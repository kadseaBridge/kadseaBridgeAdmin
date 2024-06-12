// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: gfast/app/admin/service/bridge_order.go
// 生成人：jimmy
// ==========================================================================

package service

import (
	"bytes"
	"context"
	"fmt"
	"gfast/app/admin/dao"
	"gfast/app/admin/model"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/xuri/excelize/v2"
)

type bridgeOrder struct {
}

var BridgeOrder = new(bridgeOrder)

// GetList 获取任务列表
func (s *bridgeOrder) GetList(req *dao.BridgeOrderSearchReq) (total, page int, orderlist []*model.BridgeOrderRsp, err error) {
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

	var list []*model.BridgeOrder
	err = m.Page(page, req.PageSize).Order(order).Scan(&list)

	rsp := make([]*model.BridgeOrderRsp, len(list))

	for i, order := range list {

		tx, gas, err := PayDetail.GetGasAndTxByOrderId(req.Ctx, order.OrderId)
		if err != nil {
			err = gerror.New("获取链名称名称失败")
		}
		sourceCoinName, err := Coin.GetNameByAddress(req.Ctx, order.SourceCoinAddress)
		if err != nil {
			err = gerror.New("获取币种名称失败")
		}

		targetCoinName, err := Coin.GetNameByAddress(req.Ctx, order.SourceCoinAddress)
		if err != nil {
			err = gerror.New("获取币种名称失败")
		}

		sourceChainName, err := Chain.GetNameByChainId(req.Ctx, order.SourceChainId)
		if err != nil {
			err = gerror.New("获取链名称名称失败")
		}

		targetChainName, err := Chain.GetNameByChainId(req.Ctx, order.TargetChainId)
		if err != nil {
			err = gerror.New("获取链名称名称失败")
		}
		rsp[i] = &model.BridgeOrderRsp{
			Id:              order.Id,
			SourceAddress:   order.SourceAddress,
			TargetAddress:   order.TargetAddress,
			SourceCoinName:  sourceCoinName,
			TargetCoinName:  targetCoinName,
			SourceChainName: sourceChainName,
			TargetChainName: targetChainName,
			InHash:          tx,
			OutHash:         order.TransactionHash,
			OrderId:         order.OrderId,
			Amount:          order.Amount,
			Status:          order.Status,
			CreateAt:        order.CreateAt,
			UpdateAt:        order.UpdateAt,
			BlockNumber:     order.BlockNumber,
			Fee:             order.Fee,
			GasFee:          gas,
			Remark:          order.Remark,
		}
	}

	orderlist = rsp

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

func (s *bridgeOrder) ExportOrders(req *dao.BridgeOrderSearchReq) ([]byte, error) {
	// 获取数据库连接
	_, _, list, err := s.GetList(req)
	if err != nil {
		err = gerror.New("参数错误")
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	// 设置表头
	headers := []string{"序号", "转出币种", "转入币种", "转出链", "转入链", "数量", "手续费", "gas费", "转入钱包地址", "转出钱包地址", "转出TXID", "转入TXID", "发起时间", "结束时间", "状态"}

	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	// 填充数据
	for i, order := range list {

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), order.Id)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), order.SourceCoinName)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), order.TargetCoinName)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), order.SourceChainName)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), order.TargetChainName)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), order.Amount)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), order.Fee)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), order.GasFee)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), order.SourceAddress)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), order.TargetAddress)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2), order.OutHash)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", i+2), order.InHash)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+2), order.CreateAt)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+2), order.UpdateAt)
		f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+2), order.Status)
	}

	// 将 Excel 文件写入缓冲区
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	// 添加日志输出
	g.Log().Info("Excel file created successfully")

	return buf.Bytes(), nil
}
