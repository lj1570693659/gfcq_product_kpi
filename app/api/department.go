package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Department 部门信息API管理对象
var Department = new(departmentApi)

type departmentApi struct{}

// GetOne SignUp @summary 获取部门信息详情
// @tags    部门信息服务
// @produce json
// @param   entity  body model.departmentApiGetOneReq true "注册请求"
// @router  /system/organize/department/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *departmentApi) GetOne(r *ghttp.Request) {
	var input *model.DepartmentApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Department.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetList SignUp @summary 获取部门信息列表
// @tags    部门信息服务
// @produce json
// @param   entity  body model.DepartmentApiGetListReq true "注册请求"
// @router  /system/organize/department/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *departmentApi) GetList(r *ghttp.Request) {
	var input *model.DepartmentApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Department.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    部门信息服务
// @produce json
// @param   entity  body model.DepartmentApiChangeReq true "创建部门"
// @router  /system/organize/department/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *departmentApi) Create(r *ghttp.Request) {
	var input *model.DepartmentApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.Department.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    部门信息服务
// @produce json
// @param   entity  body model.DepartmentApiDeleteReq true "删除请求"
// @router  /system/organize/department/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *departmentApi) Delete(r *ghttp.Request) {
	var input *model.DepartmentApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.DepartmentFailEmployee, err.Error())
	}

	if err := service.Department.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    部门信息服务
// @produce json
// @param   entity  body model.DepartmentApiChangeReq true "注册请求"
// @router  /system/organize/department/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *departmentApi) Modify(r *ghttp.Request) {
	var input *model.DepartmentApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.DepartmentFailEmployee, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.DepartmentFailEmployee, "编辑对象数据丢失")
	}

	if err := service.Department.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
