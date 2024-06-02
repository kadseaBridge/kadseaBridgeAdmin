// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: gfast/app/admin/service/coin.go
// 生成人：jimmy
// ==========================================================================


package service
import (
    "context"
	comModel "gfast/app/common/model"
	"gfast/app/admin/dao"
	"gfast/app/admin/model"	
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
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
    order:= "id asc"
    if req.OrderBy!=""{
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
func (s *coin) GetInfoById(ctx context.Context,id int64) (info *model.Coin, err error) {
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
func (s *coin) Add(ctx context.Context,req *dao.CoinAddReq) (err error) {
	_, err = dao.Coin.Ctx(ctx).Insert(req)
	return
}
// Edit 修改
func (s *coin) Edit(ctx context.Context,req *dao.CoinEditReq) error {    
	_, err := dao.Coin.Ctx(ctx).FieldsEx(dao.Coin.Columns.Id).Where(dao.Coin.Columns.Id, req.Id).
		Update(req)
	return err
}
// DeleteByIds 删除
func (s *coin) DeleteByIds(ctx context.Context,ids []int) (err error) {
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
