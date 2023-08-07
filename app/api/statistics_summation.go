package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// StatisticsSummation 总量统计API
var StatisticsSummation = new(statisticsSummationApi)

type statisticsSummationApi struct{}

// GetInspire SignUp @summary 激励总量汇总
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/summation/inspire [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetInspire(r *ghttp.Request) {
	res, err := service.StatisticsSummation.GetInspire(r.Context())
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetStage SignUp @summary 激励阀点汇总
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/summation/stage [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetStage(r *ghttp.Request) {
	res, err := service.StatisticsSummation.GetStage(r.Context())
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProductStage SignUp @summary 项目阶段统计
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/product/stage [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetProductStage(r *ghttp.Request) {
	res, err := service.StatisticsSummation.GetProductStage(r.Context())
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProductStageScore SignUp @summary 项目阶段统计
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/product/score [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetProductStageScore(r *ghttp.Request) {
	res, err := service.StatisticsSummation.GetProductStageScore(r.Context())
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProductStageTop SignUp @summary 项目阶段TOP排名
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/product/top [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetProductStageTop(r *ghttp.Request) {
	res, err := service.StatisticsSummation.GetProductStageTop(r.Context())
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProductMemberLevel SignUp @summary 绩效等级统计
// @tags    总量统计
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /statistics/level/index [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *statisticsSummationApi) GetProductMemberLevel(r *ghttp.Request) {
	var input *model.ProductMemberStaticWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	res, err := service.StatisticsSummation.GetProductMemberLevel(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}
