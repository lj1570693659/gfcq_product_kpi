package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// BudgetAccess 项目激励预算API管理对象
var BudgetAccess = new(budgetAccessApi)

type budgetAccessApi struct{}

// GetAll SignUp @summary 获取激励预算信息列表
// @tags    项目激励预算
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/budget/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetAccessApi) GetAll(r *ghttp.Request) {
	var input *model.BudgetAccess

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.BudgetAccess.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetList SignUp @summary 获取激励预算信息列表
// @tags    项目激励预算
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/budget/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetAccessApi) GetList(r *ghttp.Request) {
	var input *model.BudgetAccessApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.BudgetAccess.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    项目激励预算
// @produce json
// @param   entity  body model.BudgetAccessApiChangeReq true "创建激励预算"
// @router  /config/inspirit/budget/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetAccessApi) Create(r *ghttp.Request) {
	var input *model.BudgetAccessApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if err := service.BudgetAccess.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    项目激励预算
// @produce json
// @param   entity  body model.BudgetAccessApiChangeReq true "注册请求"
// @router  /config/inspirit/budget/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetAccessApi) Modify(r *ghttp.Request) {
	var input *model.BudgetAccessApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelConfirm, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailLevelConfirm, "编辑对象数据丢失")
	}

	if err := service.BudgetAccess.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    项目激励预算
// @produce json
// @param   entity  body model.BudgetAccessApiDeleteReq true "删除请求"
// @router  /config/inspirit/budget/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetAccessApi) Delete(r *ghttp.Request) {
	var input *model.BudgetAccessApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailLevelConfirm, err.Error())
	}

	if err := service.BudgetAccess.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailLevelConfirm, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
