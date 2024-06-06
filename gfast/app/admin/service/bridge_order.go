// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: gfast/app/admin/service/bridge_order.go
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

type bridgeOrder struct {
}

var BridgeOrder = new(bridgeOrder)

// GetList 获取任务列表
func (s *bridgeOrder) GetList(req *dao.BridgeOrderSearchReq) (total, page int, list []*model.BridgeOrder, err error) {
	m := dao.BridgeOrder.Ctx(req.Ctx)
	if req.SourceAddress != "" {
		m = m.Where(dao.BridgeOrder.Columns.SourceAddress+" = ?", req.SourceAddress)
	}
	if req.TargetAddress != "" {
		m = m.Where(dao.BridgeOrder.Columns.TargetAddress+" = ?", req.TargetAddress)
	}
	if req.SourceCoinAddress != "" {
		m = m.Where(dao.BridgeOrder.Columns.SourceCoinAddress+" = ?", req.SourceCoinAddress)
	}
	if req.TargetCoinAddress != "" {
		m = m.Where(dao.BridgeOrder.Columns.TargetCoinAddress+" = ?", req.TargetCoinAddress)
	}
	if req.SourceChainId != "" {
		m = m.Where(dao.BridgeOrder.Columns.SourceChainId+" = ?", req.SourceChainId)
	}
	if req.TargetChainId != "" {
		m = m.Where(dao.BridgeOrder.Columns.TargetChainId+" = ?", req.TargetChainId)
	}
	if req.TransactionHash != "" {
		m = m.Where(dao.BridgeOrder.Columns.TransactionHash+" = ?", req.TransactionHash)
	}
	if req.Status != "" {
		m = m.Where(dao.BridgeOrder.Columns.Status+" = ?", req.Status)
	}
	if len(req.StartRange) == 2 {
		m = m.Where(dao.BridgeOrder.Columns.CreateAt+" BETWEEN ? AND ?", req.StartRange[0], req.StartRange[1])
	}
	if len(req.EndRange) == 2 {
		m = m.Where(dao.BridgeOrder.Columns.CreateAt+"status = 5 AND BETWEEN ? AND ?", req.EndRange[0], req.EndRange[1])
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
func (s *bridgeOrder) GetInfoById(ctx context.Context, id int64) (info *model.BridgeOrder, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.BridgeOrder.Ctx(ctx).Where(dao.BridgeOrder.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *bridgeOrder) Add(ctx context.Context, req *dao.BridgeOrderAddReq) (err error) {
	_, err = dao.BridgeOrder.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *bridgeOrder) Edit(ctx context.Context, req *dao.BridgeOrderEditReq) error {
	_, err := dao.BridgeOrder.Ctx(ctx).FieldsEx(dao.BridgeOrder.Columns.Id).Where(dao.BridgeOrder.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *bridgeOrder) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.BridgeOrder.Ctx(ctx).Delete(dao.BridgeOrder.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// ChangeStatus 修改状态
func (s *bridgeOrder) ChangeStatus(ctx context.Context, req *dao.BridgeOrderStatusReq) error {
	_, err := dao.BridgeOrder.Ctx(ctx).WherePri(req.Id).Update(g.Map{
		dao.BridgeOrder.Columns.Status: req.Status,
	})
	return err
}
