// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-14 17:06:57
// 生成路径: adminBackend/app/admin/service/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package service

import (
	"context"
	"gfast/app/admin/dao"
	"gfast/app/admin/model"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type dailyBridgeStats struct {
}

var DailyBridgeStats = new(dailyBridgeStats)

// GetList 获取任务列表
func (s *dailyBridgeStats) GetList(req *dao.DailyBridgeStatsSearchReq) (total, page int, list []*model.DailyBridgeStats, err error) {
	m := dao.DailyBridgeStats.Ctx(req.Ctx)
	if req.Coin != "" {
		m = m.Where(dao.DailyBridgeStats.Columns.Coin+" like ?", "%"+req.Coin+"%")
	}
	if req.ChainType != "" {
		m = m.Where(dao.DailyBridgeStats.Columns.ChainType+" like ?", "%"+req.ChainType+"%")
	}

	if req.EndDate != nil || req.StartDate != nil {
		if req.EndDate != nil && req.StartDate != nil {
			m = m.Where(dao.DailyBridgeStats.Columns.Date+" BETWEEN ? AND ?", req.StartDate, req.EndDate)
		} else if req.StartDate != nil {
			m = m.Where(dao.DailyBridgeStats.Columns.Date+"= ?", req.StartDate)
		} else if req.EndDate != nil {
			m = m.Where(dao.DailyBridgeStats.Columns.Date+"= ?", req.EndDate)
		}

	}

	total, err = m.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	page = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = comModel.PageSize
	}
	order := "id asc"
	if req.OrderBy != "" {
		order = req.OrderBy
	}
	err = m.Page(page, req.PageSize).Order(order).Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}

// GetInfoById 通过id获取
func (s *dailyBridgeStats) GetInfoById(ctx context.Context, id int) (info *model.DailyBridgeStats, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.DailyBridgeStats.Ctx(ctx).Where(dao.DailyBridgeStats.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *dailyBridgeStats) Add(ctx context.Context, req *dao.DailyBridgeStatsAddReq) (err error) {
	_, err = dao.DailyBridgeStats.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *dailyBridgeStats) Edit(ctx context.Context, req *dao.DailyBridgeStatsEditReq) error {
	_, err := dao.DailyBridgeStats.Ctx(ctx).FieldsEx(dao.DailyBridgeStats.Columns.Id, dao.DailyBridgeStats.Columns.CreatedAt).Where(dao.DailyBridgeStats.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *dailyBridgeStats) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.DailyBridgeStats.Ctx(ctx).Delete(dao.DailyBridgeStats.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}
