// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: gfast/app/admin/dao/coin.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
)

// coinDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type coinDao struct {
	*internal.CoinDao
}

var (
	// Coin is globally public accessible object for table tools_gen_table operations.
	Coin = coinDao{
		internal.NewCoinDao(),
	}
)

// Fill with you ideas below.

// CoinSearchReq 分页请求参数
type CoinSearchReq struct {
	Symbol  string `p:"symbol"`  //代币简称
	ChainId string `p:"chainId"` //链Id
	Address string `p:"address"` //代币地址
	comModel.PageReq
}

// CoinAddReq 添加操作请求参数
type CoinAddReq struct {
	Name      string `p:"name" v:"required#代币名称不能为空"`
	Symbol    string `p:"symbol" v:"required#代币简称不能为空"`
	ChainId   string `p:"chainId" v:"required#链Id不能为空"`
	Address   string `p:"address" v:"required#代币地址不能为空"`
	IsEnable  int    `p:"isEnable" v:"required#是否上架不能为空"`
	TokenType string `p:"type" v:"required#币种类型不能为空"`
	Decimals  int    `p:"decimals" v:"required#精度不能为空"`
	Icon      string `p:"icon" v:"required#代币图标不能为空"`
}

// CoinEditReq 修改操作请求参数
type CoinEditReq struct {
	Id        int64  `p:"id" v:"required#主键ID不能为空"`
	Name      string `p:"name" v:"required#代币名称不能为空"`
	Symbol    string `p:"symbol" v:"required#代币简称不能为空"`
	ChainId   string `p:"chainId" v:"required#链Id不能为空"`
	Address   string `p:"address" v:"required#代币地址不能为空"`
	IsEnable  int    `p:"isEnable" v:"required#是否上架不能为空"`
	TokenType string `p:"tokenType" v:"required#币种类型不能为空"`
	Decimals  int    `p:"decimals" v:"required#精度不能为空"`
	Icon      string `p:"icon" v:"required#代币图标不能为空"`
}
