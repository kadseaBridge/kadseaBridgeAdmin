// ==========================================================================
// GFast自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-27 18:48:21
// 生成路径: gfast/app/work/router/bridge_config.go
// 生成人：gfast
// ==========================================================================


package router
import (
    "gfast/app/work/api"
    "gfast/middleware"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"    
    sysApi "gfast/app/system/api"    
)
//加载路由
func init() {
    s := g.Server()
    s.Group("/", func(group *ghttp.RouterGroup) {
        group.Group("/work", func(group *ghttp.RouterGroup) {
            group.Group("/bridgeConfig", func(group *ghttp.RouterGroup) {
                //gToken拦截器                
                sysApi.GfToken.AuthMiddleware(group)                
                //context拦截器
                group.Middleware(middleware.Ctx, middleware.Auth)                
                //后台操作日志记录
                group.Hook("/*", ghttp.HookAfterOutput, sysApi.SysOperLog.OperationLog)                
                group.GET("list", api.BridgeConfig.List)
                group.GET("get", api.BridgeConfig.Get)
                group.POST("add", api.BridgeConfig.Add)
                group.PUT("edit", api.BridgeConfig.Edit)
                group.DELETE("delete", api.BridgeConfig.Delete)                
            })
        })
    })
}
