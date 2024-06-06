// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: gfast/app/admin/service/coin.go
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

type coin struct {
}

var Coin = new(coin)

// GetList 获取任务列表
func (s *coin) GetList(req *dao.CoinSearchReq) (total, page int, list []*model.Coin, err error) {
	m := dao.Coin.Ctx(req.Ctx)
	if req.Symbol != "" {
		m = m.Where(dao.Coin.Columns.Symbol+" = ?", req.Symbol)
	}
	if req.ChainId != "" {
		m = m.Where(dao.Coin.Columns.ChainId+" = ?", req.ChainId)
	}
	if req.Address != "" {
		m = m.Where(dao.Coin.Columns.Address+" = ?", req.Address)
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
func (s *coin) GetInfoById(ctx context.Context, id int64) (info *model.Coin, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Coin.Ctx(ctx).Where(dao.Coin.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *coin) Add(ctx context.Context, req *dao.CoinAddReq) (err error) {
	_, err = dao.Coin.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *coin) Edit(ctx context.Context, req *dao.CoinEditReq) error {
	_, err := dao.Coin.Ctx(ctx).FieldsEx(dao.Coin.Columns.Id).Where(dao.Coin.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *coin) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Coin.Ctx(ctx).Delete(dao.Coin.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

func (s *coin) ExportCoins(req *dao.CoinSearchReq) ([]byte, error) {
	// 获取数据库连接
	_, _, list, err := s.GetList(req)
	if err != nil {
		err = gerror.New("参数错误")
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	// 设置表头
	headers := []string{"Id", "Name", "Symbol", "ChainId", "Address", "IsEnable", "CreateAt", "UpdateAt", "TokenType", "Decimals", "Icon"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	// 填充数据
	for i, coin := range list {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), coin.Id)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), coin.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), coin.Symbol)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), coin.ChainId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), coin.Address)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), coin.IsEnable)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), coin.CreateAt.String())
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), coin.UpdateAt.String())
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), coin.TokenType)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), coin.Decimals)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2), coin.Icon)
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
