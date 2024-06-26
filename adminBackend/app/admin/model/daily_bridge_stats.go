// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-14 17:06:57
// 生成路径: adminBackend/app/admin/model/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// DailyBridgeStats is the golang structure for table daily_bridge_stats.
type DailyBridgeStats struct {
	Id                 int         `orm:"id,primary" json:"id"`                          // 序号
	Date               *gtime.Time `orm:"date" json:"date"`                              // 时间
	Coin               string      `orm:"coin" json:"coin"`                              // 币种
	ChainType          string      `orm:"chain_type" json:"chainType"`                   // 链类型
	TransferIn         float64     `orm:"transfer_in" json:"transferIn"`                 // 100
	TransferOut        float64     `orm:"transfer_out" json:"transferOut"`               // 100
	TransferDifference float64     `orm:"transfer_difference" json:"transferDifference"` // 跨链差额
	TransferFee        float64     `orm:"transfer_fee" json:"transferFee"`               // 跨链手续费
	PlatformAssets     float64     `orm:"platform_assets" json:"platformAssets"`         // 平台资产
	CreatedAt          *gtime.Time `orm:"created_at" json:"createdAt"`                   //
	UpdatedAt          *gtime.Time `orm:"updated_at" json:"updatedAt"`                   //
}
