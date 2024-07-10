// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-08 18:02:01
// 生成路径: gfast/app/admin/api/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package api

import (
	"fmt"
	"gfast/app/admin/dao"
	"gfast/app/admin/service"
	sysApi "gfast/app/system/api"
	systemService "gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"regexp"
	"time"
)

type dailyBridgeStats struct {
	sysApi.SystemBase
}

var DailyBridgeStats = new(dailyBridgeStats)

// List 列表
func (c *dailyBridgeStats) List(r *ghttp.Request) {
	var req *dao.DailyBridgeStatsSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.DailyBridgeStats.GetList(req)
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
func (c *dailyBridgeStats) Add(r *ghttp.Request) {
	var req *dao.DailyBridgeStatsAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.DailyBridgeStats.Add(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *dailyBridgeStats) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	info, err := service.DailyBridgeStats.GetInfoById(r.GetCtx(), id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *dailyBridgeStats) Edit(r *ghttp.Request) {
	var req *dao.DailyBridgeStatsEditReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.DailyBridgeStats.Edit(r.GetCtx(), req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *dailyBridgeStats) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.DailyBridgeStats.DeleteByIds(r.GetCtx(), ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

// 获取某一天的跨链统计
func (c *dailyBridgeStats) DailyStats() {
	param := &dao.DailyBridgeStatsSearchReq{}
	//
	now := time.Now().AddDate(0, 0, -1)
	// 格式化当前时间为 2024-07-09 00:00:00 +0000 UTC
	layout := "2006-01-02"
	formattedTime := now.Format(layout)
	// 解析格式化后的时间为 time.Time 类型
	startDate, err := time.Parse(layout, formattedTime)
	if err != nil {
		fmt.Println("解析时间时出错:", err)
		return
	}
	service.DailyBridgeStats.DailyStats(param, startDate)

}

func (c *dailyBridgeStats) StartDailyStats() {
	param := &dao.DailyBridgeStatsSearchReq{}

	t := systemService.TimeTaskList.GetByName("startDailyStats")
	if t == nil {
		return
	}
	inputDate := t.Param[0]
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	dateRegex, err := regexp.Compile(datePattern)
	if err != nil {
		g.Log().Error("Error compiling regex::", err)
		return
	}

	//testDate := "2024-07-09"
	if !dateRegex.MatchString(inputDate) {
		fmt.Println("The date format is invalid.")
		g.Log().Error("定时任务输入的日期格式错误:")
		return
	}

	layout := "2006-01-02"
	//inputDate := "2024-07-09"
	startDate, err := time.Parse(layout, inputDate)
	if err != nil {
		g.Log().Error("日期格式错误:", err)
		return
	}
	// 获取当前日期
	endDate := time.Now().Truncate(24 * time.Hour)

	for date := startDate; !date.After(endDate); date = date.Add(24 * time.Hour) {
		service.DailyBridgeStats.DailyStats(param, date)

	}

}
