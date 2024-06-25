// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-12 23:29:47
// 生成路径: gfast/app/admin/api/chain.go
// 生成人：jimmy
// ==========================================================================

package api

import (
	"gfast/app/admin/dao"
	"gfast/app/admin/service"
	sysApi "gfast/app/system/api"
	"gfast/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type chain struct {
	sysApi.SystemBase
}

var Chain = new(chain)

// List 列表
func (c *chain) List(r *ghttp.Request) {
	var req *dao.ChainSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.Chain.GetList(req)
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
func (c *chain) Add(r *ghttp.Request) {
	var req *dao.ChainAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	if !utils.VerifyAddress(req.BridgeContractAddress) {
		c.FailJsonExit(r, "Bridge Contract Address error")
	}
	err := service.Chain.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *chain) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.Chain.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *chain) Edit(r *ghttp.Request) {
	var req *dao.ChainEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.Chain.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *chain) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.Chain.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}
