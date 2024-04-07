package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"GoFrameExample/internal/controller/user"
	"GoFrameExample/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/user", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewUser(),
				)

				group.Group("/profile", func(group1 *ghttp.RouterGroup) {
					group1.Middleware(ghttp.MiddlewareHandlerResponse)
					group1.Middleware(middleware.LoginCheck, middleware.RouterCheck)
					group1.Bind(
						user.NewProfile(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
