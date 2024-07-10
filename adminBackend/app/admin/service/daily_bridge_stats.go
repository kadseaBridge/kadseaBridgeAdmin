// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-07-08 18:02:01
// 生成路径: gfast/app/admin/service/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package service

import (
	"context"
	"gfast/app/admin/dao"
	"gfast/app/admin/model"
	comModel "gfast/app/common/model"
	commonService "gfast/app/common/service"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"strings"
	"time"
)

type dailyBridgeStats struct {
}

var DailyBridgeStats = new(dailyBridgeStats)

// GetList 获取任务列表
func (s *dailyBridgeStats) GetList(req *dao.DailyBridgeStatsSearchReq) (total, page int, list []*model.DailyBridgeStats, err error) {
	m := dao.DailyBridgeStats.Ctx(req.Ctx)
	//if req.CoinAddress != "" {
	//	m = m.Where(dao.DailyBridgeStats.Columns.CoinAddress+" = ?", req.CoinAddress)
	//}

	if req.CoinName != "" {
		addressList, err := Coin.GetAddressBySymbol(req.Ctx, req.CoinName)
		if err != nil {
			err = gerror.New("获取coin address 失败")
		}
		m = m.Where(dao.DailyBridgeStats.Columns.CoinAddress+" IN(?)", addressList)
	}

	if req.ChainName != "" {

		chainId, err := Chain.GetIdByChainName(req.Ctx, req.ChainName)
		if err != nil {
			err = gerror.New("获取ChainId 失败")
		}
		m = m.Where(dao.DailyBridgeStats.Columns.ChainId+" = ?", chainId)
	}

	if req.StartDate != nil || req.EndDate != nil {
		if req.StartDate != nil && req.EndDate != nil {
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

// service
func (s *dailyBridgeStats) BatchAdd(ctx context.Context, reqs []*model.DailyBridgeStats) (err error) {
	if len(reqs) == 0 {
		return nil
	}
	// 使用 Batch 插入数据
	_, err = dao.DailyBridgeStats.Ctx(ctx).Data(reqs).Batch(len(reqs)).Insert()
	return err
}

func (s *dailyBridgeStats) DailyStats(req *dao.DailyBridgeStatsSearchReq, date time.Time) {
	var orders []model.BridgeOrder
	startDate := gtime.NewFromTime(date.Truncate(24 * time.Hour))
	endDate := gtime.NewFromTime(date.Truncate(24 * time.Hour).Add(24 * time.Hour))

	m := dao.BridgeOrder.Ctx(req.Ctx)

	m = m.Where(dao.BridgeOrder.Columns.UpdateAt+" BETWEEN ? AND ? AND "+dao.BridgeOrder.Columns.Status+" = ?", startDate, endDate, 5)

	err := m.Scan(&orders)
	if err != nil {
		g.Log().Error("获取订单数据失败:", err)
	}

	statsMap := make(map[string]*model.DailyBridgeStats)
	for _, order := range orders {
		sourceKey := order.SourceChainId + "_" + order.SourceCoinAddress
		sourceStats := statsMap[sourceKey]
		if sourceStats == nil {
			sourceStats = &model.DailyBridgeStats{
				ChainId:     order.SourceChainId,
				CoinAddress: order.SourceCoinAddress,
				Date:        startDate,
			}
			statsMap[sourceKey] = sourceStats
		}
		sourceStats.TransferOut += order.Amount - order.Fee
		sourceStats.Fee += order.Fee

		targetKey := order.TargetChainId + "_" + order.TargetCoinAddress
		targetStats := statsMap[targetKey]
		if targetStats == nil {

			targetStats = &model.DailyBridgeStats{
				ChainId:     order.TargetChainId,
				CoinAddress: order.TargetCoinAddress,
				Date:        startDate,
			}
			statsMap[targetKey] = targetStats
		}
		targetStats.TransferIn += order.Amount
	}

	stats := make([]*model.DailyBridgeStats, 0, len(statsMap))
	for _, stat := range statsMap {
		stat.TransferDifference = stat.TransferIn - stat.TransferOut
		stat.PlatformAssets = getBalance(stat.ChainId, stat.CoinAddress)
		stats = append(stats, stat)
	}

	err = DailyBridgeStats.BatchAdd(req.Ctx, stats)
	if err != nil {
		g.Log().Error(err)
	}

}

func getBalance(chainId, tokenAddress string) float64 {

	// tron 的财务地址
	var tronAccount = "TRwntfF9HM4W2M4neaheG1Syu8eyjZEijm"
	// evm的财务地址
	var evmAccount = "0x2b83877aCE845279f59919aeb912946C8b5Abe26"
	rpc, err := Chain.GetRpcByChainId(context.Background(), chainId)
	if err != nil {
		g.Log().Error("获取rpc失败:", err)
		return 0
	}

	if strings.Contains(rpc, "tron") {
		bal, err := commonService.NewRpcUtil().GetTronBalance(rpc, tokenAddress, tronAccount)

		if err != nil {
			g.Log().Error("获取Tron余额失败:", err)
			return 0
		}
		return bal

	} else {
		bal, err := commonService.NewRpcUtil().GetEvmBalance(rpc, tokenAddress, evmAccount)
		if err != nil {
			g.Log().Error("获取Evm余额失败:", err)
			return 0
		}
		return bal
	}

}
