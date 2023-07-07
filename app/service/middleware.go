package service

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"net/http"
)

// Middleware 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// Ctx 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:         gconv.Uint(user.Id),
			WorkNumber: user.WorkNumber,
			Password:   user.Password,
		}
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// Auth 鉴权中间件，只有登录成功之后才能通过
func (s *middlewareService) Auth(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, http.StatusForbidden, "")
	}
}

// CORS 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
