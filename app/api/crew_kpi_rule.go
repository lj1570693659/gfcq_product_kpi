package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewKpiRule 绩效等级对应比例
var CrewKpiRule = new(crewKpiRuleApi)

type crewKpiRuleApi struct{}

// GetAll SignUp @summary 获取解决问题贡献信息列表
// @tags    解决问题贡献
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/solve/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewKpiRuleApi) GetAll(r *ghttp.Request) {
	var input *model.CrewKpiRule

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailKpiRule, err.Error())
	}

	res, err := service.CrewKpiRule.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewKpiRuleApiChangeReq true "绩效等级评分"
// @router  /config/inspirit/kpiRule/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewKpiRuleApi) Create(r *ghttp.Request) {
	var input *model.CrewKpiRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailKpiRule, err.Error())
	}

	if err := service.CrewKpiRule.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailKpiRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewKpiRuleApiChangeReq true "绩效等级评分"
// @router  /config/inspirit/kpiRule/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewKpiRuleApi) Modify(r *ghttp.Request) {
	var input *model.CrewKpiRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailKpiRule, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailKpiRule, "编辑对象数据丢失")
	}

	if err := service.CrewKpiRule.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailKpiRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewKpiRuleApiDeleteReq true "删除请求"
// @router  /config/inspirit/kpiRule/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewKpiRuleApi) Delete(r *ghttp.Request) {
	var input *model.CrewKpiRuleApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailKpiRule, err.Error())
	}

	if err := service.CrewKpiRule.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailKpiRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
