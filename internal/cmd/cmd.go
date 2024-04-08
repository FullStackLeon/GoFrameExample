package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"GoFrameExample/internal/controller/user"
	"GoFrameExample/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 用户组：注册和登录相关
			s.Group("/user", func(userGroup *ghttp.RouterGroup) {
				userGroup.Middleware(ghttp.MiddlewareHandlerResponse)
				userGroup.Bind(
					user.NewUser(),
				)

				// Profile组：用户Profile相关
				userGroup.Group("/profile", func(profileGroup *ghttp.RouterGroup) {
					profileGroup.Middleware(
						ghttp.MiddlewareHandlerResponse,
						service.Middleware().LoginCheck,
						service.Middleware().RouterCheck,
					)
					profileGroup.Bind(
						user.NewProfile(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
