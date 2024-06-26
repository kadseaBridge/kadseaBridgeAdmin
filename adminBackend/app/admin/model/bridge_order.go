// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: adminBackend/app/admin/model/bridge_order.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// BridgeOrder is the golang structure for table bridge_order.
type BridgeOrder struct {
	Id                int64       `orm:"id,primary" json:"id"`                         //
	SourceAddress     string      `orm:"source_address" json:"sourceAddress"`          // 转出币种
	TargetAddress     string      `orm:"target_address" json:"targetAddress"`          // 转入币种
	SourceCoinAddress string      `orm:"source_coin_address" json:"sourceCoinAddress"` // 转出钱包地址
	TargetCoinAddress string      `orm:"target_coin_address" json:"targetCoinAddress"` // 转入钱包地址
	SourceChainId     string      `orm:"source_chainId" json:"sourceChainId"`          // 转出链
	TargetChainId     string      `orm:"target_chainId" json:"targetChainId"`          // 转入链
	TransactionHash   string      `orm:"transaction_hash" json:"transactionHash"`      // 交易哈希
	OrderId           string      `orm:"orderId" json:"orderId"`                       // 跨链订单id
	Amount            float64     `orm:"amount" json:"amount"`                         // 数量
	Status            int         `orm:"status" json:"status"`                         // 跨链记录状态
	CreateAt          *gtime.Time `orm:"create_at" json:"createAt"`                    // 发起时间
	UpdateAt          *gtime.Time `orm:"update_at" json:"updateAt"`                    // 结束时间
	BlockNumber       int64       `orm:"block_number" json:"blockNumber"`              // 区块高度
	Fee               float64     `orm:"fee" json:"fee"`                               // 手续费
	Remark            string      `orm:"remark" json:"remark"`                         // 评论
}

type BridgeOrderRsp struct {
	Id              int64       `orm:"id,primary" json:"id"`                     //
	SourceAddress   string      `orm:"source_address" json:"sourceAddress"`      // 转出币种
	TargetAddress   string      `orm:"target_address" json:"targetAddress"`      // 转入币种
	SourceCoinName  string      `orm:"source_coin_name" json:"sourceCoinName"`   // 转出钱包地址
	TargetCoinName  string      `orm:"target_coin_name" json:"targetCoinName"`   // 转入钱包地址
	SourceChainName string      `orm:"source_chain_name" json:"sourceChainName"` // 转出链
	TargetChainName string      `orm:"target_chain_name" json:"targetChainName"` // 转入链
	InHash          string      `orm:"in_hash" json:"inHash"`                    // 转入哈希
	OutHash         string      `orm:"out_hash" json:"outHash"`                  // 转出哈希
	OrderId         string      `orm:"orderId" json:"orderId"`                   // 跨链订单id
	Amount          float64     `orm:"amount" json:"amount"`                     // 数量
	Status          int         `orm:"status" json:"status"`                     // 跨链记录状态
	CreateAt        *gtime.Time `orm:"create_at" json:"createAt"`                // 发起时间
	UpdateAt        *gtime.Time `orm:"update_at" json:"updateAt"`                // 结束时间
	BlockNumber     int64       `orm:"block_number" json:"blockNumber"`          // 区块高度
	Fee             float64     `orm:"fee" json:"fee"`                           // 手续费
	GasFee          float64     `orm:"gas_fee" json:"gasFee"`                    // gas手续费
	Remark          string      `orm:"remark" json:"remark"`                     // 评论
}
