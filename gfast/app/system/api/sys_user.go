package api

import (
	"fmt"
	commonService "gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"gfast/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type user struct {
	SystemBase
}

var User = new(user)

// UserList
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户管理
// @Param req query model.SysUserSearchReq  true "data"
// @Success 0 {object} library.Response "{"code": 200, "data": [...]}"
// @Router /system/auth/userList [get]
// @Security
func (c *user) UserList(r *ghttp.Request) {
	var req *model.SysUserSearchReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}

	if err := service.SysUser.WriteDeptIdsOfSearchReq(req); err != nil {
		c.FailJsonExit(r, err.Error())
	}

	total, page, userList, err := service.SysUser.GetUserList(req)
	if err != nil {
		c.FailJsonExit(r, "获取用户列表数据失败")
	}

	users, err := service.SysUser.GetUsersRoleDept(userList)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}

	c.SusJsonExit(r, g.Map{
		"total":       total,
		"currentPage": page,
		"userList":    users,
	})
}

// Get 获取添加用户信息
func (c *user) Get(r *ghttp.Request) {
	roleListEntities, err := service.SysRole.GetRoleList()
	if err != nil {
		g.Log().Error(err)
		c.FailJsonExit(r, "获取角色数据失败")
	}
	posts, err := service.SysUser.GetUsedPost()
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, g.Map{
		"roleList": roleListEntities,
		"posts":    posts,
	})
}

// AddUser 添加用户提交
func (c *user) AddUser(r *ghttp.Request) {
	var req *model.AddUserReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.SysUser.AddUser(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "添加管理员成功")
}

// GetEditUser 获取编辑的用户
func (c *user) GetEditUser(r *ghttp.Request) {
	id := r.GetUint64("id")
	res, err := service.SysUser.GetEditUser(id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, res)
}

// EditUser 编辑用户提交
func (c *user) EditUser(r *ghttp.Request) {
	var req *model.EditUserReq
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	err := service.SysUser.EditUser(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "修改管理员成功")
}

// GetInfo 获取登陆用户信息
func (c *user) GetInfo(r *ghttp.Request) {
	userInfo := c.GetCurrentUser(r.Context())
	rolesList := make([]string, 0, 10)
	var permissions []string
	userId := userInfo.Id
	allRoles, err := service.SysRole.GetRoleList()
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	roles, err := service.SysUser.GetAdminRole(userId, allRoles)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	name := make([]string, len(roles))
	roleIds := make([]uint, len(roles))
	for k, v := range roles {
		name[k] = v.Name
		roleIds[k] = v.Id
	}
	isSuperAdmin := false
	service.SysUser.NotCheckAuthAdminIds.Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == userId {
			isSuperAdmin = true
			return false
		}
		return true
	})
	if isSuperAdmin {
		permissions = []string{"*/*/*"}
	} else {
		permissions, err = service.SysUser.GetPermissions(roleIds)
	}
	rolesList = name
	c.SusJsonExit(r, g.Map{
		"user":        userInfo,
		"roles":       rolesList,
		"permissions": permissions,
	})
}

// GetRouters 获取用户菜单
func (c *user) GetRouters(r *ghttp.Request) {
	//获取用户信息
	userInfo := c.GetCurrentUser(r.Context())
	//是否超管
	isSuperAdmin := false
	userId := userInfo.Id
	//获取无需验证权限的用户id
	service.SysUser.NotCheckAuthAdminIds.Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == userId {
			isSuperAdmin = true
			return false
		}
		return true
	})
	//获取用户角色信息
	allRoles, err := service.SysRole.GetRoleList()
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	roles, e := service.SysUser.GetAdminRole(userId, allRoles)
	if e != nil {
		c.FailJsonExit(r, e.Error())
	}
	name := make([]string, len(roles))
	roleIds := make([]uint, len(roles))
	for k, v := range roles {
		name[k] = v.Name
		roleIds[k] = v.Id
	}
	var menuList []service.UserMenus
	//获取菜单信息
	if isSuperAdmin {
		//超管获取所有菜单
		menuList, err = service.SysUser.GetAllMenus()
	} else {
		menuList, err = service.SysUser.GetAdminMenusByRoleIds(roleIds)
	}
	if err != nil {
		c.FailJsonExit(r, e.Error())
	}
	if menuList == nil {
		menuList = make([]service.UserMenus, 0)
	}
	adminMenus := make([]service.UserMenus, 0, len(menuList))
	//过滤非后台模块的菜单
	for _, v := range menuList {
		if v.ModuleType == "sys_admin" {
			adminMenus = append(adminMenus, v)
		}
	}
	c.SusJsonExit(r, adminMenus)
}

func (c *user) ResetUserPwd(r *ghttp.Request) {
	var req *model.SysUserResetPwdReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if err := service.SysUser.ResetUserPwd(req); err != nil {
		c.FailJsonExit(r, err.Error())
	} else {
		c.SusJsonExit(r, "用户密码重置成功")
	}
}

func (c *user) ChangeUserStatus(r *ghttp.Request) {
	var req *model.SysUserStatusReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if err := service.SysUser.ChangeUserStatus(req); err != nil {
		c.FailJsonExit(r, err.Error())
	} else {
		c.SusJsonExit(r, "用户状态设置成功")
	}
}

func (c *user) DeleteUser(r *ghttp.Request) {
	ids := r.GetInts("ids")
	if len(ids) > 0 {
		err := service.SysUser.DeleteUser(r.GetCtx(), ids)
		if err != nil {
			g.Log().Error(err)
			c.FailJsonExit(r, "删除失败")
		}
	} else {
		c.FailJsonExit(r, "删除失败，参数错误")
	}

	c.SusJsonExit(r, "删除成功")
}

// Login   后台用户登陆验证
func (c *user) Login(r *ghttp.Request) (string, interface{}) {
	var ctx = r.GetCtx()
	var apiReq *model.LoginParamsReq
	if err := r.Parse(&apiReq); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
	}
	//判断验证码是否正确
	debug := genv.GetWithCmd("gf.debug")
	fmt.Println(apiReq.VerifyCode)
	if debug.Int() != 1 {
		if !commonService.Captcha.VerifyString(apiReq.VerifyKey, apiReq.VerifyCode) {
			c.FailJson(true, r, "验证码输入错误")
		}
	}

	ip := library.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	if user, err := service.SysUser.GetAdminUserByUsernamePassword(ctx, apiReq); err != nil {
		//保存日志（异步）
		service.SysLoginLog.Invoke(&model.LoginLogParams{
			Status:    0,
			Username:  apiReq.Username,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		c.FailJsonExit(r, err.Error())
	} else if user != nil {

		// 用户未绑定谷歌验证码
		if user.GoogleAuth == "" {
			secret, qrCode, err := commonService.NewGoogleAuth2().GenerateSecretAndQRCode(user.UserName)
			if err != nil {
				c.FailJsonExit(r, "生成谷歌验证码失败")
			}
			// 绑定谷歌验证码
			user.GoogleAuth = secret
			if err := service.SysUser.UpdateGoogleAuth(user.Id, secret); err != nil {
				c.FailJsonExit(r, "绑定谷歌验证码密钥失败")
			}
			c.SusJsonExit(r, g.Map{
				"bindGoogleAuth": true,
				"qrcode":         qrCode,
			})
		} else {
			// 验证谷歌验证码
			//if !commonService.NewGoogleAuth2().Verify(user.GoogleAuth, apiReq.GoogleCode) {
			//	c.FailJsonExit(r, "谷歌验证码错误")
			//}
			//r.SetParam("userInfo", user)
			//service.SysUser.UpdateLoginInfo(user.Id, apiReq.Username, ip, userAgent, "登录成功", "系统后台")
			//var keys string
			//if MultiLogin {
			//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip)
			//} else {
			//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
			//}
			//return keys, user

			//r.SetParam("userInfo", user)

			c.SusJsonExit(r, g.Map{
				"googleAuthRequired": true,
			})
		}

		//r.SetParam("userInfo", user)
		////更新用户登录记录 写入日志信息
		//service.SysUser.UpdateLoginInfo(user.Id, apiReq.Username, ip, userAgent, "登录成功", "系统后台")
		//var keys string
		//if MultiLogin {
		//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip)
		//} else {
		//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
		//}
		//return keys, user

	}
	return "", nil

}

//func (c *user) VerifyGoogleAuth(r *ghttp.Request) (string, interface{}) {
//	var ctx = r.GetCtx()
//	var apiReq *model.GoogleAuthVerifyReq
//	if err := r.Parse(&apiReq); err != nil {
//		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
//	}
//
//	ip := library.GetClientIp(r)
//	userAgent := r.Header.Get("User-Agent")
//	userInfo, err := service.SysUser.GetAdminByUserId(ctx, apiReq.UserId)
//	if err != nil {
//		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
//	}
//	if userInfo == nil {
//		c.FailJsonExit(r, "用户信息缺失")
//	}
//
//	// 假设这里有验证谷歌验证码的逻辑，并且验证通过
//	if commonService.NewGoogleAuth2().VerifyCode(userInfo.GoogleAuth, apiReq.GoogleCode) {
//		// 验证通过，设置用户信息
//		r.SetParam("userInfo", userInfo)
//		//更新用户登录记录 写入日志信息
//		service.SysUser.UpdateLoginInfo(userInfo.Id, userInfo.UserName, ip, userAgent, "登录成功", "系统后台")
//		var keys string
//		if MultiLogin {
//			keys = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword+ip)
//		} else {
//			keys = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword)
//		}
//		return keys, userInfo
//		//c.SusJsonExit(r, g.Map{
//		//	"loginSuccess": true,
//		//	"userInfo":     userInfo,
//		//})
//	} else {
//		c.FailJsonExit(r, "谷歌验证码错误")
//	}
//
//	return "", nil
//}
