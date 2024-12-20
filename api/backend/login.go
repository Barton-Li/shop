package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"shop/internal/model/entity"
	"time"
)

type LoginDoReq struct {
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// jwt
type LoginDoRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// for gtoken
type LoginRes struct {
	Type       string                  `json:"type"`
	Token      string                  `json:"token"`
	ExpireIn   int                     `json:"expire_in"`
	IsAdmin    bool                    `json:"is_admin"`
	RoleIds    string                  `json:"role_ids"`
	Permission []entity.PermissionInfo `json:"permission"`
}
type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post" summary:"刷新token" tags:"登录"`
}
type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"退出登录" tags:"登录"`
}
type LogoutRes struct{}
