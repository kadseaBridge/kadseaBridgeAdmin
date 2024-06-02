// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-05-30 18:31:15
// 生成路径: gfast/app/admin/dao/chain.go
// 生成人：jimmy
// ==========================================================================


package dao
import (
    comModel "gfast/app/common/model"
    "gfast/app/admin/dao/internal"    
    "github.com/gogf/gf/os/gtime"    
)
// chainDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type chainDao struct {
	*internal.ChainDao
}
var (
    // Chain is globally public accessible object for table tools_gen_table operations.
    Chain = chainDao{
        internal.NewChainDao(),
    }
)


// Fill with you ideas below.


// ChainSearchReq 分页请求参数
type ChainSearchReq struct {    
    Name  string `p:"name"` //链名称    
    ChainId  string `p:"chainId"` //链id    
    IsEnable  string `p:"isEnable"` //是否启用    
    comModel.PageReq
}
// ChainAddReq 添加操作请求参数
type ChainAddReq struct {    
    Name  string   `p:"name" v:"required#链名称不能为空"`    
    ChainId  string   `p:"chainId" v:"required#链id不能为空"`    
    Rpc  string   `p:"rpc" v:"required#rpc链接不能为空"`    
    Explorer  string   `p:"explorer" v:"required#浏览器链接不能为空"`    
    IsEnable  int   `p:"isEnable" v:"required#是否启用不能为空"`    
    CreateAt  *gtime.Time   `p:"createAt" `    
    UpdateAt  *gtime.Time   `p:"updateAt" `    
    ChainType  int   `p:"chainType" v:"required#0 ethereum 1 tron  2 btc不能为空"`    
    Icon  string   `p:"icon" `    
}
// ChainEditReq 修改操作请求参数
type ChainEditReq struct {
    Id    int64  `p:"id" v:"required#主键ID不能为空"`    
    Name  string `p:"name" v:"required#链名称不能为空"`    
    ChainId  string `p:"chainId" v:"required#链id不能为空"`    
    Rpc  string `p:"rpc" v:"required#rpc链接不能为空"`    
    Explorer  string `p:"explorer" v:"required#浏览器链接不能为空"`    
    IsEnable  int `p:"isEnable" v:"required#是否启用不能为空"`    
    CreateAt  *gtime.Time `p:"createAt" `    
    UpdateAt  *gtime.Time `p:"updateAt" `    
    ChainType  int `p:"chainType" v:"required#0 ethereum 1 tron  2 btc不能为空"`    
    Icon  string `p:"icon" `    
}
