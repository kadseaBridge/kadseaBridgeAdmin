// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-06-08 23:23:41
// 生成路径: gfast/app/admin/dao/pay_detail.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/os/gtime"
)

// payDetailDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type payDetailDao struct {
	*internal.PayDetailDao
}

var (
	// PayDetail is globally public accessible object for table tools_gen_table operations.
	PayDetail = payDetailDao{
		internal.NewPayDetailDao(),
	}
)

// Fill with you ideas below.

// PayDetailSearchReq 分页请求参数
type PayDetailSearchReq struct {
	OrderId string `p:"orderId"` //
	comModel.PageReq
}

// PayDetailAddReq 添加操作请求参数
type PayDetailAddReq struct {
	OrderId         string      `p:"orderId" v:"required#不能为空"`
	PayChainId      string      `p:"payChainId" v:"required#不能为空"`
	PayCoinAddress  string      `p:"payCoinAddress" v:"required#不能为空"`
	TransactionHash string      `p:"transactionHash" v:"required#不能为空"`
	PayAmount       float64     `p:"payAmount" v:"required#不能为空"`
	PayTime         *gtime.Time `p:"payTime" v:"required#不能为空"`
	PayAddress      string      `p:"payAddress" v:"required#不能为空"`
	CreateAt        *gtime.Time `p:"createAt" `
	UpdateAt        *gtime.Time `p:"updateAt" `
	Remark          string      `p:"remark" `
	BlockNumber     int64       `p:"blockNumber" `
	PayGasFee       float64     `p:"payGasFee" v:"required#不能为空"`
}

// PayDetailEditReq 修改操作请求参数
type PayDetailEditReq struct {
	Id int64 `p:"id" v:"required#主键ID不能为空"`
}
