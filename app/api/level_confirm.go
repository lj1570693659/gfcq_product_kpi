package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// LevelConfirm 项目优先级确认信息API管理对象
var LevelConfirm = new(levelConfirmApi)

type levelConfirmApi struct{}

// GetList SignUp @summary 获取部门信息列表
// @tags    项目优先级配置
// @produce json
// @param   entity  body model.LevelConfirmApiGetList true "注册请求"
// @router  /config/product/confirm/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelConfirmApi) GetList(r *ghttp.Request) {
	var input *model.LevelConfirmApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.LevelConfirm.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    项目优先级配置
// @produce json
// @param   entity  body model.LevelConfirmApiChangeReq true "创建部门"
// @router  /config/product/confirm/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelConfirmApi) Create(r *ghttp.Request) {
	var input *model.LevelConfirmApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.LevelConfirm.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    项目优先级配置
// @produce json
// @param   entity  body model.LevelConfirmApiChangeReq true "注册请求"
// @router  /config/product/confirm/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelConfirmApi) Modify(r *ghttp.Request) {
	var input *model.LevelConfirmApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelConfirm, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailLevelConfirm, "编辑对象数据丢失")
	}

	if err := service.LevelConfirm.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    项目优先级配置
// @produce json
// @param   entity  body model.LevelConfirmApiDeleteReq true "删除请求"
// @router  /config/product/confirm/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *levelConfirmApi) Delete(r *ghttp.Request) {
	var input *model.LevelConfirmApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelConfirm, err.Error())
	}

	if err := service.LevelConfirm.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailLevelConfirm, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
