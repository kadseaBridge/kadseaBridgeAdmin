// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-06-12 23:29:47
// 生成路径: gfast/app/admin/dao/chain.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/os/gtime"
)

// chainDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type chainDao struct {
	*internal.ChainDao
}

var (
	// Chain is globally public accessible object for table tools_gen_table operations.
	Chain = chainDao{
		internal.NewChainDao(),
	}
)

// Fill with you ideas below.

// ChainSearchReq 分页请求参数
type ChainSearchReq struct {
	Name string `p:"name"` //链名
	comModel.PageReq
}

// ChainAddReq 添加操作请求参数
type ChainAddReq struct {
	Name                  string      `p:"name" v:"required#链名不能为空"`
	ChainId               string      `p:"chainId" v:"required#链Id不能为空"`
	Rpc                   string      `p:"rpc" v:"required#链rpc不能为空"`
	Explorer              string      `p:"explorer" `
	IsEnable              int         `p:"isEnable" `
	CreateAt              *gtime.Time `p:"createAt" `
	UpdateAt              *gtime.Time `p:"updateAt" `
	ChainType             int         `p:"type" v:"required#0 ethereum 1 tron  2 btc不能为空"`
	Icon                  string      `p:"icon" `
	IsDelete              int         `p:"isDelete" `
	MaxGasPrice           float64     `p:"maxGasPrice" v:"required#maxGasPrice不能为空"`
	EnablePay             int         `p:"enablePay" v:"required#该公链作为目标时，是否允许支付不能为空"`
	BridgeContractAddress string      `p:"bridgeContractAddress" v:"required#跨链桥合约地址不能为空"`
}

// ChainEditReq 修改操作请求参数
type ChainEditReq struct {
	Id                    int64       `p:"id" v:"required#主键ID不能为空"`
	Name                  string      `p:"name" v:"required#链名不能为空"`
	ChainId               string      `p:"chainId" v:"required#链Id不能为空"`
	Rpc                   string      `p:"rpc" v:"required#链rpc不能为空"`
	Explorer              string      `p:"explorer" `
	IsEnable              int         `p:"isEnable"`
	CreateAt              *gtime.Time `p:"createAt" `
	UpdateAt              *gtime.Time `p:"updateAt" `
	ChainType             int         `p:"type" v:"required#0 ethereum 1 tron  2 btc不能为空"`
	Icon                  string      `p:"icon" `
	IsDelete              int         `p:"isDelete" `
	MaxGasPrice           float64     `p:"maxGasPrice" v:"required#maxGasPrice不能为空"`
	EnablePay             int         `p:"enablePay" v:"required#该公链作为目标时，是否允许支付不能为空"`
	BridgeContractAddress string      `p:"bridgeContractAddress" v:"required#跨链桥合约地址不能为空"`
}
