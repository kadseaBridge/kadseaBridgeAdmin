// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-08 23:23:41
// 生成路径: gfast/app/admin/service/pay_detail.go
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

type payDetail struct {
}

var PayDetail = new(payDetail)

// GetList 获取任务列表
func (s *payDetail) GetList(req *dao.PayDetailSearchReq) (total, page int, list []*model.PayDetail, err error) {
	m := dao.PayDetail.Ctx(req.Ctx)
	if req.OrderId != "" {
		m = m.Where(dao.PayDetail.Columns.OrderId+" = ?", req.OrderId)
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
func (s *payDetail) GetInfoById(ctx context.Context, id int64) (info *model.PayDetail, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.PayDetail.Ctx(ctx).Where(dao.PayDetail.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *payDetail) Add(ctx context.Context, req *dao.PayDetailAddReq) (err error) {
	_, err = dao.PayDetail.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *payDetail) Edit(ctx context.Context, req *dao.PayDetailEditReq) error {
	_, err := dao.PayDetail.Ctx(ctx).FieldsEx(dao.PayDetail.Columns.Id).Where(dao.PayDetail.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *payDetail) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.PayDetail.Ctx(ctx).Delete(dao.PayDetail.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// GetInfoById 通过id获取
func (s *payDetail) GetGasAndTxByOrderId(ctx context.Context, orderId string) (tx string, gas float64, err error) {
	if orderId == "" {
		err = gerror.New("参数错误")
		return
	}

	// 创建一个临时结构体来接收只包含 Name 字段的数据
	var result struct {
		TransactionHash string
		PayGasFee       float64
	}

	// 只选择 Name 字段
	err = dao.PayDetail.Ctx(ctx).Where(dao.PayDetail.Columns.OrderId, orderId).Fields("transaction_hash", "pay_gas_fee").Scan(&result)
	if err != nil {
		g.Log().Error(err)
		return
	}
	//if result.TransactionHash == "" ||  {
	//	err = gerror.New("获取信息失败")
	//	return
	//}

	tx = result.TransactionHash
	gas = result.PayGasFee
	return
}

// GetInfoById 通过id获取
func (s *payDetail) GetOrderIdByTx(ctx context.Context, tx string) (orderId string, err error) {
	if tx == "" {
		err = gerror.New("参数错误")
		return
	}

	//m := dao.BridgeOrder.Ctx(ctx)
	//
	//m = m.Where(dao.PayDetail.Columns.TransactionHash+" LIKE ?", "%"+tx+"%")

	// 创建一个临时结构体来接收只包含 Name 字段的数据
	var result struct {
		OrderId string
	}

	// 只选择 Name 字段
	//err = dao.PayDetail.Ctx(ctx).Where(dao.PayDetail.Columns.TransactionHash, tx).Fields("order_id").Scan(&result)
	err = dao.PayDetail.Ctx(ctx).Where(dao.PayDetail.Columns.TransactionHash+" LIKE ?", "%"+tx+"%").Fields("order_id").Scan(&result)
	if err != nil {
		g.Log().Error(err)
		return
	}
	if result.OrderId == "" {
		err = gerror.New("获取OrderId失败")
		return
	}

	orderId = result.OrderId
	return
}
