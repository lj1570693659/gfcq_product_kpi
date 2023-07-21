package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewOvertimeRule 加班贡献
var CrewOvertimeRule = new(crewOvertimeRuleApi)

type crewOvertimeRuleApi struct{}

// GetAll SignUp @summary 获取激励预算信息列表
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewOvertimeRule true "注册请求"
// @router  /config/inspirit/overtime/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewOvertimeRuleApi) GetAll(r *ghttp.Request) {
	var input *model.CrewOvertimeRule

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailOvertimeRule, err.Error())
	}

	res, err := service.CrewOvertimeRule.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewOvertimeRuleApiChangeReq true "创建激励预算"
// @router  /config/inspirit/overtime/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewOvertimeRuleApi) Create(r *ghttp.Request) {
	var input *model.CrewOvertimeRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailOvertimeRule, err.Error())
	}

	if err := service.CrewOvertimeRule.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailOvertimeRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewOvertimeRuleApiChangeReq true "注册请求"
// @router  /config/inspirit/overtime/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewOvertimeRuleApi) Modify(r *ghttp.Request) {
	var input *model.CrewOvertimeRuleApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailOvertimeRule, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailOvertimeRule, "编辑对象数据丢失")
	}

	if err := service.CrewOvertimeRule.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailOvertimeRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    加班贡献
// @produce json
// @param   entity  body model.CrewOvertimeRuleApiDeleteReq true "删除请求"
// @router  /config/inspirit/overtime/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewOvertimeRuleApi) Delete(r *ghttp.Request) {
	var input *model.CrewOvertimeRuleApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailOvertimeRule, err.Error())
	}

	if err := service.CrewOvertimeRule.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailOvertimeRule, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
