// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: gfast/app/admin/api/bridge_config.go
// 生成人：jimmy
// ==========================================================================

package api

import (
	"fmt"
	"gfast/app/admin/dao"
	"gfast/app/admin/service"
	sysApi "gfast/app/system/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type bridgeConfig struct {
	sysApi.SystemBase
}

var BridgeConfig = new(bridgeConfig)

// List 列表
func (c *bridgeConfig) List(r *ghttp.Request) {
	var req *dao.BridgeConfigSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.BridgeConfig.GetList(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	c.SusJsonExit(r, result)
}

// Add 添加
func (c *bridgeConfig) Add(r *ghttp.Request) {
	var req *dao.BridgeConfigAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	err := service.BridgeConfig.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *bridgeConfig) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.BridgeConfig.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *bridgeConfig) Edit(r *ghttp.Request) {
	var req *dao.BridgeConfigEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.BridgeConfig.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *bridgeConfig) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.BridgeConfig.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

// Export
func (c *bridgeConfig) Export(r *ghttp.Request) {
	var req *dao.BridgeConfigSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	data, err := service.BridgeConfig.ExportBridgeConfigs(req)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{"error": err.Error()})
	}

	// 设置响应头以指示文件下载
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Content-Disposition", `attachment; filename="coins.xlsx"`)
	r.Response.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

	// 返回文件内容给前端
	r.Response.Write(data)

}
