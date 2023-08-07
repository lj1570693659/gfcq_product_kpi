package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
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

// CreateModeStage GetStageAll GetList SignUp @summary 获取项目阶段列表
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

// ModifyModeStage GetStageAll GetList SignUp @summary 获取项目阶段列表
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

// DeleteModeStage GetStageAll GetList SignUp @summary 获取项目阶段列表
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

// Create SignUp @summary 完善部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiChangeReq true "创建部门"
// @router  /config/product/type/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) Create(r *ghttp.Request) {
	var input *model.TypeApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}

	if _, err := service.Type.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiChangeReq true "注册请求"
// @router  /config/product/type/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) Modify(r *ghttp.Request) {
	var input *model.TypeApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailMode, "编辑对象数据丢失")
	}

	if _, err := service.Type.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiDeleteReq true "删除请求"
// @router  /config/product/type/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *typeApi) Delete(r *ghttp.Request) {
	var input *model.TypeApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}

	if _, err := service.Type.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
