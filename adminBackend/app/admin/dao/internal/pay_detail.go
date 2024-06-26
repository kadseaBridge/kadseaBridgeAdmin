// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-08 23:23:41
// 生成路径: adminBackend/app/admin/dao/internal/pay_detail.go
// 生成人：jimmy
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// PayDetailDao is the manager for logic model data accessing and custom defined data operations functions management.
type PayDetailDao struct {
	Table   string           // Table is the underlying table name of the DAO.
	Group   string           // Group is the database configuration group name of current DAO.
	Columns PayDetailColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// PayDetailColumns defines and stores column names for table pay_detail.
type PayDetailColumns struct {
	Id              string //
	OrderId         string //
	PayChainId      string //
	PayCoinAddress  string //
	TransactionHash string //
	PayAmount       string //
	PayTime         string //
	PayAddress      string //
	CreateAt        string //
	UpdateAt        string //
	Remark          string //
	BlockNumber     string //
	PayGasFee       string //
}

var payDetailColumns = PayDetailColumns{
	Id:              "id",
	OrderId:         "order_id",
	PayChainId:      "pay_chainId",
	PayCoinAddress:  "pay_coin_address",
	TransactionHash: "transaction_hash",
	PayAmount:       "pay_amount",
	PayTime:         "pay_time",
	PayAddress:      "pay_address",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
	Remark:          "remark",
	BlockNumber:     "block_number",
	PayGasFee:       "pay_gas_fee",
}

// NewPayDetailDao creates and returns a new DAO object for table data access.
func NewPayDetailDao() *PayDetailDao {
	return &PayDetailDao{
		Group:   "default",
		Table:   "pay_detail",
		Columns: payDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayDetailDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
