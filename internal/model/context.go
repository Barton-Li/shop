package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
	Data    g.Map
}
type ContextUser struct {
	Id      int
	Name    string
	IsAdmin int
}
