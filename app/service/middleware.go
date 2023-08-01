package service

import (
	"fmt"
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
	fmt.Println("----------------------")
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{}
		customCtx.User.UserInfo = &model.UserInfo{
			Id:         gconv.Uint(user.Id),
			WorkNumber: user.WorkNumber,
			Password:   user.Password,
		}

		// 完善上下文员工信息
		employeeInfo, err := Employee.GetOne(r.Context(), &model.EmployeeApiGetOneReq{
			model.Employee{
				WorkNumber: user.WorkNumber,
			},
		})

		if err != nil {
			response.JsonExit(r, http.StatusForbidden, err.Error())
		}
		Context.SetUserEmployee(r.Context(), employeeInfo.EmployeeInfo)
		Context.SetUserDepartment(r.Context(), employeeInfo.DepartmentInfo)
		Context.SetUserJob(r.Context(), employeeInfo.JobInfo)
		Context.SetUserProduct(r.Context(), employeeInfo.ProductInfo)
	}
	fmt.Println("----33333333------------------")
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// LoggedIn 鉴权中间件，验证是否登录
func (s *middlewareService) LoggedIn(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		fmt.Println("islogin------------------")
		r.Middleware.Next()
	} else {
		fmt.Println("islogin----------------false--")
		response.JsonExit(r, http.StatusForbidden, "")
	}
}

// Role 鉴权中间件，验证是否在允许角色组内
func (s *middlewareService) Role(r *ghttp.Request) {
	ok, err := Casbin.CheckAuth(r.Context(), Context.Get(r.Context()).User, r, ROLE)
	fmt.Println("Role----------------false--", ok, err)
	if err != nil {
		response.JsonExit(r, http.StatusForbidden, err.Error())
	}
	if ok {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, http.StatusForbidden, err.Error())
	}
}

// BusinessRole 鉴权中间件，验证是否在项目组内 TODO
func (s *middlewareService) BusinessRole(r *ghttp.Request) {
	ok, err := Casbin.CheckAuth(r.Context(), Context.Get(r.Context()).User, r, BUSINESS_ROLE)
	fmt.Println("BusinessRole----------------false--", ok, err)
	if err != nil {
		response.JsonExit(r, http.StatusForbidden, err.Error())
	}
	if ok {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, http.StatusForbidden, err.Error())
	}
}

// CORS 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
