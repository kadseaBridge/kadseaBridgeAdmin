// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: adminBackend/app/admin/api/bridge_order.go
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

type bridgeOrder struct {
	sysApi.SystemBase
}

var BridgeOrder = new(bridgeOrder)

// List 列表
func (c *bridgeOrder) List(r *ghttp.Request) {
	var req *dao.BridgeOrderSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.BridgeOrder.GetList(req)
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
func (c *bridgeOrder) Add(r *ghttp.Request) {
	var req *dao.BridgeOrderAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.BridgeOrder.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *bridgeOrder) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.BridgeOrder.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *bridgeOrder) Edit(r *ghttp.Request) {
	var req *dao.BridgeOrderEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.BridgeOrder.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *bridgeOrder) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.BridgeOrder.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

// ChangeStatus 修改状态
func (c *bridgeOrder) ChangeStatus(r *ghttp.Request) {
	var req *dao.BridgeOrderStatusReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if err := service.BridgeOrder.ChangeStatus(r.GetCtx(), req); err != nil {
		c.FailJsonExit(r, err.Error())
	} else {
		c.SusJsonExit(r, "状态设置成功")
	}
}

func (c *bridgeOrder) Export(r *ghttp.Request) {
	var req *dao.BridgeOrderSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	data, err := service.BridgeOrder.ExportOrders(req)
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
