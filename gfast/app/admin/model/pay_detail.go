// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-08 23:23:41
// 生成路径: gfast/app/admin/model/pay_detail.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// PayDetail is the golang structure for table pay_detail.
type PayDetail struct {
	Id              int64       `orm:"id,primary" json:"id"`                    //
	OrderId         string      `orm:"order_id" json:"orderId"`                 //
	PayChainId      string      `orm:"pay_chainId" json:"payChainId"`           //
	PayCoinAddress  string      `orm:"pay_coin_address" json:"payCoinAddress"`  //
	TransactionHash string      `orm:"transaction_hash" json:"transactionHash"` //
	PayAmount       float64     `orm:"pay_amount" json:"payAmount"`             //
	PayTime         *gtime.Time `orm:"pay_time" json:"payTime"`                 //
	PayAddress      string      `orm:"pay_address" json:"payAddress"`           //
	CreateAt        *gtime.Time `orm:"create_at" json:"createAt"`               //
	UpdateAt        *gtime.Time `orm:"update_at" json:"updateAt"`               //
	Remark          string      `orm:"remark" json:"remark"`                    //
	BlockNumber     int64       `orm:"block_number" json:"blockNumber"`         //
	PayGasFee       float64     `orm:"pay_gas_fee" json:"payGasFee"`            //
}
