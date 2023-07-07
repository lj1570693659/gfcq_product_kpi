package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// 员工信息API管理对象
var Employee = new(employeeApi)

type employeeApi struct{}

// SignUp @summary 判断是否已完善员工信息接口
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) IsSyncEmployee(r *ghttp.Request) {
	if service.Employee.IsSyncEmployee(r.Context()) {
		response.JsonExit(r, response.Success, "已同步员工信息")
	} else {
		response.JsonExit(r, response.NotSyncEmployee, "请先完善员工信息")
	}

}

// SignUp @summary 完善员工信息接口
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) Create(r *ghttp.Request) {
	var input *model.EmployeeApiCreateReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.Employee.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
