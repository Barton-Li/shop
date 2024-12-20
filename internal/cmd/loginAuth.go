package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model/entity"
	"shop/utility"
	"shop/utility/response"
	"strconv"
)

func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:        consts.CacheModeRedis,
		ServerName:       consts.BackendServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/backend/admin/info"},
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"},
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	err = gfAdminToken.Start()
	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	return consts.GtokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}

	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	// 定义一个变量用于存储从响应数据中获取的管理员信息
	var adminInfo entity.AdminInfo
	// 尝试将响应数据中的字符串转换为管理员信息结构体
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	// 如果转换失败，返回一个错误响应
	if err != nil {
		response.Auth(r)
		return
	}
	// 将管理员ID设置到请求的上下文中
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	// 将管理员名称设置到请求的上下文中
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	// 将管理员是否为管理员的标志设置到请求的上下文中
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	// 将管理员角色ID列表设置到请求的上下文中
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	// 调用下一个中间件
	r.Middleware.Next()
}
