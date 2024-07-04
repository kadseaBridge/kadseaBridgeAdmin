// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-07-04 17:54:12
// 生成路径: gfast/app/admin/dao/internal/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// DailyBridgeStatsDao is the manager for logic model data accessing and custom defined data operations functions management.
type DailyBridgeStatsDao struct {
	Table   string                  // Table is the underlying table name of the DAO.
	Group   string                  // Group is the database configuration group name of current DAO.
	Columns DailyBridgeStatsColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// DailyBridgeStatsColumns defines and stores column names for table daily_bridge_stats.
type DailyBridgeStatsColumns struct {
	Id                 string // id
	Date               string // 日期
	CoinName           string // 币种
	ChainName          string // 链类型
	TransferIn         string // 跨链转入
	TransferOut        string // 跨链转出
	TransferDifference string // 跨链差额
	Fee                string // 跨链手续费
	PlatformAssets     string // 平台资产
	FinancialBalance   string // 财务余额
	CreatedAt          string //
	UpdatedAt          string //
}

var dailyBridgeStatsColumns = DailyBridgeStatsColumns{
	Id:                 "id",
	Date:               "date",
	CoinName:           "coin_name",
	ChainName:          "chain_name",
	TransferIn:         "transfer_in",
	TransferOut:        "transfer_out",
	TransferDifference: "transfer_difference",
	Fee:                "fee",
	PlatformAssets:     "platform_assets",
	FinancialBalance:   "financial_balance",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewDailyBridgeStatsDao creates and returns a new DAO object for table data access.
func NewDailyBridgeStatsDao() *DailyBridgeStatsDao {
	return &DailyBridgeStatsDao{
		Group:   "default",
		Table:   "daily_bridge_stats",
		Columns: dailyBridgeStatsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DailyBridgeStatsDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DailyBridgeStatsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DailyBridgeStatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
