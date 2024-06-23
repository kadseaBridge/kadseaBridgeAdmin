package api

import (
	"gfast/app/common/api"
	comModel "gfast/app/common/model"
	commonService "gfast/app/common/service"
	"gfast/app/system/model"
	"gfast/app/system/service"
	"gfast/library"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/gconv"
	"strings"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
	"github.com/mssola/user_agent"
)

type auth struct {
	api.CommonBase
}

var (
	Auth       = new(auth)
	MultiLogin = g.Cfg().GetBool("gToken.system.MultiLogin")
	GfToken    = &gtoken.GfToken{
		CacheMode:      g.Cfg().GetInt8("gToken.system.CacheMode"),
		CacheKey:       g.Cfg().GetString("gToken.system.CacheKey"),
		Timeout:        g.Cfg().GetInt("gToken.system.Timeout"),
		MaxRefresh:     g.Cfg().GetInt("gToken.system.MaxRefresh"),
		TokenDelimiter: g.Cfg().GetString("gToken.system.TokenDelimiter"),
		EncryptKey:     g.Cfg().GetBytes("gToken.system.EncryptKey"),
		AuthFailMsg:    g.Cfg().GetString("gToken.system.AuthFailMsg"),
		MultiLogin:     MultiLogin,
		LoginPath:      "/login",
		//LoginBeforeFunc: Auth.login,
		//LoginPath:        "/system/auth/verifyGoogleCode",
		LoginBeforeFunc:  Auth.VerifyGoogleCode,
		LoginAfterFunc:   Auth.loginAfter,
		LogoutPath:       "/logout",
		AuthExcludePaths: g.SliceStr{"/login", "/system/auth/login", "/system/auth/verifyGoogleCode"},
		AuthAfterFunc:    Auth.authAfterFunc,
		LogoutBeforeFunc: Auth.loginOut,
	}
)

// Login   后台用户登陆验证
//func (c *auth) login(r *ghttp.Request) (string, interface{}) {
//	var ctx = r.GetCtx()
//	var apiReq *model.LoginParamsReq
//	if err := r.Parse(&apiReq); err != nil {
//		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
//	}
//	//判断验证码是否正确
//	debug := genv.GetWithCmd("gf.debug")
//	fmt.Println(apiReq.VerifyCode)
//	if debug.Int() != 1 {
//		if !commonService.Captcha.VerifyString(apiReq.VerifyKey, apiReq.VerifyCode) {
//			c.FailJson(true, r, "验证码输入错误")
//		}
//	}
//
//	ip := library.GetClientIp(r)
//	userAgent := r.Header.Get("User-Agent")
//	if user, err := service.SysUser.GetAdminUserByUsernamePassword(ctx, apiReq); err != nil {
//		//保存日志（异步）
//		service.SysLoginLog.Invoke(&model.LoginLogParams{
//			Status:    0,
//			Username:  apiReq.Username,
//			Ip:        ip,
//			UserAgent: userAgent,
//			Msg:       err.Error(),
//			Module:    "系统后台",
//		})
//		c.FailJsonExit(r, err.Error())
//	} else if user != nil {
//
//		// 用户未绑定谷歌验证码
//		if user.GoogleAuth == "" {
//			secret, qrCode, err := commonService.NewGoogleAuth2().GenerateSecretAndQRCode(user.UserName)
//			if err != nil {
//				c.FailJsonExit(r, "生成谷歌验证码失败")
//			}
//			// 绑定谷歌验证码
//			user.GoogleAuth = secret
//			if err := service.SysUser.UpdateGoogleAuth(user.Id, secret); err != nil {
//				c.FailJsonExit(r, "绑定谷歌验证码密钥失败")
//			}
//			c.SusJsonExit(r, g.Map{
//				"bindGoogleAuth": true,
//				"qrcode":         qrCode,
//			})
//		} else {
//			// 验证谷歌验证码
//			//if !commonService.NewGoogleAuth2().Verify(user.GoogleAuth, apiReq.GoogleCode) {
//			//	c.FailJsonExit(r, "谷歌验证码错误")
//			//}
//			//r.SetParam("userInfo", user)
//			//service.SysUser.UpdateLoginInfo(user.Id, apiReq.Username, ip, userAgent, "登录成功", "系统后台")
//			//var keys string
//			//if MultiLogin {
//			//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip)
//			//} else {
//			//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
//			//}
//			//return keys, user
//
//			//r.SetParam("userInfo", user)
//
//			c.SusJsonExit(r, g.Map{
//				"googleAuthRequired": true,
//			})
//		}
//
//		//r.SetParam("userInfo", user)
//		////更新用户登录记录 写入日志信息
//		//service.SysUser.UpdateLoginInfo(user.Id, apiReq.Username, ip, userAgent, "登录成功", "系统后台")
//		//var keys string
//		//if MultiLogin {
//		//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip)
//		//} else {
//		//	keys = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
//		//}
//		//return keys, user
//
//	}
//	return "", nil
//
//}

// VerifyGoogleCode 成功验证谷歌验证码后再执行
// 此处假设你在其他地方完成了谷歌验证的逻辑，然后调用这个函数
func (c *auth) VerifyGoogleCode(r *ghttp.Request) (string, interface{}) {
	var ctx = r.GetCtx()
	var apiReq *model.GoogleAuthVerifyReq
	if err := r.Parse(&apiReq); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
	}

	ip := library.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	userInfo, err := service.SysUser.GetUserByUsername(ctx, apiReq.UserName)
	if err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).Current().Error())
	}
	if userInfo == nil {
		c.FailJsonExit(r, "用户信息缺失")
	}

	// 假设这里有验证谷歌验证码的逻辑，并且验证通过
	if commonService.NewGoogleAuth2().VerifyCode(apiReq.GoogleCode, userInfo.GoogleAuth) {
		// 验证通过，设置用户信息
		r.SetParam("userInfo", userInfo)
		//更新用户登录记录 写入日志信息
		service.SysUser.UpdateLoginInfo(userInfo.Id, userInfo.UserName, ip, userAgent, "登录成功", "系统后台")
		var keys string
		if MultiLogin {
			keys = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword+ip)
		} else {
			keys = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword)
		}
		return keys, userInfo
	} else {
		c.FailJsonExit(r, "谷歌验证码错误")
	}

	return "", nil
}

// 登录成功返回
func (c *auth) loginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		token := respData.GetString("token")
		uuid := respData.GetString("uuid")
		var userInfo *model.LoginUserRes
		r.GetParamVar("userInfo").Struct(&userInfo)
		//保存用户在线状态token到数据库
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		os := ua.OS()
		explorer, _ := ua.Browser()
		onlineData := &model.SysUserOnline{
			Uuid:       uuid,
			Token:      token,
			CreateTime: gtime.Now(),
			UserName:   userInfo.UserName,
			Ip:         library.GetClientIp(r),
			Explorer:   explorer,
			Os:         os,
		}
		//保存用户在线状态(异步)
		service.Online.Invoke(onlineData)
		c.SusJsonExit(r, g.Map{
			"token": token,
		})
	}
}

// gToken验证后返回
func (c *auth) authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if r.Method == "OPTIONS" || respData.Success() {
		r.Middleware.Next()
	} else {
		c.JsonExit(r, respData.Code, "用户信息验证失败")
	}
}

// 后台退出登陆
func (c *auth) loginOut(r *ghttp.Request) bool {
	//删除在线用户状态
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" && parts[1] != "" {
			//删除在线用户状态操作
			service.Online.DeleteOnlineByToken(parts[1])
		}
	}
	authHeader = r.GetString("token")
	if authHeader != "" {
		//删除在线用户状态操作
		service.Online.DeleteOnlineByToken(authHeader)
	}
	return true
}

// CheckUserOnline 检查在线用户
func (c *auth) CheckUserOnline() {
	param := &model.SysUserOnlineSearchReq{
		PageReq: comModel.PageReq{
			PageNum:  1,
			PageSize: 50,
		},
	}
	var total int
	for {
		var (
			list []*model.SysUserOnline
			err  error
		)
		total, _, list, err = service.Online.GetOnlineListPage(param, true)
		if err != nil {
			g.Log().Error(err)
			break
		}
		if list == nil {
			break
		}
		for _, v := range list {
			if b := c.UserIsOnline(v.Token); !b {
				service.Online.DeleteOnlineByToken(v.Token)
			}
		}
		if param.PageNum*param.PageSize >= total {
			break
		}
		param.PageNum++
	}
}

// UserIsOnline 判断用户是否在线
func (c *auth) UserIsOnline(token string) bool {
	uuid, userKey := c.GetUuidUserKeyByToken(token)
	cacheKey := GfToken.CacheKey + userKey
	switch GfToken.CacheMode {
	case gtoken.CacheModeCache:
		userCacheValue, _ := gcache.Get(cacheKey)
		if userCacheValue == nil {
			return false
		}
		return true
	case gtoken.CacheModeRedis:
		var userCache g.Map
		userCacheJson, err := g.Redis().Do("GET", cacheKey)
		if err != nil {
			g.Log().Error("[GToken]cache get error", err)
			return false
		}
		if userCacheJson == nil {
			return false
		}
		err = gjson.DecodeTo(userCacheJson, &userCache)
		if err != nil {
			g.Log().Error("[GToken]cache get json error", err)
			return false
		}
		if uuid != userCache["uuid"] {
			return false
		}
		return true
	}
	return false
}

// GetUuidUserKeyByToken 通过token获取uuid和userKey
func (c *auth) GetUuidUserKeyByToken(token string) (uuid, userKey string) {
	decryptToken := GfToken.DecryptToken(token)
	if !decryptToken.Success() {
		return
	}
	userKey = decryptToken.GetString("userKey")
	uuid = decryptToken.GetString("uuid")
	return
}
