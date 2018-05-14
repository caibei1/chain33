package evm

import (
	"math/big"

	"encoding/binary"
	"encoding/hex"
	log "github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/account"
	"gitlab.33.cn/chain33/chain33/client"
	"gitlab.33.cn/chain33/chain33/executor/drivers"
	"gitlab.33.cn/chain33/chain33/executor/drivers/evm/vm/common"
	"gitlab.33.cn/chain33/chain33/executor/drivers/evm/vm/model"
	"gitlab.33.cn/chain33/chain33/executor/drivers/evm/vm/runtime"
	"gitlab.33.cn/chain33/chain33/executor/drivers/evm/vm/state"
	"gitlab.33.cn/chain33/chain33/types"
	"os"
)

const (
	// 交易payload中，前8个字节固定存储转账信息
	BALANCE_SIZE = 8
	// 在一个交易中，GasLimit设置为交易费用的倍数(整数)
	TX_GAS_TIMES_FEE = 2
)

var (
	GasPrice = big.NewInt(1)
	//目前这里先用默认配置，后面在增加具体实现
	VMConfig = &runtime.Config{}
)

// EVM执行器结构
type EVMExecutor struct {
	drivers.DriverBase
	vmCfg *runtime.Config

	mStateDB *state.MemoryStateDB
}

func init() {
	evm := NewEVMExecutor()
	// TODO 注册的驱动高度需要更新为上线时的正确高度
	drivers.Register(evm.GetName(), evm, 0)
}

func NewEVMExecutor() *EVMExecutor {
	exec := &EVMExecutor{}

	exec.vmCfg = VMConfig
	exec.vmCfg.Tracer = runtime.NewJSONLogger(os.Stdout)

	// 打开EVM调试开关
	//exec.vmCfg.Debug = true

	exec.SetChild(exec)
	return exec
}

func (evm *EVMExecutor) GetName() string {
	return "evm"
}

func (evm *EVMExecutor) SetEnv(height, blockTime int64, coinBase string, difficulty uint64) {
	// 需要从这里识别出当前执行的Transaction所在的区块高度
	// 因为执行器框架在调用每一个Transaction时，都会先设置StateDB，在设置区块环境
	// 因此，在这里判断当前设置的区块高度和上一次缓存的区块高度是否为同一高度，即可判断是否在同一个区块内执行的Transaction
	if height != evm.DriverBase.GetHeight() || blockTime != evm.DriverBase.GetBlockTime() || coinBase != evm.DriverBase.GetCoinBase() || difficulty != evm.DriverBase.GetDifficulty() {
		// 这时说明区块发生了变化，需要集成原来的设置逻辑，并执行自定义操作
		evm.DriverBase.SetEnv(height, blockTime, coinBase, difficulty)

		// 在生成新的区块状态DB之前，把先前区块中生成的合约日志集中打印出来
		if evm.mStateDB != nil {
			evm.mStateDB.WritePreimages(height)
		}

		// 重新初始化MemoryStateDB
		// 需要注意的时，在执行器中只执行单个Transaction，但是并没有提交区块的动作
		// 所以，这个mStateDB只用来缓存一个区块内执行的Transaction引起的状态数据变更
		evm.mStateDB = state.NewMemoryStateDB(evm.DriverBase.GetStateDB(), evm.DriverBase.GetLocalDB(), evm.DriverBase.GetCoinsAccount())
	}
	// 两者都和上次的设置相同，不需要任何操作
}

// 在区块上的执行操作，同一个区块内的多个交易会循环调用此方法进行处理；
// 返回的结果types.Receipt数据，将会被统一写入到本地状态数据库中；
// 本操作返回的ReceiptLog数据，在后面调用ExecLocal时会被传入，同样ExecLocal返回的数据将会被写入blockchain.db；
// FIXME 目前evm执行器暂时没有ExecLocal，执行默认逻辑，后面根据需要再考虑增加；
func (evm *EVMExecutor) Exec(tx *types.Transaction, index int) (*types.Receipt, error) {
	// 先转换消息
	msg := evm.GetMessage(tx)

	// 获取当前区块的上下文信息构造EVM上下文
	context := evm.NewEVMContext(msg)

	// 创建EVM运行时对象
	env := runtime.NewEVM(context, evm.mStateDB, *evm.vmCfg)

	isCreate := msg.To() == nil

	var (
		ret         = []byte("")
		vmerr       error
		leftOverGas = uint64(0)
		addr        common.Address
	)

	// 状态机中设置当前交易状态
	evm.mStateDB.Prepare(common.BytesToHash(tx.Hash()), index)

	// 合约执行之前，预先扣除GasLimit费用
	evm.mStateDB.SubBalance(msg.From(), big.NewInt(1).Mul(big.NewInt(int64(msg.GasLimit())), msg.GasPrice()))

	if isCreate {
		ret, addr, leftOverGas, vmerr = env.Create(runtime.AccountRef(msg.From()), msg.Data(), context.GasLimit, msg.Value())
	} else {
		addr = *msg.To()
		ret, leftOverGas, vmerr = env.Call(runtime.AccountRef(msg.From()), *msg.To(), msg.Data(), context.GasLimit, msg.Value())
	}

	if vmerr != nil {
		log.Info("VM returned with error", "err", vmerr)

		// 只有在出现账户余额不足的时候，才认为是有效的的错误
		// 其它错误情况目前不需要处理
		if vmerr == model.ErrInsufficientBalance {
			log.Error("evm contract error", vmerr)
			return nil, vmerr
		}
	}

	// 根据真实使用的Gas，将多余的费用返还
	// 此处将合约执行过程中给予的奖励也一并返还
	refundGas := leftOverGas + evm.mStateDB.GetRefund()
	evm.mStateDB.AddBalance(msg.From(), big.NewInt(1).Mul(big.NewInt(int64(refundGas)), msg.GasPrice()))

	// 计算消耗了多少Gas
	// 注意：这里和以太坊EVM计费有几处不同：
	// 1. GasPrice 始终为1，不支持动态调整  TODO 后继版本考虑支持动态调整
	// 2. 不计算IntrinsicGas，即合约代码的存储字节计费以及合约创建或调用动作本身的计费
	// 3. 因为前面计费内容较少，不考虑单独的refund奖励（合约执行过程中给予的refund支持返还）
	usedGas := msg.GasLimit() - refundGas

	// 将消耗的费用奖励给区块作者
	if context.Coinbase != common.EmptyAddress() {
		evm.mStateDB.AddBalance(context.Coinbase, big.NewInt(1).Mul(big.NewInt(int64(usedGas)), msg.GasPrice()))
	}

	logMsg := "call contract details:"
	if isCreate {
		logMsg = "create contract details:"
	}

	log.Info(logMsg, "caller address", msg.From().Str(), "contract address", addr.Str(), "usedGas", usedGas, "return data", hex.EncodeToString(ret))

	// 打印合约中生成的日志
	evm.mStateDB.PrintLogs()

	// 这里还需要再进行其它错误的判断（余额不足的错误，前面已经返回）
	// 因为其它情况的错误，即使合约执行失败，也消耗了资源，调用者需要为此支付费用(其实在本交易回滚时，不会扣除用户的任何费用)
	if vmerr != nil {
		log.Error("evm contract exec error", "error info",vmerr)
		return nil, vmerr
	}

	// 从状态机中获取数据变更和变更日志
	data, logs := evm.mStateDB.GetChangedData(evm.mStateDB.GetLastSnapshot())
	if logs != nil {
		logs = append(logs, evm.mStateDB.GetReceiptLogs(addr, isCreate)...)
	}
	if data != nil {
		receipt := &types.Receipt{Ty: types.ExecOk, KV: data, Logs: logs}
		return receipt, nil
	}
	return nil,nil
}

func (evm *EVMExecutor) GetMStateDB() *state.MemoryStateDB {
	return evm.mStateDB
}

func (evm *EVMExecutor) GetVMConfig() *runtime.Config {
	return evm.vmCfg
}

// 目前的交易中，如果是coins交易，金额是放在payload的，但是合约不行，需要修改Transaction结构
func (evm *EVMExecutor) GetMessage(tx *types.Transaction) (msg common.Message) {

	// 此处暂时不考虑消息发送签名的处理，chain33在mempool中对签名做了检查
	from := getCaller(tx)
	to := getReceiver(tx)

	// 注意Transaction中的payload内容同时包含转账金额和合约代码
	// payload[:8]为转账金额，payload[8:]为合约代码
	amount := binary.BigEndian.Uint64(tx.Payload[:BALANCE_SIZE])

	gasLimit := uint64(tx.Fee * TX_GAS_TIMES_FEE)
	// 合约的GasLimit即为调用者为本次合约调用准备支付的手续费
	msg = common.NewMessage(from, to, uint64(tx.Nonce), big.NewInt(int64(amount)), gasLimit, GasPrice, tx.Payload[BALANCE_SIZE:], false)
	return msg
}

// 构造一个新的EVM上下文对象
func (evm *EVMExecutor) NewEVMContext(msg common.Message) runtime.Context {
	return runtime.Context{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(evm.GetApi()),
		Origin:      msg.From(),
		Coinbase:    common.StringToAddress(evm.GetCoinBase()),
		BlockNumber: new(big.Int).SetInt64(evm.GetHeight()),
		Time:        new(big.Int).SetInt64(evm.GetBlockTime()),
		Difficulty:  new(big.Int).SetUint64(evm.GetDifficulty()),
		GasLimit:    msg.GasLimit(),
		GasPrice:    new(big.Int).Set(msg.GasPrice()),
	}
}

// 从交易信息中获取交易发起人地址
func getCaller(tx *types.Transaction) common.Address {
	return common.StringToAddress(account.From(tx).String())
}

// 从交易信息中获取交易目标地址，在创建合约交易中，此地址为空
func getReceiver(tx *types.Transaction) *common.Address {
	if tx.To == "" {
		return nil
	}
	addr := common.StringToAddress(tx.To)
	return &addr
}

// 检查合约调用账户是否有充足的金额进行转账交易操作
func CanTransfer(db state.StateDB, addr common.Address, amount *big.Int) bool {
	if amount.Uint64() == 0 {
		return true
	}
	return db.CanTransfer(addr, amount)
}

// 在内存数据库中执行转账操作（只修改内存中的金额）
func Transfer(db state.StateDB, sender, recipient common.Address, amount *big.Int) {
	if amount.Uint64() == 0 {
		return
	}
	db.Transfer(sender, recipient, amount)
}

// 获取制定高度区块的哈希
func GetHashFn(api client.QueueProtocolAPI) func(blockHeight uint64) common.Hash {
	return func(blockHeight uint64) common.Hash {
		if api != nil {
			reply, err := api.GetBlockHash(&types.ReqInt{int64(blockHeight)})
			if nil != err {
				log.Error("Call GetBlockHash Failed.", err)
			}
			return common.BytesToHash(reply.Hash)
		}
		return common.Hash{}
	}
}
