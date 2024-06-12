// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-12 23:29:47
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
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns ChainColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ChainColumns defines and stores column names for table chain.
type ChainColumns struct {
	Id                    string //
	Name                  string // 链名
	ChainId               string // 链Id
	Rpc                   string // 链rpc
	Explorer              string // 链浏览器地址
	IsEnable              string // 状态
	CreateAt              string // 创建时间
	UpdateAt              string // 更新时间
	ChainType             string // 0 ethereum 1 tron  2 btc
	Icon                  string // 链图标
	IsDelete              string // 软删除
	MaxGasPrice           string // 支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时， 为0时，所有订单都不进行支付， 如果想要所有的订单都进行支付，不管gasprice 则设置很大   单位：wei
	EnablePay             string // 该公链作为目标时，是否允许支付
	BridgeContractAddress string // 跨链桥合约地址
}

var chainColumns = ChainColumns{
	Id:                    "id",
	Name:                  "name",
	ChainId:               "chainId",
	Rpc:                   "rpc",
	Explorer:              "explorer",
	IsEnable:              "isEnable",
	CreateAt:              "create_at",
	UpdateAt:              "update_at",
	ChainType:             "type",
	Icon:                  "icon",
	IsDelete:              "isDelete",
	MaxGasPrice:           "max_gas_price",
	EnablePay:             "enablePay",
	BridgeContractAddress: "bridge_contract_address",
}

// NewChainDao creates and returns a new DAO object for table data access.
func NewChainDao() *ChainDao {
	return &ChainDao{
		Group:   "default",
		Table:   "chain",
		Columns: chainColumns,
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
