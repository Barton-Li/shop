package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	// 初始化一个 AdminInfo 结构体指针 adminInfo
	adminInfo := new(entity.AdminInfo)
	// 在数据库中查询用户名等于输入用户名的记录，并将结果扫描到 adminInfo 中
	err := dao.AddressInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	// 如果查询过程中发生错误，返回错误
	if err != nil {
		return err
	}
	// 对比输入的密码经过加密后的结果是否与数据库中存储的密码一致
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		// 如果不一致，返回一个包含错误信息的新错误
		return gerror.New("账号或密码错误")
	}
	// 将用户信息存储在会话中
	if err := service.Session().SetUser(ctx, adminInfo); err != nil {
		// 如果存储过程中发生错误，返回错误
		return err
	}
	// 在上下文中设置用户信息
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      adminInfo.Id,
		Name:    adminInfo.Name,
		IsAdmin: adminInfo.IsAdmin,
	})
	// 返回 nil，表示登录成功
	return nil
}

func (s *sLogin) Logout(ctx context.Context) error {
	// 从会话中移除用户信息
	return service.Session().RemoveUser(ctx)
}
