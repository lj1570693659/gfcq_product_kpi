package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// JobLevel 职级信息API管理对象
var JobLevel = new(jobLevelApi)

type jobLevelApi struct{}

// GetOne SignUp @summary 获取部门信息详情
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevel true "注册请求"
// @router  /system/organize/level/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) GetOne(r *ghttp.Request) {
	var input *model.JobLevel

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.JobLevel.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetList SignUp @summary 获取部门信息列表
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevelApiGetListReq true "注册请求"
// @router  /system/organize/level/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) GetList(r *ghttp.Request) {
	var input *model.JobLevelApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.JobLevel.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetAll SignUp @summary 获取部门信息列表
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevelApiGetListReq true "注册请求"
// @router  /system/organize/level/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) GetAll(r *ghttp.Request) {
	var input *model.JobLevel

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.JobLevel.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevelApiChangeReq true "创建部门"
// @router  /system/organize/level/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) Create(r *ghttp.Request) {
	var input *model.JobLevelApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.JobLevel.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevelApiChangeReq true "注册请求"
// @router  /system/organize/level/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) Modify(r *ghttp.Request) {
	var input *model.JobLevelApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.JobLevelFailEmployee, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.JobLevelFailEmployee, "编辑对象数据丢失")
	}

	if err := service.JobLevel.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    职级信息服务
// @produce json
// @param   entity  body model.JobLevelApiDeleteReq true "删除请求"
// @router  /system/organize/level/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *jobLevelApi) Delete(r *ghttp.Request) {
	var input *model.JobLevelApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.JobLevelFailEmployee, err.Error())
	}

	if err := service.JobLevel.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
