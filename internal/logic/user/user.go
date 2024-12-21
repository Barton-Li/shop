package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/do"
	"shop/internal/service"
	"shop/utility"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}
func init() {
	service.RegisterUser(New())
}
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}
	userSlat := grand.S(10)
	in.UserSalt = userSlat
	in.Password = utility.EncryptPassword(in.Password, userSlat)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, nil
}
