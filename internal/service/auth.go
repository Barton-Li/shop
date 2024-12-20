package service

import (
	"context"
	"shop/internal/model"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}
func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "shop",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 60,
		MaxRefresh:      time.Minute * 60,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	if user := Admin().GetAdminByNamePassword(ctx, in); user != nil {
		return user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}
