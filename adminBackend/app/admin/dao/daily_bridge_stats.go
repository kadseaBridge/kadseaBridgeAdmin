// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2024-07-04 17:54:12
// 生成路径: gfast/app/admin/dao/daily_bridge_stats.go
// 生成人：jimmy
// ==========================================================================

package dao

import (
	"gfast/app/admin/dao/internal"
	comModel "gfast/app/common/model"
	"github.com/gogf/gf/os/gtime"
)

// dailyBridgeStatsDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type dailyBridgeStatsDao struct {
	*internal.DailyBridgeStatsDao
}

var (
	// DailyBridgeStats is globally public accessible object for table tools_gen_table operations.
	DailyBridgeStats = dailyBridgeStatsDao{
		internal.NewDailyBridgeStatsDao(),
	}
)

// Fill with you ideas below.

// DailyBridgeStatsSearchReq 分页请求参数
type DailyBridgeStatsSearchReq struct {
	Date      *gtime.Time `p:"date"`      //日期
	CoinName  string      `p:"coinName"`  //币种
	ChainName string      `p:"chainName"` //链类型
	comModel.PageReq
}

// DailyBridgeStatsAddReq 添加操作请求参数
type DailyBridgeStatsAddReq struct {
}

// DailyBridgeStatsEditReq 修改操作请求参数
type DailyBridgeStatsEditReq struct {
	Id int `p:"id" v:"required#主键ID不能为空"`
}
