// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: adminBackend/app/admin/model/coin.go
// 生成人：jimmy
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// Coin is the golang structure for table coin.
type Coin struct {
	Id       int64       `orm:"id,primary" json:"id"`      //
	Name     string      `orm:"name" json:"name"`          // 代币名称
	Symbol   string      `orm:"symbol" json:"symbol"`      // 代币简称
	ChainId  string      `orm:"chainId" json:"chainId"`    // 链Id
	Address  string      `orm:"address" json:"address"`    // 代币地址
	IsEnable int         `orm:"isEnable" json:"isEnable"`  // 是否上架
	CreateAt *gtime.Time `orm:"create_at" json:"createAt"` //
	UpdateAt *gtime.Time `orm:"update_at" json:"updateAt"` //
	Type     string      `orm:"type" json:"tokenType"`     // 币种类型
	Decimals int         `orm:"decimals" json:"decimals"`  // 精度
	Icon     string      `orm:"icon" json:"icon"`          // 代币图标
}
