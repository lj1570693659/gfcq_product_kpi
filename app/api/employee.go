package api

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Employee 员工信息API管理对象
var Employee = new(employeeApi)

type employeeApi struct{}

// IsSyncEmployee @summary 判断是否已完善员工信息
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /system/organize/employee/isSyncEmployee [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) IsSyncEmployee(r *ghttp.Request) {
	if service.Employee.IsSyncEmployee(r.Context()) {
		response.JsonExit(r, response.Success, "已同步员工信息")
	} else {
		response.JsonExit(r, response.NotSyncEmployee, "请先完善员工信息")
	}

}

// GetOne @summary 获取员工信息详情
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.EmployeeApiGetOneReq true "注册请求"
// @router  /system/organize/employee/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) GetOne(r *ghttp.Request) {
	var input *model.EmployeeApiGetOneReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Employee.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.NotSyncEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetList @summary 获取员工信息列表
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.EmployeeApiGetListReq true "注册请求"
// @router  /system/organize/employee/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) GetList(r *ghttp.Request) {
	var input *model.EmployeeApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Employee.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetAll @summary 获取员工信息列表
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.Employee true "注册请求"
// @router  /system/organize/employee/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) GetAll(r *ghttp.Request) {
	var input *model.Employee

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Employee.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create @summary 完善员工信息
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.EmployeeApiCreateReq true "注册请求"
// @router  /system/organize/employee/create [POST]
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

// Modify @summary 更新员工信息
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /system/organize/employee/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) Modify(r *ghttp.Request) {
	var input *model.EmployeeApiModifyReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.Employee.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// GetCheckIn @summary 获取员工信息列表
// @tags    员工基础信息服务
// @produce json
// @param   entity  body model.Employee true "注册请求"
// @router  /system/organize/employee/checkIn [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *employeeApi) GetCheckIn(r *ghttp.Request) {
	var input *model.GetCheckIn

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}
	fmt.Println("-------------------------", input)
	res, err := service.Employee.GetCheckIn(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
