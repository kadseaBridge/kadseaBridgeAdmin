// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gfast/app/system/model/internal"
)

// SysAuthRule is the golang structure for table sys_auth_rule.
type SysAuthRule internal.SysAuthRule

// Fill with you ideas below.
type SysAuthRuleInfoRes struct {
	Id         uint   `orm:"id,primary"  json:"id"`         //
	Pid        uint   `orm:"pid"         json:"pid"`        // 父ID
	Name       string `orm:"name,unique" json:"name"`       // 规则名称
	Title      string `orm:"title"       json:"title"`      // 规则名称
	Icon       string `orm:"icon"        json:"icon"`       // 图标
	Condition  string `orm:"condition"   json:"condition"`  // 条件
	Remark     string `orm:"remark"      json:"remark"`     // 备注
	MenuType   uint   `orm:"menu_type"   json:"menuType"`   // 类型 0目录 1菜单 2按钮
	Weigh      int    `orm:"weigh"       json:"weigh"`      // 权重
	Status     uint   `orm:"status"      json:"status"`     // 状态
	AlwaysShow uint   `orm:"always_show" json:"alwaysShow"` // 显示状态
	Path       string `orm:"path"        json:"path"`       // 路由地址
	JumpPath   string `orm:"jump_path"   json:"jumpPath"`   // 跳转路由
	Component  string `orm:"component"   json:"component"`  // 组件路径
	IsFrame    uint   `orm:"is_frame"    json:"isFrame"`    // 是否外链 1是 0否
	ModuleType string `orm:"module_type" json:"moduleType"` // 所属模块
	ModelId    uint   `orm:"model_id"    json:"modelId"`    // 模型ID
}

type SysAuthRuleReqSearch struct {
	Status string `p:"status" `
	Title  string `p:"menuName" `
}

// SysAuthRuleTreeRes 菜单树形结构
type SysAuthRuleTreeRes struct {
	*SysAuthRuleInfoRes
	Children []*SysAuthRuleTreeRes `json:"children"`
}

func (req *SysAuthRuleReqSearch) IsEmpty() bool {
	return req.Status == "" && req.Title == ""
}

//菜单对象
type MenuReq struct {
	MenuType   uint   `orm:"menu_type" p:"menuType"  v:"min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max"`
	Pid        uint   `orm:"pid"  p:"parentId"  v:"min:0"`
	Name       string `orm:"name,unique" p:"name" v:"required#请填写规则名称"`
	Title      string `orm:"title" p:"menuName" v:"required|length:1,100#请填写标题|标题长度在:min到:max位"`
	Icon       string `orm:"icon"  p:"icon"`
	Weigh      int    `orm:"weigh"  p:"orderNum" `
	Condition  string `orm:"condition" p:"condition" `
	Remark     string `orm:"remark" p:"remark" `
	Status     uint   `orm:"status"  p:"status" `
	AlwaysShow uint   `orm:"always_show" p:"visible"`
	Path       string `orm:"path"   p:"path"`
	Component  string `orm:"component"  p:"component" v:"required-if:menuType,1#组件路径不能为空"`
	IsFrame    uint   `orm:"is_frame" p:"is_frame"`
	ModuleType string `orm:"module_type" p:"moduleType"`
	ModelId    uint   `orm:"model_id"    p:"modelId"`
}
