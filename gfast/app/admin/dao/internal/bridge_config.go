// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:01
// 生成路径: gfast/app/admin/dao/internal/bridge_config.go
// 生成人：jimmy
// ==========================================================================


package internal
import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)
// BridgeConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
type BridgeConfigDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns BridgeConfigColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}
// BridgeConfigColumns defines and stores column names for table bridge_config.
type BridgeConfigColumns struct {    
    Id  string  //    
    SourceChainId  string  // 当前链    
    TargetChainId  string  // 对手链    
    SourceCoinAddress  string  // 币种    
    FeeFixed  string  // 固定金额手续费，    
    FeePercent  string  // 手续费百分比，如果不需要则设置为0    
    DayTotal  string  // 跨入每日审核数量    
    OnceTotal  string  // 跨入单次审核数量    
    IsEnable  string  // 状态    
    TargetCoinAddress  string  // 目标链合约地址    
    UpdateAt  string  // 更新时间    
    CreateAt  string  // 创建时间    
}
var bridgeConfigColumns = BridgeConfigColumns{    
    Id:  "id",    
    SourceChainId:  "source_chainId",    
    TargetChainId:  "target_chainId",    
    SourceCoinAddress:  "source_coin_address",    
    FeeFixed:  "fee_fixed",    
    FeePercent:  "fee_percent",    
    DayTotal:  "day_total",    
    OnceTotal:  "once_total",    
    IsEnable:  "isEnable",    
    TargetCoinAddress:  "target_coin_address",    
    UpdateAt:  "update_at",    
    CreateAt:  "create_at",    
}
// NewBridgeConfigDao creates and returns a new DAO object for table data access.
func NewBridgeConfigDao() *BridgeConfigDao {
	return &BridgeConfigDao{
        Group:    "default",
        Table: "bridge_config",
        Columns:bridgeConfigColumns,
	}
}
// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BridgeConfigDao) DB() gdb.DB {
	return g.DB(dao.Group)
}
// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BridgeConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}
// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BridgeConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}