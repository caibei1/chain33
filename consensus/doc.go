package consensus

//共识相关的模块
//模块功能：模块主要的功能是实现共识排序的功能，包括完整的共识的实现。
/**
1. solo
提供单节点排序功能，适用于开发过程，便于调试。
多节点情况下部署solo,只有主节点处理交易排序打包，其余节点不作处理。
主节点打包完之后，广播给其它节点。

2. raft
3. bft

*/
//接口设计：
