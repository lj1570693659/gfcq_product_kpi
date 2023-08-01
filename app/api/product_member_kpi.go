package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductMemberKpi 项目小组成员绩效录入
var ProductMemberKpi = new(productMemberKpiApi)

type productMemberKpiApi struct{}

// Export SignUp @summary 团队成员信息导出
// @tags    团队成员绩效
// @produce json
// @param   entity  body model.ProductMemberExport true "团队成员信息导出"
// @router  /achieve/product/member/export [POST]
// @success 200 {object} response.JsonResponse "团队绩效"
func (a *productMemberKpiApi) Export(r *ghttp.Request) {
	var input *model.ProductMemberExport

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKpi, err.Error())
	}

	filepath, err := service.ProductMemberKpi.Export(r.Context(), input)

	if err != nil {
		response.JsonExit(r, response.CreateFailProductStageKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", filepath)
	}
}

// Import SignUp @summary 团队成员绩效导入
// @tags    团队成员绩效
// @produce json
// @param   entity  body model.ProductMemberKpiImportReq true "团队成员绩效导入"
// @router  /achieve/product/member/import [POST]
// @success 200 {object} response.JsonResponse "团队绩效"
func (a *productMemberKpiApi) Import(r *ghttp.Request) {
	var input *model.ProductMemberKpiImportReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKpi, err.Error())
	}

	err := service.ProductMemberKpi.Import(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailProductMemberKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// Create SignUp @summary 增加团队成员绩效
// @tags    团队成员绩效
// @produce json
// @param   entity  body model.ProductMemberKpiChangeReq true "增加团队成员绩效"
// @router  /achieve/product/member/create [POST]
// @success 200 {object} response.JsonResponse "团队绩效"
func (a *productMemberKpiApi) Create(r *ghttp.Request) {
	var input model.ProductMemberKpiChangeReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKpi, err.Error())
	}

	err := service.ProductMemberKpi.Create(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailProductMemberKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// Modify SignUp @summary 编辑团队成员绩效
// @tags    团队成员绩效
// @produce json
// @param   entity  body model.ProductMemberKpiChangeReq true "编辑团队成员绩效"
// @router  /achieve/product/member/modify [PUT]
// @success 200 {object} response.JsonResponse "团队绩效"
func (a *productMemberKpiApi) Modify(r *ghttp.Request) {
	var input model.ProductMemberKpiChangeReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKpi, err.Error())
	}

	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailProductMemberKpi, "编辑对象数据丢失")
	}

	err := service.ProductMemberKpi.Modify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.ModifyFailProductMemberKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// GetList SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiGetListReq true "项目阶段绩效清单"
// @router  /achieve/product/member/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberKpiApi) GetList(r *ghttp.Request) {
	var input model.ProductMemberKpiApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKpi, err.Error())
	}

	res, err := service.ProductMemberKpi.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberKpi, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
