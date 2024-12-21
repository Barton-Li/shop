package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop/internal/consts"
	"shop/internal/controller"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//	不需要登录验证
				group.Bind(

					controller.Admin.Create,
					controller.Login,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Admin.List,   // 管理员
						controller.Admin.Update, // 管理员
						controller.Admin.Delete, // 管理员
						controller.Admin.Info,   // 查询当前管理员信息
						controller.Category,     // 栏目
						controller.Data,         // 数据统计
						controller.File,         //从0到1实现文件入库
						controller.Upload,       //实现可跨项目使用的文件上云工具类
						controller.Goods,        // 商品
						controller.Rotation,     // 轮播图
						controller.Role,         // 角色
						controller.GoodsOptions, // 商品规格
						controller.Permission,   // 权限
						controller.Position,     // 手工位
						controller.Order.List,   // 订单列表
						controller.Order.Detail, // 订单详情
					)
				})
			})
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.User.Register,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					group.Bind(
						controller.User.Info,
						controller.User.UpdatePassword,
						controller.Order.Add,
					)
				})
			})
			s.SetPort(8000)
			s.Run()
			return nil
		},
	}
)
