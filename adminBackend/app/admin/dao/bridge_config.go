// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: adminBackend/app/admin/dao/bridge_config.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
)

// bridgeConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type bridgeConfigDao struct {
	*internal.BridgeConfigDao
}

var (
	// BridgeConfig is globally public accessible object for table tools_gen_table operations.
	BridgeConfig = bridgeConfigDao{
		internal.NewBridgeConfigDao(),
	}
)

// Fill with you ideas below.

// BridgeConfigSearchReq 分页请求参数
type BridgeConfigSearchReq struct {
	SourceChainId     string `p:"sourceChainId"`     //当前链
	SourceCoinAddress string `p:"sourceCoinAddress"` //币种
	comModel.PageReq
}

// BridgeConfigAddReq 添加操作请求参数
type BridgeConfigAddReq struct {
	SourceChainId     string  `p:"sourceChainId" v:"required#当前链不能为空"`
	TargetChainId     string  `p:"targetChainId" v:"required#对手链不能为空"`
	SourceCoinAddress string  `p:"sourceCoinAddress" v:"required#币种不能为空"`
	FeeFixed          float64 `p:"feeFixed" v:"required#固定金额手续费，不能为空"`
	FeePercent        float64 `p:"feePercent" v:"required#手续费百分比，如果不需要则设置为0不能为空"`
	DayTotal          float64 `p:"dayTotal" v:"required#跨入每日审核数量不能为空"`
	OnceTotal         float64 `p:"onceTotal" v:"required#跨入单次审核数量不能为空"`
	IsEnable          int     `p:"isEnable" v:"required#状态不能为空"`
	TargetCoinAddress string  `p:"targetCoinAddress"`
}

// BridgeConfigEditReq 修改操作请求参数
type BridgeConfigEditReq struct {
	Id                int64   `p:"id" v:"required#主键ID不能为空"`
	SourceChainId     string  `p:"sourceChainId" v:"required#当前链不能为空"`
	TargetChainId     string  `p:"targetChainId" v:"required#对手链不能为空"`
	SourceCoinAddress string  `p:"sourceCoinAddress" v:"required#币种不能为空"`
	FeeFixed          float64 `p:"feeFixed" v:"required#固定金额手续费，不能为空"`
	FeePercent        float64 `p:"feePercent" v:"required#手续费百分比，如果不需要则设置为0不能为空"`
	DayTotal          float64 `p:"dayTotal" v:"required#跨入每日审核数量不能为空"`
	OnceTotal         float64 `p:"onceTotal" v:"required#跨入单次审核数量不能为空"`
	IsEnable          int     `p:"isEnable" v:"required#状态不能为空"`
	TargetCoinAddress string  `p:"targetCoinAddress"`
}
