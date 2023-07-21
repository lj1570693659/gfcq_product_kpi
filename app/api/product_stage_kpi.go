package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductStageKpi 项目阶段总绩效信息管理对象
var ProductStageKpi = new(productStageKpiApi)

type productStageKpiApi struct{}

// Create SignUp @summary 项目阶段绩效录入
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiChangeReq true "项目阶段绩效录入"
// @router  /achieve/product/create [POST]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productStageKpiApi) Create(r *ghttp.Request) {
	var input *model.ProductStageKpiApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductStageKpi, err.Error())
	}

	if out, err := service.ProductStageKpi.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailProductStageKpi, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Modify SignUp @summary 项目阶段绩效调整
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiChangeReq true "项目阶段绩效调整"
// @router  /achieve/product/modify [PUT]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productStageKpiApi) Modify(r *ghttp.Request) {
	var input *model.ProductStageKpiApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductStageKpi, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailProductStageKpi, "编辑对象数据丢失")
	}

	if err := service.ProductStageKpi.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailProductStageKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// GetList SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiGetListReq true "项目阶段绩效清单"
// @router  /achieve/product/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productStageKpiApi) GetList(r *ghttp.Request) {
	var input *model.ProductStageKpiApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductStageKpi, err.Error())
	}

	res, err := service.ProductStageKpi.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductStageKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetOne SignUp @summary 项目阶段绩效详情
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpi true "项目阶段绩效详情"
// @router  /achieve/product/info [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productStageKpiApi) GetOne(r *ghttp.Request) {
	var input *model.ProductStageKpi

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductStageKpi, err.Error())
	}

	res, err := service.ProductStageKpi.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetOneFailProductStageKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
