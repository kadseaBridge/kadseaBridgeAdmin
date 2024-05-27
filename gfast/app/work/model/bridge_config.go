// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-27 18:48:21
// 生成路径: gfast/app/work/model/bridge_config.go
// 生成人：gfast
// ==========================================================================


package model
import "github.com/gogf/gf/os/gtime"
// BridgeConfig is the golang structure for table bridge_config.
type BridgeConfig struct {	
         Id       int64         `orm:"id,primary" json:"id"`    // 编号    
         SourceChainId    string         `orm:"source_chainId" json:"sourceChainId"`    // 当前链    
         TargetChainId    string         `orm:"target_chainId" json:"targetChainId"`    // 目标链    
         SourceCoinAddress    string         `orm:"source_coin_address" json:"sourceCoinAddress"`    // 当前代币    
         FeeFixed    float64         `orm:"fee_fixed" json:"feeFixed"`    // 固定金额手续费，    
         FeePercent    float64         `orm:"fee_percent" json:"feePercent"`    // 手续费百分比，如果不需要则设置为0    
         DayTotal    float64         `orm:"day_total" json:"dayTotal"`    // 单用户每日跨链不需要审核的总额度    
         OnceTotal    float64         `orm:"once_total" json:"onceTotal"`    // 单用户每次跨链不需要审核的额度    
         IsEnable    int         `orm:"isEnable" json:"isEnable"`    // 是否可用    
         TargetCoinAddress    string         `orm:"target_coin_address" json:"targetCoinAddress"`    // 目标代币    
         UpdateAt    *gtime.Time         `orm:"update_at" json:"updateAt"`    //    
         CreateAt    *gtime.Time         `orm:"create_at" json:"createAt"`    //    
}