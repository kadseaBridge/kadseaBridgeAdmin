// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: gfast/app/admin/service/bridge_config.go
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
	"strings"
)

type bridgeConfig struct {
}

var BridgeConfig = new(bridgeConfig)

// GetList 获取任务列表
func (s *bridgeConfig) GetList(req *dao.BridgeConfigSearchReq) (total, page int, list []*model.BridgeConfig, err error) {
	m := dao.BridgeConfig.Ctx(req.Ctx)
	if req.SourceChainId != "" {
		m = m.Where(dao.BridgeConfig.Columns.SourceChainId+" = ?", req.SourceChainId)
	}
	if req.SourceCoinAddress != "" {
		m = m.Where(dao.BridgeConfig.Columns.SourceCoinAddress+" = ?", req.SourceCoinAddress)
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
func (s *bridgeConfig) GetInfoById(ctx context.Context, id int64) (info *model.BridgeConfig, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.BridgeConfig.Ctx(ctx).Where(dao.BridgeConfig.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *bridgeConfig) Add(ctx context.Context, req *dao.BridgeConfigAddReq) (err error) {

	if len(req.TargetCoinAddress) != 42 {
		req.TargetCoinAddress, err = Coin.GetAddressByNameAndChainId(ctx, req.SourceCoinAddress, req.TargetChainId)
		if err != nil {
			err = gerror.New("目标链地址失败")
			return
		}
	}

	if len(req.SourceCoinAddress) != 42 {
		req.SourceCoinAddress, err = Coin.GetAddressByNameAndChainId(ctx, req.SourceCoinAddress, req.SourceChainId)
		if err != nil {
			err = gerror.New("获取当前链地址失败")
			return
		}
	}

	if len(req.TargetCoinAddress) != 42 || len(req.SourceCoinAddress) != 42 {
		err = gerror.New("targetCoinAddress 或者 sourceCoinAddress 错误！ ")
		return
	}

	_, err = dao.BridgeConfig.Ctx(ctx).Insert(req)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("不能添加重复数据")
		}
		return err
	}
	return
}

// Edit 修改
func (s *bridgeConfig) Edit(ctx context.Context, req *dao.BridgeConfigEditReq) (err error) {

	_, err = dao.BridgeConfig.Ctx(ctx).FieldsEx(dao.BridgeConfig.Columns.Id).Where(dao.BridgeConfig.Columns.Id, req.Id).Update(req)
	return err
}

// DeleteByIds 删除
func (s *bridgeConfig) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.BridgeConfig.Ctx(ctx).Delete(dao.BridgeConfig.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

func (s *bridgeConfig) ExportBridgeConfigs(req *dao.BridgeConfigSearchReq) ([]byte, error) {
	// 获取数据库连接
	_, _, list, err := s.GetList(req)
	if err != nil {
		err = gerror.New("参数错误")
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	// 设置表头
	headers := []string{"序号", "当前链", "币种", "对手链", "跨入手续费", "跨入每日审核数量", "跨入单次审核数量", "状态"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	// 填充数据
	for i, bridgeConfig := range list {
		coinName, err := Coin.GetNameByAddress(req.Ctx, bridgeConfig.SourceCoinAddress)
		if err != nil {
			err = gerror.New("获取币种名称失败")
		}

		SourceChainName, err := Chain.GetNameByChainId(req.Ctx, bridgeConfig.SourceChainId)
		if err != nil {
			err = gerror.New("获取链名称名称失败")
		}

		TargetChainName, err := Chain.GetNameByChainId(req.Ctx, bridgeConfig.TargetChainId)
		if err != nil {
			err = gerror.New("获取链名称名称失败")
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), bridgeConfig.Id)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), SourceChainName)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), coinName)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), TargetChainName)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), bridgeConfig.FeeFixed+bridgeConfig.FeePercent)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), bridgeConfig.DayTotal)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), bridgeConfig.OnceTotal)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), bridgeConfig.IsEnable)

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
