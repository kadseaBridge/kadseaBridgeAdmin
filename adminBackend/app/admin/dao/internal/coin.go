// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-05-30 18:31:17
// 生成路径: adminBackend/app/admin/dao/internal/coin.go
// 生成人：jimmy
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CoinDao is the manager for logic model data accessing and custom defined data operations functions management.
type CoinDao struct {
	Table   string      // Table is the underlying table name of the DAO.
	Group   string      // Group is the database configuration group name of current DAO.
	Columns CoinColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CoinColumns defines and stores column names for table coin.
type CoinColumns struct {
	Id        string //
	Name      string // 代币名称
	Symbol    string // 代币简称
	ChainId   string // 链Id
	Address   string // 代币地址
	IsEnable  string // 是否上架
	CreateAt  string //
	UpdateAt  string //
	TokenType string // 币种类型
	Decimals  string // 精度
	Icon      string // 代币图标
}

var coinColumns = CoinColumns{
	Id:        "id",
	Name:      "name",
	Symbol:    "symbol",
	ChainId:   "chainId",
	Address:   "address",
	IsEnable:  "isEnable",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	TokenType: "type",
	Decimals:  "decimals",
	Icon:      "icon",
}

// NewCoinDao creates and returns a new DAO object for table data access.
func NewCoinDao() *CoinDao {
	return &CoinDao{
		Group:   "default",
		Table:   "coin",
		Columns: coinColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CoinDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CoinDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
