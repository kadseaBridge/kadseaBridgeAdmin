// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/model/internal"
)

// SysUser is the golang structure for table sys_user.
type SysUser internal.SysUser

// Fill with you ideas below.

// LoginParamsReq 登陆参数
type LoginParamsReq struct {
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	GoogleCode string `p:"googleCode"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type GoogleAuthVerifyReq struct {
	UserId     uint64 `p:"userId" v:"required#userId不能为空"`
	GoogleCode string `p:"googleCode" v:"required#googleCode不能为空" `
}

// LoginUserRes 登录返回
type LoginUserRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`           //
	UserName     string `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string `orm:"user_password"    json:"userPassword"` // 登录密码;cmf_password加密
	GoogleAuth   string `orm:"googleAuth"    json:"googleAuth"`      // google验证码
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`      // 是否后台管理员 1 是  0   否
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
	DeptId       uint64 `orm:"dept_id"       json:"deptId"`          //部门id
}

// SysUserSearchReq 用户搜索请求参数
type SysUserSearchReq struct {
	DeptId      string  `p:"deptId"` //部门id
	DeptIds     []int64 //所属部门id数据
	Phonenumber string  `p:"phonenumber"`
	Status      string  `p:"status"`
	KeyWords    string  `p:"userName"`
	comModel.PageReq
}

// SetUserReq 添加修改用户公用请求字段
type SetUserReq struct {
	DeptId      uint64  `p:"deptId" v:"required#用户部门不能为空"` //所属部门
	Email       string  `p:"email" v:"email#邮箱格式错误"`       //邮箱
	NickName    string  `p:"nickName" v:"required#用户昵称不能为空"`
	Phonenumber string  `p:"phonenumber" v:"required|phone#手机号不能为空|手机号格式错误"`
	PostIds     []int64 `p:"postIds"`
	Remark      string  `p:"remark"`
	RoleIds     []int64 `p:"roleIds"`
	Sex         int     `p:"sex"`
	Status      uint    `p:"status"`
	IsAdmin     int     `p:"isAdmin"` // 是否后台管理员 1 是  0   否
}

// AddUserReq 添加用户参数
type AddUserReq struct {
	SetUserReq
	UserName   string `p:"userName" v:"required#用户账号不能为空"`
	Password   string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	UserSalt   string
	Googleauth string
}

type EditUserReq struct {
	SetUserReq
	UserId int `p:"userId" v:"required#用户id不能为空"`
}

type SysUserRoleDeptRes struct {
	*SysUser
	Dept     *SysDept `json:"dept"`
	RoleInfo []*struct {
		RoleId uint   `json:"roleId"`
		Name   string `json:"name"`
	} `json:"roleInfo"`
	Post []*struct {
		PostId   int64  `json:"postId"`
		PostName string `json:"postName"`
	} `json:"post"`
}

// SysUserResetPwdReq 重置用户密码状态参数
type SysUserResetPwdReq struct {
	Id       uint64 `p:"userId" v:"required#用户id不能为空"`
	Password string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
}

// SysUserStatusReq 设置用户状态参数
type SysUserStatusReq struct {
	Id         uint64 `p:"userId" v:"required#用户id不能为空"`
	UserStatus uint   `p:"status" v:"required#用户状态不能为空"`
}

// ProfileUpReq 个人中心修改用户信息参数
type ProfileUpReq struct {
	UserId       uint64
	UserNickname string `p:"userNickname" v:"required#用户昵称不能为空" orm:"user_nickname"` // 用户昵称
	Mobile       string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误" orm:"mobile,unique"`
	UserEmail    string `p:"userEmail" v:"email#邮箱格式错误" orm:"user_email"`
	Sex          int    `p:"sex" orm:"sex"`
	Googleauth   string
}

// ProfileUpdatePwdReq 个人中心修改用户密码
type ProfileUpdatePwdReq struct {
	UserId      uint64
	OldPassword string `p:"oldPassword" v:"required#旧密码不能为空"`
	NewPassword string `p:"newPassword" v:"required#新密码不能为空"`
}
