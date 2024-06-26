// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: adminBackend/app/admin/dao/internal/bridge_order.go
// 生成人：jimmy
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// BridgeOrderDao is the manager for logic model data accessing and custom defined data operations functions management.
type BridgeOrderDao struct {
	Table   string             // Table is the underlying table name of the DAO.
	Group   string             // Group is the database configuration group name of current DAO.
	Columns BridgeOrderColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// BridgeOrderColumns defines and stores column names for table bridge_order.
type BridgeOrderColumns struct {
	Id                string //
	SourceAddress     string // 转出币种
	TargetAddress     string // 转入币种
	SourceCoinAddress string // 转出钱包地址
	TargetCoinAddress string // 转入钱包地址
	SourceChainId     string // 转出链
	TargetChainId     string // 转入链
	TransactionHash   string // 交易哈希
	OrderId           string // 跨链订单id
	Amount            string // 数量
	Status            string // 跨链记录状态
	CreateAt          string // 发起时间
	UpdateAt          string // 结束时间
	BlockNumber       string // 区块高度
	Fee               string // 手续费
	Remark            string // 评论
}

var bridgeOrderColumns = BridgeOrderColumns{
	Id:                "id",
	SourceAddress:     "source_address",
	TargetAddress:     "target_address",
	SourceCoinAddress: "source_coin_address",
	TargetCoinAddress: "target_coin_address",
	SourceChainId:     "source_chainId",
	TargetChainId:     "target_chainId",
	TransactionHash:   "transaction_hash",
	OrderId:           "orderId",
	Amount:            "amount",
	Status:            "status",
	CreateAt:          "create_at",
	UpdateAt:          "update_at",
	BlockNumber:       "block_number",
	Fee:               "fee",
	Remark:            "remark",
}

// NewBridgeOrderDao creates and returns a new DAO object for table data access.
func NewBridgeOrderDao() *BridgeOrderDao {
	return &BridgeOrderDao{
		Group:   "default",
		Table:   "bridge_order",
		Columns: bridgeOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BridgeOrderDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BridgeOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BridgeOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
