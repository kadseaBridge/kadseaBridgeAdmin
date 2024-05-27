// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-05-27 18:48:21
// 生成路径: gfast/app/work/dao/bridge_config.go
// 生成人：gfast
// ==========================================================================


package dao
import (
    comModel "gfast/app/common/model"
    "gfast/app/work/dao/internal"    
    "github.com/gogf/gf/os/gtime"    
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
    SourceChainId  string `p:"sourceChainId"` //当前链    
    TargetChainId  string `p:"targetChainId"` //目标链    
    SourceCoinAddress  string `p:"sourceCoinAddress"` //当前代币    
    IsEnable  string `p:"isEnable"` //是否可用    
    TargetCoinAddress  string `p:"targetCoinAddress"` //目标代币    
    UpdateAt  *gtime.Time `p:"updateAt"` //    
    CreateAt  *gtime.Time `p:"createAt"` //    
    comModel.PageReq
}
// BridgeConfigAddReq 添加操作请求参数
type BridgeConfigAddReq struct {    
    SourceChainId  string   `p:"sourceChainId" v:"required#当前链不能为空"`    
    TargetChainId  string   `p:"targetChainId" v:"required#目标链不能为空"`    
    SourceCoinAddress  string   `p:"sourceCoinAddress" v:"required#当前代币不能为空"`    
    FeeFixed  float64   `p:"feeFixed" v:"required#固定金额手续费，不能为空"`    
    FeePercent  float64   `p:"feePercent" v:"required#手续费百分比，如果不需要则设置为0不能为空"`    
    DayTotal  float64   `p:"dayTotal" v:"required#单用户每日跨链不需要审核的总额度不能为空"`    
    OnceTotal  float64   `p:"onceTotal" v:"required#单用户每次跨链不需要审核的额度不能为空"`    
    IsEnable  int   `p:"isEnable" v:"required#是否可用不能为空"`    
    TargetCoinAddress  string   `p:"targetCoinAddress" v:"required#目标代币不能为空"`    
    UpdateAt  *gtime.Time   `p:"updateAt" `    
    CreateAt  *gtime.Time   `p:"createAt" `    
}
// BridgeConfigEditReq 修改操作请求参数
type BridgeConfigEditReq struct {
    Id    int64  `p:"id" v:"required#主键ID不能为空"`    
    SourceChainId  string `p:"sourceChainId" v:"required#当前链不能为空"`    
    TargetChainId  string `p:"targetChainId" v:"required#目标链不能为空"`    
    SourceCoinAddress  string `p:"sourceCoinAddress" v:"required#当前代币不能为空"`    
    FeeFixed  float64 `p:"feeFixed" v:"required#固定金额手续费，不能为空"`    
    FeePercent  float64 `p:"feePercent" v:"required#手续费百分比，如果不需要则设置为0不能为空"`    
    DayTotal  float64 `p:"dayTotal" v:"required#单用户每日跨链不需要审核的总额度不能为空"`    
    OnceTotal  float64 `p:"onceTotal" v:"required#单用户每次跨链不需要审核的额度不能为空"`    
    IsEnable  int `p:"isEnable" v:"required#是否可用不能为空"`    
    TargetCoinAddress  string `p:"targetCoinAddress" v:"required#目标代币不能为空"`    
    UpdateAt  *gtime.Time `p:"updateAt" `    
    CreateAt  *gtime.Time `p:"createAt" `    
}
