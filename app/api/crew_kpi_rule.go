package api

import (
	"github.com/gogf/gf/net/ghttp"
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
		response.JsonExit(r, response.FormatFailSolveRule, err.Error())
	}

	res, err := service.CrewKpiRule.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
