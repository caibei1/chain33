// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statistic.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 手续费
type TotalFee struct {
	Fee     int64 `protobuf:"varint,1,opt,name=fee" json:"fee,omitempty"`
	TxCount int64 `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *TotalFee) Reset()                    { *m = TotalFee{} }
func (m *TotalFee) String() string            { return proto.CompactTextString(m) }
func (*TotalFee) ProtoMessage()               {}
func (*TotalFee) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *TotalFee) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TotalFee) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

// 查询symbol代币总额
type ReqGetTotalCoins struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	StartKey  []byte `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqGetTotalCoins) Reset()                    { *m = ReqGetTotalCoins{} }
func (m *ReqGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReqGetTotalCoins) ProtoMessage()               {}
func (*ReqGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ReqGetTotalCoins) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetTotalCoins) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetTotalCoins) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *ReqGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqGetTotalCoins) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// 查询symbol代币总额应答
type ReplyGetTotalCoins struct {
	Count   int64  `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Num     int64  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	Amount  int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	NextKey []byte `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (m *ReplyGetTotalCoins) Reset()                    { *m = ReplyGetTotalCoins{} }
func (m *ReplyGetTotalCoins) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetTotalCoins) ProtoMessage()               {}
func (*ReplyGetTotalCoins) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *ReplyGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

// 迭代查询symbol代币总额
type IterateRangeByStateHash struct {
	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start     []byte `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End       []byte `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *IterateRangeByStateHash) Reset()                    { *m = IterateRangeByStateHash{} }
func (m *IterateRangeByStateHash) String() string            { return proto.CompactTextString(m) }
func (*IterateRangeByStateHash) ProtoMessage()               {}
func (*IterateRangeByStateHash) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *IterateRangeByStateHash) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *IterateRangeByStateHash) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IterateRangeByStateHash) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *IterateRangeByStateHash) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TicketStatistic struct {
	// 当前在挖的ticket
	CurrentOpenCount int64 `protobuf:"varint,1,opt,name=currentOpenCount" json:"currentOpenCount,omitempty"`
	// 一共挖到的ticket
	TotalMinerCount int64 `protobuf:"varint,2,opt,name=totalMinerCount" json:"totalMinerCount,omitempty"`
	// 一共取消的ticket
	TotalCancleCount int64 `protobuf:"varint,3,opt,name=totalCancleCount" json:"totalCancleCount,omitempty"`
}

func (m *TicketStatistic) Reset()                    { *m = TicketStatistic{} }
func (m *TicketStatistic) String() string            { return proto.CompactTextString(m) }
func (*TicketStatistic) ProtoMessage()               {}
func (*TicketStatistic) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *TicketStatistic) GetCurrentOpenCount() int64 {
	if m != nil {
		return m.CurrentOpenCount
	}
	return 0
}

func (m *TicketStatistic) GetTotalMinerCount() int64 {
	if m != nil {
		return m.TotalMinerCount
	}
	return 0
}

func (m *TicketStatistic) GetTotalCancleCount() int64 {
	if m != nil {
		return m.TotalCancleCount
	}
	return 0
}

type TicketMinerInfo struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	// 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status     int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	PrevStatus int32 `protobuf:"varint,3,opt,name=prevStatus" json:"prevStatus,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,4,opt,name=isGenesis" json:"isGenesis,omitempty"`
	// 创建ticket时间
	CreateTime int64 `protobuf:"varint,5,opt,name=createTime" json:"createTime,omitempty"`
	// ticket挖矿时间
	MinerTime int64 `protobuf:"varint,6,opt,name=minerTime" json:"minerTime,omitempty"`
	// 关闭ticket时间
	CloseTime int64 `protobuf:"varint,7,opt,name=closeTime" json:"closeTime,omitempty"`
	// 挖到的币的数目
	MinerValue   int64  `protobuf:"varint,8,opt,name=minerValue" json:"minerValue,omitempty"`
	MinerAddress string `protobuf:"bytes,9,opt,name=minerAddress" json:"minerAddress,omitempty"`
}

func (m *TicketMinerInfo) Reset()                    { *m = TicketMinerInfo{} }
func (m *TicketMinerInfo) String() string            { return proto.CompactTextString(m) }
func (*TicketMinerInfo) ProtoMessage()               {}
func (*TicketMinerInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *TicketMinerInfo) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *TicketMinerInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TicketMinerInfo) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *TicketMinerInfo) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *TicketMinerInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *TicketMinerInfo) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

type TotalAmount struct {
	// 统计的总数
	Total int64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *TotalAmount) Reset()                    { *m = TotalAmount{} }
func (m *TotalAmount) String() string            { return proto.CompactTextString(m) }
func (*TotalAmount) ProtoMessage()               {}
func (*TotalAmount) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *TotalAmount) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 查询symbol在合约中的代币总额，如果execAddr为空，则为查询symbol在所有合约中的代币总额
type ReqGetExecBalance struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Addr      []byte `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	ExecAddr  []byte `protobuf:"bytes,4,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqGetExecBalance) Reset()                    { *m = ReqGetExecBalance{} }
func (m *ReqGetExecBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqGetExecBalance) ProtoMessage()               {}
func (*ReqGetExecBalance) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *ReqGetExecBalance) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetExecBalance) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetExecBalance) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type ExecBalanceItem struct {
	ExecAddr []byte `protobuf:"bytes,1,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Frozen   int64  `protobuf:"varint,2,opt,name=frozen" json:"frozen,omitempty"`
	Active   int64  `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
}

func (m *ExecBalanceItem) Reset()                    { *m = ExecBalanceItem{} }
func (m *ExecBalanceItem) String() string            { return proto.CompactTextString(m) }
func (*ExecBalanceItem) ProtoMessage()               {}
func (*ExecBalanceItem) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *ExecBalanceItem) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ExecBalanceItem) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *ExecBalanceItem) GetActive() int64 {
	if m != nil {
		return m.Active
	}
	return 0
}

// 查询symbol在合约中的代币总额应答
type ReplyGetExecBalance struct {
	Amount       int64              `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	AmountFrozen int64              `protobuf:"varint,2,opt,name=amountFrozen" json:"amountFrozen,omitempty"`
	AmountActive int64              `protobuf:"varint,3,opt,name=amountActive" json:"amountActive,omitempty"`
	Items        []*ExecBalanceItem `protobuf:"bytes,4,rep,name=items" json:"items,omitempty"`
}

func (m *ReplyGetExecBalance) Reset()                    { *m = ReplyGetExecBalance{} }
func (m *ReplyGetExecBalance) String() string            { return proto.CompactTextString(m) }
func (*ReplyGetExecBalance) ProtoMessage()               {}
func (*ReplyGetExecBalance) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ReplyGetExecBalance) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountFrozen() int64 {
	if m != nil {
		return m.AmountFrozen
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountActive() int64 {
	if m != nil {
		return m.AmountActive
	}
	return 0
}

func (m *ReplyGetExecBalance) GetItems() []*ExecBalanceItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// 迭代查询symbol代币在合约中的总额
type IterateExecBalanceByStateHash struct {
	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Addr      []byte `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Prefix    []byte `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *IterateExecBalanceByStateHash) Reset()                    { *m = IterateExecBalanceByStateHash{} }
func (m *IterateExecBalanceByStateHash) String() string            { return proto.CompactTextString(m) }
func (*IterateExecBalanceByStateHash) ProtoMessage()               {}
func (*IterateExecBalanceByStateHash) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *IterateExecBalanceByStateHash) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *IterateExecBalanceByStateHash) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *IterateExecBalanceByStateHash) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func init() {
	proto.RegisterType((*TotalFee)(nil), "types.TotalFee")
	proto.RegisterType((*ReqGetTotalCoins)(nil), "types.ReqGetTotalCoins")
	proto.RegisterType((*ReplyGetTotalCoins)(nil), "types.ReplyGetTotalCoins")
	proto.RegisterType((*IterateRangeByStateHash)(nil), "types.IterateRangeByStateHash")
	proto.RegisterType((*TicketStatistic)(nil), "types.TicketStatistic")
	proto.RegisterType((*TicketMinerInfo)(nil), "types.TicketMinerInfo")
	proto.RegisterType((*TotalAmount)(nil), "types.TotalAmount")
	proto.RegisterType((*ReqGetExecBalance)(nil), "types.ReqGetExecBalance")
	proto.RegisterType((*ExecBalanceItem)(nil), "types.ExecBalanceItem")
	proto.RegisterType((*ReplyGetExecBalance)(nil), "types.ReplyGetExecBalance")
	proto.RegisterType((*IterateExecBalanceByStateHash)(nil), "types.IterateExecBalanceByStateHash")
}

func init() { proto.RegisterFile("statistic.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0xbb, 0x4d, 0x9a, 0x9e, 0x16, 0x12, 0xc7, 0x52, 0x17, 0x51, 0x91, 0xd1, 0x8b, 0x22,
	0x12, 0xc1, 0x80, 0xf7, 0x6d, 0xb0, 0x35, 0x88, 0x08, 0xdb, 0xe2, 0x85, 0xe0, 0xc5, 0x74, 0x72,
	0xd2, 0x0e, 0xee, 0xce, 0xc6, 0x99, 0x49, 0x49, 0x7c, 0x0b, 0xf5, 0x11, 0xfa, 0xa2, 0x32, 0x67,
	0x67, 0x37, 0xbb, 0xa9, 0x82, 0x78, 0xb5, 0xf3, 0x7d, 0x67, 0xe6, 0xfc, 0x7f, 0x0b, 0x7d, 0xeb,
	0x84, 0x53, 0xd6, 0x29, 0x39, 0x9c, 0x9b, 0xc2, 0x15, 0xac, 0xe3, 0x56, 0x73, 0xb4, 0xfc, 0x0d,
	0xf4, 0x2e, 0x0a, 0x27, 0xb2, 0x53, 0x44, 0x36, 0x80, 0x78, 0x86, 0x98, 0x44, 0x4f, 0xa3, 0xa3,
	0x38, 0xf5, 0x47, 0x96, 0xc0, 0x8e, 0x5b, 0x8e, 0x8b, 0x85, 0x76, 0xc9, 0x16, 0xb1, 0x15, 0xe4,
	0x3f, 0x23, 0x18, 0xa4, 0xf8, 0xed, 0x0c, 0x1d, 0x3d, 0x1f, 0x17, 0x4a, 0x5b, 0x76, 0x08, 0x5d,
	0xbb, 0xca, 0x2f, 0x8b, 0x8c, 0x7c, 0xec, 0xa6, 0x01, 0xb1, 0x47, 0xb0, 0xeb, 0xc3, 0xe3, 0x3b,
	0x61, 0xaf, 0xc9, 0xd1, 0x7e, 0xba, 0x26, 0xd8, 0x43, 0xe8, 0x59, 0x27, 0x8c, 0x7b, 0x8f, 0xab,
	0x24, 0x26, 0x63, 0x8d, 0xd9, 0x01, 0x74, 0x24, 0x85, 0xdf, 0xa6, 0xf0, 0x25, 0xf0, 0x71, 0x70,
	0x89, 0x12, 0x4d, 0xd2, 0x29, 0xe3, 0x94, 0x88, 0x6b, 0x60, 0x29, 0xce, 0xb3, 0x55, 0x3b, 0xab,
	0xda, 0x47, 0xd4, 0xf4, 0x31, 0x80, 0x58, 0x2f, 0xf2, 0x50, 0x96, 0x3f, 0x7a, 0xaf, 0x22, 0xa7,
	0x8b, 0x31, 0x91, 0x01, 0xf9, 0x26, 0x68, 0x5c, 0x52, 0x7a, 0xdb, 0x94, 0x5e, 0x05, 0xf9, 0x02,
	0x1e, 0x4c, 0x1c, 0x1a, 0xe1, 0x30, 0x15, 0xfa, 0x0a, 0x4f, 0x56, 0xe7, 0x75, 0x51, 0xad, 0x92,
	0xa3, 0xcd, 0x92, 0x0f, 0xa0, 0x43, 0x25, 0x86, 0x66, 0x94, 0xc0, 0xa7, 0x84, 0x7a, 0x1a, 0x7a,
	0xe0, 0x8f, 0x7f, 0x2e, 0x9f, 0xff, 0x8a, 0xa0, 0x7f, 0xa1, 0xe4, 0x57, 0x74, 0xe7, 0xd5, 0x50,
	0xd9, 0x0b, 0x18, 0xc8, 0x85, 0x31, 0xa8, 0xdd, 0xc7, 0x39, 0xea, 0x71, 0xa3, 0xde, 0x3b, 0x3c,
	0x3b, 0x82, 0xbe, 0xf3, 0xed, 0xf9, 0xa0, 0x34, 0x9a, 0xe6, 0x74, 0x37, 0x69, 0xef, 0x95, 0xa8,
	0xb1, 0xd0, 0x32, 0xc3, 0x71, 0xa3, 0x39, 0x77, 0x78, 0x7e, 0xbb, 0x55, 0x65, 0x45, 0x0e, 0x26,
	0x7a, 0x56, 0xf8, 0xd1, 0x3a, 0xa2, 0x26, 0xd3, 0xb0, 0x12, 0x35, 0xa6, 0x65, 0x71, 0xc2, 0x2d,
	0x2c, 0x05, 0xef, 0xa4, 0x01, 0xb1, 0x27, 0x00, 0x73, 0x83, 0x37, 0xe7, 0xa5, 0x2d, 0x26, 0x5b,
	0x83, 0xf1, 0x9d, 0x55, 0xf6, 0x0c, 0x35, 0x5a, 0x65, 0xa9, 0x2f, 0xbd, 0x74, 0x4d, 0xf8, 0xd7,
	0xd2, 0xa0, 0x70, 0x78, 0xa1, 0x72, 0xa4, 0xf5, 0x88, 0xd3, 0x06, 0xe3, 0x5f, 0xe7, 0x3e, 0x3d,
	0x32, 0x77, 0xc9, 0xbc, 0x26, 0xbc, 0x55, 0x66, 0x85, 0x2d, 0x1f, 0xef, 0x94, 0xd6, 0x9a, 0xf0,
	0xbe, 0xe9, 0xea, 0x27, 0x91, 0x2d, 0x30, 0xe9, 0x95, 0xbe, 0xd7, 0x0c, 0xe3, 0xb0, 0x4f, 0xe8,
	0x78, 0x3a, 0x35, 0x68, 0x6d, 0xb2, 0x4b, 0x15, 0xb7, 0x38, 0xfe, 0x0c, 0xf6, 0x68, 0x35, 0x8f,
	0xcb, 0xdd, 0x3a, 0x80, 0x0e, 0x35, 0xb2, 0xda, 0x4d, 0x02, 0xfc, 0x47, 0x04, 0xf7, 0x4a, 0x71,
	0xbd, 0x5d, 0xa2, 0x3c, 0x11, 0x99, 0xd0, 0x12, 0xff, 0x53, 0x5d, 0x0c, 0xb6, 0xc5, 0x74, 0x6a,
	0xc2, 0x56, 0xd1, 0xd9, 0x8f, 0xc5, 0x2b, 0xc6, 0xe7, 0x14, 0x56, 0xba, 0xc6, 0x7f, 0xd5, 0xd6,
	0x17, 0xe8, 0x37, 0x92, 0x99, 0x38, 0xcc, 0x5b, 0x6e, 0xa2, 0xbb, 0x6e, 0x66, 0xa6, 0xf8, 0x8e,
	0x3a, 0xac, 0x56, 0x40, 0x24, 0x32, 0xe9, 0xd4, 0x0d, 0xd6, 0x22, 0x23, 0xc4, 0x6f, 0x23, 0xb8,
	0x5f, 0x69, 0x77, 0xa3, 0xe8, 0x20, 0xca, 0xa8, 0x25, 0x4a, 0x0e, 0xfb, 0xe5, 0xe9, 0xb4, 0x19,
	0xa5, 0xc5, 0xad, 0xef, 0x1c, 0x37, 0x23, 0xb6, 0x38, 0xf6, 0x12, 0x3a, 0xca, 0x61, 0xee, 0x37,
	0x29, 0x3e, 0xda, 0x7b, 0x7d, 0x38, 0xa4, 0xdf, 0xe2, 0x70, 0xa3, 0xd4, 0xb4, 0xbc, 0xc4, 0x15,
	0x3c, 0x0e, 0x82, 0x6f, 0x5c, 0xf8, 0x77, 0xd9, 0x57, 0xb3, 0xd8, 0x6a, 0xcc, 0xe2, 0x10, 0xba,
	0x73, 0x83, 0x33, 0xb5, 0x0c, 0x13, 0x0a, 0xe8, 0xe4, 0xf9, 0x67, 0x7e, 0xa5, 0x5c, 0x26, 0x2e,
	0x87, 0xa3, 0xd1, 0x50, 0xea, 0x57, 0xf2, 0x5a, 0x28, 0x3d, 0x1a, 0xd5, 0x5f, 0xca, 0xf3, 0xb2,
	0x4b, 0x3f, 0xf3, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6e, 0xf7, 0xf8, 0xdf, 0x05,
	0x00, 0x00,
}
