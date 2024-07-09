// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-07-08 18:02:01
// 生成路径: gfast/app/admin/model/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// DailyBridgeStats is the golang structure for table daily_bridge_stats.
type DailyBridgeStats struct {
	Id                 int         `orm:"id,primary" json:"id"`                          // id
	Date               *gtime.Time `orm:"date" json:"date"`                              // 时间
	CoinAddress        string      `orm:"coin_address" json:"coinAddress"`               // 币种
	ChainId            string      `orm:"chain_id" json:"chainId"`                       // 链类型
	TransferIn         float64     `orm:"transfer_in" json:"transferIn"`                 // 跨链转入
	TransferOut        float64     `orm:"transfer_out" json:"transferOut"`               // 跨链转出
	TransferDifference float64     `orm:"transfer_difference" json:"transferDifference"` // 跨链差额
	Fee                float64     `orm:"fee" json:"fee"`                                // 跨链手续费
	PlatformAssets     float64     `orm:"platform_assets" json:"platformAssets"`         // 平台资产
	FinancialBalance   float64     `orm:"financial_balance" json:"financialBalance"`     // 财务地址余额
	CreatedAt          *gtime.Time `orm:"created_at" json:"createdAt"`                   //
	UpdatedAt          *gtime.Time `orm:"updated_at" json:"updatedAt"`                   //
	TronNotActiveFee   float64     `p:"tron_not_active_fee"`
	TronTrc20HaveNoFee float64     `p:"tron_trc20_have_no_fee"`
}
