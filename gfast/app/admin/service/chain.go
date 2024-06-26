// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-06-12 23:29:47
// 生成路径: gfast/app/admin/service/chain.go
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

type chain struct {
}

var Chain = new(chain)

// GetList 获取任务列表
func (s *chain) GetList(req *dao.ChainSearchReq) (total, page int, list []*model.Chain, err error) {
	m := dao.Chain.Ctx(req.Ctx)
	if req.Name != "" {
		m = m.Where(dao.Chain.Columns.Name+" like ?", "%"+req.Name+"%")
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
func (s *chain) GetInfoById(ctx context.Context, id int64) (info *model.Chain, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.Chain.Ctx(ctx).Where(dao.Chain.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *chain) Add(ctx context.Context, req *dao.ChainAddReq) (err error) {
	_, err = dao.Chain.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *chain) Edit(ctx context.Context, req *dao.ChainEditReq) error {
	_, err := dao.Chain.Ctx(ctx).FieldsEx(dao.Chain.Columns.Id).Where(dao.Chain.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *chain) DeleteByIds(ctx context.Context, ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}
	_, err = dao.Chain.Ctx(ctx).Delete(dao.Chain.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

// GetInfoById 通过id获取
func (s *chain) GetNameByChainId(ctx context.Context, chainId string) (name string, err error) {
	if chainId == "" {
		err = gerror.New("参数错误")
		return
	}

	// 创建一个临时结构体来接收只包含 Name 字段的数据
	var result struct {
		Name string
	}

	// 只选择 Name 字段
	err = dao.Chain.Ctx(ctx).Where(dao.Chain.Columns.ChainId, chainId).Fields("name").Scan(&result)
	if err != nil {
		g.Log().Error(err)
		return
	}
	if result.Name == "" {
		err = gerror.New("获取信息失败")
		return
	}

	name = result.Name
	return
}
