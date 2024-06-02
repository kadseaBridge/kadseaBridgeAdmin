// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: gfast/app/admin/service/bridge_config.go
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
func (s *bridgeConfig) GetInfoById(ctx context.Context,id int64) (info *model.BridgeConfig, err error) {
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
func (s *bridgeConfig) Add(ctx context.Context,req *dao.BridgeConfigAddReq) (err error) {
	_, err = dao.BridgeConfig.Ctx(ctx).Insert(req)
	return
}
// Edit 修改
func (s *bridgeConfig) Edit(ctx context.Context,req *dao.BridgeConfigEditReq) error {    
	_, err := dao.BridgeConfig.Ctx(ctx).FieldsEx(dao.BridgeConfig.Columns.Id).Where(dao.BridgeConfig.Columns.Id, req.Id).
		Update(req)
	return err
}
// DeleteByIds 删除
func (s *bridgeConfig) DeleteByIds(ctx context.Context,ids []int) (err error) {
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
