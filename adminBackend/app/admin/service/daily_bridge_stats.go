// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-04 17:54:12
// 生成路径: gfast/app/admin/service/daily_bridge_stats.go
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
	"github.com/gogf/gf/os/gtime"
	"time"
)

type dailyBridgeStats struct {
}

var DailyBridgeStats = new(dailyBridgeStats)

// GetList 获取任务列表
func (s *dailyBridgeStats) GetList(req *dao.DailyBridgeStatsSearchReq) (total, page int, list []*model.DailyBridgeStats, err error) {
	m := dao.DailyBridgeStats.Ctx(req.Ctx)
	if req.CoinName != "" {
		m = m.Where(dao.DailyBridgeStats.Columns.CoinName+" like ?", "%"+req.CoinName+"%")
	}
	if req.ChainName != "" {
		m = m.Where(dao.DailyBridgeStats.Columns.ChainName+" like ?", "%"+req.ChainName+"%")
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
	order := "id desc"
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

func (s *dailyBridgeStats) BatchAdd(ctx context.Context, reqs []*model.DailyBridgeStats) (err error) {
	if len(reqs) == 0 {
		return nil
	}
	// 使用 Batch 插入数据
	_, err = dao.DailyBridgeStats.Ctx(ctx).Data(reqs).Batch(len(reqs)).Insert()
	return err
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

func (s *dailyBridgeStats) DailyStats(req *dao.DailyBridgeStatsSearchReq, date time.Time) {
	var orders []model.BridgeOrder
	startDate := gtime.NewFromTime(date.Truncate(24 * time.Hour))
	endDate := gtime.NewFromTime(date.Truncate(24 * time.Hour).Add(24 * time.Hour))

	m := dao.BridgeOrder.Ctx(req.Ctx)

	m = m.Where(dao.BridgeOrder.Columns.UpdateAt+" BETWEEN ? AND ? AND "+dao.BridgeOrder.Columns.Status+" = ?", startDate, endDate, 5)

	err := m.Scan(&orders)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}

	statsMap := make(map[string]*model.DailyBridgeStats)
	for _, order := range orders {
		sourceKey := order.SourceChainId + "_" + order.SourceCoinAddress
		sourceStats := statsMap[sourceKey]
		if sourceStats == nil {
			sourceChainName, err := Chain.GetNameByChainId(req.Ctx, order.SourceChainId)
			if err != nil {
				err = gerror.New("获取chainName失败")
				g.Log().Error(err)

			}
			sourceCoinName, err := Coin.GetNameByAddress(req.Ctx, order.SourceCoinAddress)

			if err != nil {
				err = gerror.New("获取coinName失败")
				g.Log().Error(err)

			}
			sourceStats = &model.DailyBridgeStats{
				ChainName: sourceChainName,
				CoinName:  sourceCoinName,
				Date:      startDate,
			}
			statsMap[sourceKey] = sourceStats
		}
		sourceStats.TransferOut += order.Amount - order.Fee
		sourceStats.Fee += order.Fee

		targetKey := order.TargetChainId + "_" + order.TargetCoinAddress
		targetStats := statsMap[targetKey]
		if targetStats == nil {

			targetChainName, err := Chain.GetNameByChainId(req.Ctx, order.TargetChainId)
			if err != nil {
				err = gerror.New("获取chainName失败")
				g.Log().Error(err)

			}
			targetCoinName, err := Coin.GetNameByAddress(req.Ctx, order.TargetCoinAddress)

			if err != nil {
				err = gerror.New("获取coinName失败")
				g.Log().Error(err)

			}

			targetStats = &model.DailyBridgeStats{
				ChainName: targetChainName,
				CoinName:  targetCoinName,
				Date:      startDate,
			}
			statsMap[targetKey] = targetStats
		}
		targetStats.TransferIn += order.Amount
	}

	stats := make([]*model.DailyBridgeStats, 0, len(statsMap))
	for _, stat := range statsMap {
		stat.TransferDifference = stat.TransferIn - stat.TransferOut
		stats = append(stats, stat)
	}

	err = DailyBridgeStats.BatchAdd(req.Ctx, stats)
	if err != nil {
		g.Log().Error(err)
	}

}
