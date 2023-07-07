package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/api"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		//group.ALL("/chat", api.Chat)

		// 登录账号相关
		group.ALL("/user", api.User)
		// 员工信息相关
		group.Group("/employee", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.GET("/isSyncEmployee", api.Employee.IsSyncEmployee)
			group.POST("/create", api.Employee.Create)
		})
	})
}
