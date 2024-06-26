// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: gfast/app/admin/model/bridge_config.go
// 生成人：jimmy
// ==========================================================================


package model
import "github.com/gogf/gf/os/gtime"
// BridgeConfig is the golang structure for table bridge_config.
type BridgeConfig struct {	
         Id       int64         `orm:"id,primary" json:"id"`    //    
         SourceChainId    string         `orm:"source_chainId" json:"sourceChainId"`    // 当前链    
         TargetChainId    string         `orm:"target_chainId" json:"targetChainId"`    // 对手链    
         SourceCoinAddress    string         `orm:"source_coin_address" json:"sourceCoinAddress"`    // 币种    
         FeeFixed    float64         `orm:"fee_fixed" json:"feeFixed"`    // 固定金额手续费，    
         FeePercent    float64         `orm:"fee_percent" json:"feePercent"`    // 手续费百分比，如果不需要则设置为0    
         DayTotal    float64         `orm:"day_total" json:"dayTotal"`    // 跨入每日审核数量    
         OnceTotal    float64         `orm:"once_total" json:"onceTotal"`    // 跨入单次审核数量    
         IsEnable    int         `orm:"isEnable" json:"isEnable"`    // 状态    
         TargetCoinAddress    string         `orm:"target_coin_address" json:"targetCoinAddress"`    // 目标链合约地址    
         UpdateAt    *gtime.Time         `orm:"update_at" json:"updateAt"`    // 更新时间    
         CreateAt    *gtime.Time         `orm:"create_at" json:"createAt"`    // 创建时间    
}