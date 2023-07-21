package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewSolveRule 解决问题贡献
var CrewSolveRule = new(crewSolveRuleApi)

type crewSolveRuleApi struct{}

// GetAll SignUp @summary 获取解决问题贡献信息列表
// @tags    解决问题贡献
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/solve/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewSolveRuleApi) GetAll(r *ghttp.Request) {
	var input *model.CrewSolveRule

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailSolveRule, err.Error())
	}

	res, err := service.CrewSolveRule.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善解决问题贡献信息
// @tags    解决问题贡献
// @produce json
// @param   entity  body model.CrewSolveRuleApiChangeReq true "创建解决问题贡献"
// @router  /config/inspirit/solve/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewSolveRuleApi) Create(r *ghttp.Request) {
	var input *model.CrewSolveRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailSolveRule, err.Error())
	}

	if err := service.CrewSolveRule.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailSolveRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新解决问题贡献信息
// @tags    解决问题贡献
// @produce json
// @param   entity  body model.CrewSolveRuleApiChangeReq true "注册请求"
// @router  /config/inspirit/solve/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewSolveRuleApi) Modify(r *ghttp.Request) {
	var input *model.CrewSolveRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailSolveRule, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailSolveRule, "编辑对象数据丢失")
	}

	if err := service.CrewSolveRule.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailSolveRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除解决问题贡献信息
// @tags    解决问题贡献
// @produce json
// @param   entity  body model.CrewSolveRuleApiDeleteReq true "删除请求"
// @router  /config/inspirit/solve/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewSolveRuleApi) Delete(r *ghttp.Request) {
	var input *model.CrewSolveRuleApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailSolveRule, err.Error())
	}

	if err := service.CrewSolveRule.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailSolveRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
