package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Job 职级信息API管理对象
var Job = new(jobApi)

type jobApi struct{}

// GetOne SignUp @summary 获取部门信息详情
// @tags    职级信息服务
// @produce json
// @param   entity  body model.Job true "注册请求"
// @router  /system/organize/level/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *jobApi) GetOne(r *ghttp.Request) {
//	var input *model.Job
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailEmployee, err.Error())
//	}
//
//	res, err := service.Job.GetOne(r.Context(), input)
//	if err != nil {
//		response.JsonExit(r, response.CreateFailEmployee, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok", res)
//	}
//
//}
//
//// GetList SignUp @summary 获取部门信息列表
//// @tags    职级信息服务
//// @produce json
//// @param   entity  body model.JobApiGetListReq true "注册请求"
//// @router  /system/organize/level/lists [GET]
//// @success 200 {object} response.JsonResponse "执行结果"
//func (a *jobApi) GetList(r *ghttp.Request) {
//	var input *model.JobApiGetListReq
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailEmployee, err.Error())
//	}
//
//	res, err := service.Job.GetList(r.Context(), input)
//	if err != nil {
//		response.JsonExit(r, response.CreateFailEmployee, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok", res)
//	}
//
//}

// GetAll SignUp @summary 获取部门信息列表
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobApiGetListReq true "注册请求"
// @router  /system/organize/level/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobApi) GetAll(r *ghttp.Request) {
	var input *model.Job

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.Job.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobApiChangeReq true "创建部门"
// @router  /system/organize/level/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *jobApi) Create(r *ghttp.Request) {
//	var input *model.JobApiChangeReq
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailEmployee, err.Error())
//	}
//
//	if err := service.Job.Create(r.Context(), input); err != nil {
//		response.JsonExit(r, response.CreateFailEmployee, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//
//}

// Modify SignUp @summary 更新部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobApiChangeReq true "注册请求"
// @router  /system/organize/level/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *jobApi) Modify(r *ghttp.Request) {
//	var input *model.JobApiChangeReq
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.JobFailEmployee, err.Error())
//	}
//	if g.IsEmpty(input.ID) {
//		response.JsonExit(r, response.JobFailEmployee, "编辑对象数据丢失")
//	}
//
//	if err := service.Job.Modify(r.Context(), input); err != nil {
//		response.JsonExit(r, response.CreateFailEmployee, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//
//}

// Delete Modify SignUp @summary 删除部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobApiDeleteReq true "删除请求"
// @router  /system/organize/level/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *jobApi) Delete(r *ghttp.Request) {
//	var input *model.JobApiDeleteReq
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.JobFailEmployee, err.Error())
//	}
//
//	if err := service.Job.Delete(r.Context(), input); err != nil {
//		response.JsonExit(r, response.CreateFailEmployee, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//
//}
