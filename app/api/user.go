package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// 用户API管理对象
var User = new(userApi)

type userApi struct{}

// SignUp @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	}
	if err := service.User.SignUp(r.Context(), serviceReq); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// SignIn @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignIn(r *ghttp.Request) {
	var (
		data *model.UserApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	}
	if err := service.User.SignIn(r.Context(), data.WorkNumber, data.Password); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// IsSignedIn @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /user/issignedin [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (a *userApi) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, response.Success, "", service.User.IsSignedIn(r.Context()))
}

// SignOut @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /user/signout [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (a *userApi) SignOut(r *ghttp.Request) {
	if err := service.User.SignOut(r.Context()); err != nil {
		response.JsonExit(r, response.NotSignedIn, err.Error())
	}
	response.JsonExit(r, response.Success, "ok")
}

// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/profile [GET]
// @success 200 {object} model.User "用户信息"
func (a *userApi) Profile(r *ghttp.Request) {
	response.JsonExit(r, response.Success, "", service.User.GetProfile(r.Context()))
}
