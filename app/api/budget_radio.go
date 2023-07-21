package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// BudgetRadio 项目激励应发比例API管理对象
var BudgetRadio = new(budgetRadioApi)

type budgetRadioApi struct{}

// GetAll SignUp @summary 获取激励预算信息列表
// @tags    项目激励应发
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/radio/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetRadioApi) GetAll(r *ghttp.Request) {
	var input *model.BudgetRadio

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	res, err := service.BudgetRadio.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    项目激励应发
// @produce json
// @param   entity  body model.BudgetAccessApiChangeReq true "创建激励预算"
// @router  /config/inspirit/radio/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetRadioApi) Create(r *ghttp.Request) {
	var input *model.BudgetRadioApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.BudgetRadio.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    项目激励应发
// @produce json
// @param   entity  body model.BudgetAccessApiChangeReq true "注册请求"
// @router  /config/inspirit/radio/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetRadioApi) Modify(r *ghttp.Request) {
	var input *model.BudgetRadioApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailStageRadio, "编辑对象数据丢失")
	}

	if err := service.BudgetRadio.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    项目激励应发
// @produce json
// @param   entity  body model.BudgetAccessApiDeleteReq true "删除请求"
// @router  /config/inspirit/radio/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *budgetRadioApi) Delete(r *ghttp.Request) {
	var input *model.BudgetRadioApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.BudgetRadio.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
