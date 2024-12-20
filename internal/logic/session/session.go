package session

import (
	"context"
	"github.com/gogf/gf/v2/util/gutil"
	"shop/internal/model/entity"
	"shop/internal/service"
)

type sSession struct{}

const (
	sessionKeyUser = "SessionKeyUser"
)

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}

}
func (s *sSession) SetUser(ctx context.Context, user *entity.AdminInfo) (err error) {
	// 打印 sessionKeyUser 常量的值
	gutil.Dump(sessionKeyUser)
	// 打印 user 变量的值
	gutil.Dump("user:", user)
	// 将用户信息存储在会话中
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyUser, user)
}

func (s *sSession) GetUser(ctx context.Context) *entity.AdminInfo {
	// 从上下文中获取自定义上下文对象
	customCtx := service.BizCtx().Get(ctx)
	// 如果自定义上下文对象存在
	if customCtx != nil {
		// 从会话中获取用户信息
		v, _ := customCtx.Session.Get(sessionKeyUser)
		// 如果用户信息存在且不为空
		if !v.IsNil() {
			// 创建一个 AdminInfo 类型的指针 user
			var user *entity.AdminInfo
			// 将获取到的用户信息解析并赋值给 user 变量
			_ = v.Struct(&user)
			// 返回获取到的用户信息
			return user
		}
	}
	// 如果自定义上下文对象不存在或者用户信息不存在，返回一个空的 AdminInfo 对象
	return &entity.AdminInfo{}
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	// 从上下文中获取自定义上下文对象
	customCtx := service.BizCtx().Get(ctx)
	// 如果自定义上下文对象存在
	if customCtx != nil {
		// 从会话中移除用户信息
		return customCtx.Session.Remove(sessionKeyUser)
	}
	// 如果自定义上下文对象不存在，返回 nil
	return nil
}
