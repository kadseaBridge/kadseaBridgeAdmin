// ==========================================================================
// GFast自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-12 23:29:47
// 生成路径: adminBackend/app/admin/router/chain.go
// 生成人：jimmy
// ==========================================================================

package router

import (
	"gfast/app/admin/api"
	sysApi "gfast/app/system/api"
	"gfast/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 加载路由
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			group.Group("/chain", func(group *ghttp.RouterGroup) {
				//gToken拦截器
				sysApi.GfToken.AuthMiddleware(group)
				//context拦截器
				group.Middleware(middleware.Ctx, middleware.Auth)
				//后台操作日志记录
				group.Hook("/*", ghttp.HookAfterOutput, sysApi.SysOperLog.OperationLog)
				group.GET("list", api.Chain.List)
				group.GET("get", api.Chain.Get)
				group.POST("add", api.Chain.Add)
				group.PUT("edit", api.Chain.Edit)
				group.DELETE("delete", api.Chain.Delete)
			})
		})
	})
}
