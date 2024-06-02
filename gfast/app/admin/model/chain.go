// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:15
// 生成路径: gfast/app/admin/model/chain.go
// 生成人：jimmy
// ==========================================================================


package model
import "github.com/gogf/gf/os/gtime"
// Chain is the golang structure for table chain.
type Chain struct {	
         Id       int64         `orm:"id,primary" json:"id"`    //    
         Name    string         `orm:"name" json:"name"`    // 链名称    
         ChainId    string         `orm:"chainId" json:"chainId"`    // 链id    
         Rpc    string         `orm:"rpc" json:"rpc"`    // rpc链接    
         Explorer    string         `orm:"explorer" json:"explorer"`    // 浏览器链接    
         IsEnable    int         `orm:"isEnable" json:"isEnable"`    // 是否启用    
         CreateAt    *gtime.Time         `orm:"create_at" json:"createAt"`    //    
         UpdateAt    *gtime.Time         `orm:"update_at" json:"updateAt"`    //    
         ChainType    int         `orm:"type" json:"chainType"`    // 0 ethereum 1 tron  2 btc    
         Icon    string         `orm:"icon" json:"icon"`    // 图标    
}