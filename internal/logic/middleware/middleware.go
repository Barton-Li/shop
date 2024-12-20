package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility/response"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/backend/login",
	}
}
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	// 调用下一个中间件或处理函数
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	// 定义变量用于存储错误和响应数据
	var (
		// 从请求中获取错误信息
		err = r.GetError()
		// 从请求中获取处理函数的返回结果
		res = r.GetHandlerResponse()
		// 定义一个默认的响应状态码
		code gcode.Code = gcode.CodeOK
	)

	// 如果有错误发生
	if err != nil {
		// 获取错误的状态码
		code = gerror.Code(err)
		// 如果状态码为空，则设置为内部错误
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		// 以JSON格式返回错误信息
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		// 如果没有错误，以JSON格式返回响应数据
		response.JsonExit(r, code.Code(), "", res)
	}
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 创建一个新的 Context 对象
	customCtx := &model.Context{
		// 将请求的 Session 对象赋值给 Context 的 Session 属性
		Session: r.Session,
		// 创建一个新的 map 用于存储数据
		Data: make(g.Map),
	}
	// 初始化 BizCtx 服务，并将 customCtx 作为参数传递给 Init 方法
	service.BizCtx().Init(r, customCtx)

	// 从 Session 中获取用户信息
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		// 如果用户信息存在且 Id 大于 0，则创建一个新的 ContextUser 对象
		customCtx.User = &model.ContextUser{
			// 将用户实体的 Id 赋值给 ContextUser 的 Id 属性
			Id: userEntity.Id,
			// 将用户实体的 Name 赋值给 ContextUser 的 Name 属性
			Name: userEntity.Name,
			// 将用户实体的 IsAdmin 赋值给 ContextUser 的 IsAdmin 属性
			IsAdmin: userEntity.IsAdmin,
		}
	}

	// 将 customCtx 赋值给请求的上下文变量 "Context"
	r.Assigns(g.Map{
		"Context": customCtx,
	})

	// 调用下一个中间件或处理函数
	r.Middleware.Next()
}
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
