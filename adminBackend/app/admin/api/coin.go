// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: adminBackend/app/admin/api/coin.go
// 生成人：jimmy
// ==========================================================================

package api

import (
	"fmt"
	"gfast/app/admin/dao"
	"gfast/app/admin/service"
	sysApi "gfast/app/system/api"
	"gfast/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"strings"
)

type coin struct {
	sysApi.SystemBase
}

var Coin = new(coin)

// List 列表
func (c *coin) List(r *ghttp.Request) {
	var req *dao.CoinSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.Coin.GetList(req)
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
func (c *coin) Add(r *ghttp.Request) {
	var req *dao.CoinAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	if strings.Contains(req.ChainId, "tron") {
		if len(req.Address) != 34 {
			c.FailJsonExit(r, "Coin Address error")
		}
	} else {
		if !utils.VerifyAddress(req.Address) {
			c.FailJsonExit(r, "Coin Address error")
		}
	}

	// 所有新添加的币种都是上架状态
	req.IsEnable = 1
	err := service.Coin.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *coin) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.Coin.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *coin) Edit(r *ghttp.Request) {
	var req *dao.CoinEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if strings.Contains(req.ChainId, "tron") {
		if len(req.Address) != 34 {
			c.FailJsonExit(r, "Coin Address error")
		}
	} else {
		if !utils.VerifyAddress(req.Address) {
			c.FailJsonExit(r, "Coin Address error")
		}
	}

	err := service.Coin.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *coin) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.Coin.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

// Export

func (c *coin) Export(r *ghttp.Request) {
	var req *dao.CoinSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	data, err := service.Coin.ExportCoins(req)
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
