// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:15
// 生成路径: gfast/app/admin/dao/internal/chain.go
// 生成人：jimmy
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// ChainDao is the manager for logic model data accessing and custom defined data operations functions management.
type ChainDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns ChainColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// ChainColumns defines and stores column names for table chain.
type ChainColumns struct {    
    Id  string  //    
    Name  string  // 链名称    
    ChainId  string  // 链id    
    Rpc  string  // rpc链接    
    Explorer  string  // 浏览器链接    
    IsEnable  string  // 是否启用    
    CreateAt  string  //    
    UpdateAt  string  //    
    ChainType  string  // 0 ethereum 1 tron  2 btc    
    Icon  string  // 图标    
}
var chainColumns = ChainColumns{    
    Id:  "id",    
    Name:  "name",    
    ChainId:  "chainId",    
    Rpc:  "rpc",    
    Explorer:  "explorer",    
    IsEnable:  "isEnable",    
    CreateAt:  "create_at",    
    UpdateAt:  "update_at",    
    ChainType:  "type",    
    Icon:  "icon",    
}
// NewChainDao creates and returns a new DAO object for table data access.
func NewChainDao() *ChainDao {
	return &ChainDao{
        Group:    "default",
        Table: "chain",
        Columns:chainColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChainDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}