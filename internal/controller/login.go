package controller

import (
	"context"
	"shop/api/backend"

	"shop/internal/service"
)

var Login = cLogin{}

type cLogin struct{}

//	func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//		// 初始化返回结果结构体
//		res = &backend.LoginDoRes{}
//		// 调用 service.Login().Login 方法进行登录操作
//		err = service.Login().Login(ctx, model.UserLoginInput{
//			Name:     req.Name,
//			Password: req.Password,
//		})
//		// 如果登录过程中发生错误，返回错误
//		if err != nil {
//			return
//		}
//		// 从会话中获取登录用户信息
//		loginUSer := service.Session().GetUser(ctx)
//		// 将获取到的登录用户信息设置到返回结果中
//		res.User = loginUSer
//		// 返回登录结果和错误信息
//		return
//	}
func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire, err = service.Auth().RefreshToken(ctx)
	if err != nil {
		return nil, err
	}
	return
}
func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
