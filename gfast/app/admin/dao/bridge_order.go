// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-06-03 19:16:53
// 生成路径: gfast/app/admin/dao/bridge_order.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/os/gtime"
)

// bridgeOrderDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type bridgeOrderDao struct {
	*internal.BridgeOrderDao
}

var (
	// BridgeOrder is globally public accessible object for table tools_gen_table operations.
	BridgeOrder = bridgeOrderDao{
		internal.NewBridgeOrderDao(),
	}
)

// Fill with you ideas below.

//queryParams: {
//pageNum: 1,
//pageSize: 10,
//sourceAddress: undefined,
//targetAddress: undefined,
//sourceCoinAddress: undefined,
//targetCoinAddress: undefined,
//sourceChainId: undefined,
//targetChainId: undefined,
//transactionHash: undefined,
//status: undefined,
//startRange: [],
//endRange: [],
//},

// BridgeOrderSearchReq 分页请求参数
type BridgeOrderSearchReq struct {
	SourceAddress     string      `p:"sourceAddress"`     //转出币种
	TargetAddress     string      `p:"targetAddress"`     //转入币种
	SourceCoinAddress string      `p:"sourceCoinAddress"` //转出钱包地址
	TargetCoinAddress string      `p:"targetCoinAddress"` //转入钱包地址
	SourceChainId     string      `p:"sourceChainId"`     //转出链
	TargetChainId     string      `p:"targetChainId"`     //转入链
	TransactionHash   string      `p:"transactionHash"`   //交易哈希
	Status            string      `p:"status"`            //跨链记录状态
	StartDate1        *gtime.Time `p:"startDate1"`        //发起时间1
	StartDate2        *gtime.Time `p:"startDate2"`        //发起时间2
	EndDate1          *gtime.Time `p:"endDate1"`          //结束时间1
	EndDate2          *gtime.Time `p:"endDate2"`          //结束时间2
	comModel.PageReq
}

// BridgeOrderAddReq 添加操作请求参数
type BridgeOrderAddReq struct {
	SourceAddress     string      `p:"sourceAddress" v:"required#转出币种不能为空"`
	TargetAddress     string      `p:"targetAddress" v:"required#转入币种不能为空"`
	SourceCoinAddress string      `p:"sourceCoinAddress" v:"required#转出钱包地址不能为空"`
	TargetCoinAddress string      `p:"targetCoinAddress" v:"required#转入钱包地址不能为空"`
	SourceChainId     string      `p:"sourceChainId" v:"required#转出链不能为空"`
	TargetChainId     string      `p:"targetChainId" v:"required#转入链不能为空"`
	TransactionHash   string      `p:"transactionHash" v:"required#交易哈希不能为空"`
	OrderId           string      `p:"orderId" v:"required#跨链订单id不能为空"`
	Amount            float64     `p:"amount" v:"required#数量不能为空"`
	Status            int         `p:"status" v:"required#跨链记录状态不能为空"`
	CreateAt          *gtime.Time `p:"createAt" `
	UpdateAt          *gtime.Time `p:"updateAt" `
	BlockNumber       int64       `p:"blockNumber" v:"required#区块高度不能为空"`
	Fee               float64     `p:"fee" v:"required#手续费不能为空"`
	Remark            string      `p:"remark" `
}

// BridgeOrderEditReq 修改操作请求参数
//type BridgeOrderEditReq struct {
//	Id     int64 `p:"id" v:"required#主键ID不能为空"`
//	Status int   `p:"status" v:"required#跨链记录状态不能为空"`
//}
//
//// BridgeOrderStatusReq 设置用户状态参数
//type BridgeOrderStatusReq struct {
//	Id     int64 `p:"id" v:"required#主键ID不能为空"`
//	Status int   `p:"status" v:"required#跨链记录状态不能为空"`
//}

type BridgeOrderEditReq struct {
	Id           int64 `p:"id" v:"required#主键ID不能为空"`
	ReviewStatus int   `p:"reviewStatus" v:"required#跨链记录状态不能为空"`
}

// BridgeOrderStatusReq 设置用户状态参数
type BridgeOrderStatusReq struct {
	Id           int64 `p:"id" v:"required#主键ID不能为空"`
	ReviewStatus int   `p:"reviewStatus" v:"required#跨链记录状态不能为空"`
}
