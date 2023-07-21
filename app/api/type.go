package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Type 项目类型信息API管理对象
var Type = new(typeApi)

type typeApi struct{}

// GetAll GetList SignUp @summary 获取部门信息列表
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.Type true "注册请求"
// @router  /config/product/type/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) GetAll(r *ghttp.Request) {
	var input *model.ProductType

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailType, err.Error())
	}

	res, err := service.Type.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailType, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetStageAll GetStageAll GetList SignUp @summary 获取项目阶段列表
// @tags    项目阶段信息服务
// @produce json
// @param   entity  body model.Type true "注册请求"
// @router  /config/product/stage/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) GetStageAll(r *ghttp.Request) {
	var input *model.ProductModeStage

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailType, err.Error())
	}

	res, err := service.Type.GetStageAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailType, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetStageAll GetStageAll GetList SignUp @summary 获取项目阶段列表
// @tags    项目阶段信息服务
// @produce json
// @param   entity  body model.Type true "注册请求"
// @router  /config/product/stage/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) CreateModeStage(r *ghttp.Request) {
	var input *model.ProductModeStage

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	}

	res, err := service.Type.CreateModeStage(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetStageAll GetStageAll GetList SignUp @summary 获取项目阶段列表
// @tags    项目阶段信息服务
// @produce json
// @param   entity  body model.Type true "注册请求"
// @router  /config/product/stage/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) ModifyModeStage(r *ghttp.Request) {
	var input *model.ProductModeStage

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	}

	res, err := service.Type.ModifyModeStage(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetStageAll GetStageAll GetList SignUp @summary 获取项目阶段列表
// @tags    项目阶段信息服务
// @produce json
// @param   entity  body model.Type true "注册请求"
// @router  /config/product/stage/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) DeleteModeStage(r *ghttp.Request) {
	var input *model.ProductModeStage

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	}

	res, err := service.Type.DeleteModeStage(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailModeStage, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
