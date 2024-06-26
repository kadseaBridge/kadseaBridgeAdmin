// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-12 23:29:47
// 生成路径: adminBackend/app/admin/model/chain.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// Chain is the golang structure for table chain.
type Chain struct {
	Id                    int64       `orm:"id,primary" json:"id"`                                 //
	Name                  string      `orm:"name" json:"name"`                                     // 链名
	ChainId               string      `orm:"chainId" json:"chainId"`                               // 链Id
	Rpc                   string      `orm:"rpc" json:"rpc"`                                       // 链rpc
	Explorer              string      `orm:"explorer" json:"explorer"`                             // 链浏览器地址
	IsEnable              int         `orm:"isEnable" json:"isEnable"`                             // 状态
	CreateAt              *gtime.Time `orm:"create_at" json:"createAt"`                            // 创建时间
	UpdateAt              *gtime.Time `orm:"update_at" json:"updateAt"`                            // 更新时间
	ChainType             int         `orm:"type" json:"type"`                                     // 0 ethereum 1 tron  2 btc
	Icon                  string      `orm:"icon" json:"icon"`                                     // 链图标
	IsDelete              int         `orm:"isDelete" json:"isDelete"`                             // 软删除
	MaxGasPrice           float64     `orm:"max_gas_price" json:"maxGasPrice"`                     // 支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时， 为0时，所有订单都不进行支付， 如果想要所有的订单都进行支付，不管gasprice 则设置很大   单位：wei
	EnablePay             int         `orm:"enablePay" json:"enablePay"`                           // 该公链作为目标时，是否允许支付
	BridgeContractAddress string      `orm:"bridge_contract_address" json:"bridgeContractAddress"` // 跨链桥合约地址
}
