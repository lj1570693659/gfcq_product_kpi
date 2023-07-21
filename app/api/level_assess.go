package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// LevelAssess 项目等级评估信息API管理对象
var LevelAssess = new(levelAssessApi)

type levelAssessApi struct{}

// GetOne SignUp @summary 获取部门信息详情
// @tags    项目等级评估配置
// @produce json
// @param   entity  body model.LevelAssess true "注册请求"
// @router  /config/product/assess/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelAssessApi) GetOne(r *ghttp.Request) {
	var input *model.LevelAssess

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.LevelAssess.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetList SignUp @summary 获取部门信息列表
// @tags    项目等级评估配置
// @produce json
// @param   entity  body model.LevelAssessApiGetList true "注册请求"
// @router  /config/product/assess/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelAssessApi) GetList(r *ghttp.Request) {
	var input *model.LevelAssessApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.LevelAssess.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    项目等级评估配置
// @produce json
// @param   entity  body model.LevelAssessApiChangeReq true "创建部门"
// @router  /config/product/assess/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelAssessApi) Create(r *ghttp.Request) {
	var input *model.LevelAssessApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.LevelAssess.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    项目等级评估配置
// @produce json
// @param   entity  body model.LevelAssessApiChangeReq true "注册请求"
// @router  /config/product/assess/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelAssessApi) Modify(r *ghttp.Request) {
	var input *model.LevelAssessApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelAssess, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailLevelAssess, "编辑对象数据丢失")
	}

	if err := service.LevelAssess.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    项目等级评估配置
// @produce json
// @param   entity  body model.LevelAssessApiDeleteReq true "删除请求"
// @router  /config/product/assess/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelAssessApi) Delete(r *ghttp.Request) {
	var input *model.LevelAssessApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelAssess, err.Error())
	}

	if err := service.LevelAssess.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailLevelAssess, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
