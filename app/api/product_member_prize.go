package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductMemberPrize 项目小组成员激励录入
var ProductMemberPrize = new(productMemberPrizeApi)

type productMemberPrizeApi struct{}

// Export SignUp @summary 团队成员信息导出
// @tags    团队成员激励
// @produce json
// @param   entity  body model.ProductMemberExport true "团队成员信息导出"
// @router  /achieve/product/member/export [POST]
// @success 200 {object} response.JsonResponse "团队激励"
func (a *productMemberPrizeApi) Export(r *ghttp.Request) {
	//var input *model.ProductMemberExport
	//
	//if err := r.Parse(&input); err != nil {
	//	response.JsonExit(r, response.FormatFailProductMemberPrize, err.Error())
	//}
	//
	//filepath, err := service.ProductMemberPrize.Export(r.Context(), input)
	////r.Response.ServeFileDownload(filepath)
	//if err != nil {
	//	response.JsonExit(r, response.CreateFailProductStagePrize, err.Error())
	//} else {
	//	response.JsonExit(r, response.Success, "ok", filepath)
	//}
}

// Compute SignUp @summary 增加团队成员激励
// @tags    团队成员激励
// @produce json
// @param   entity  body model.ProductMemberPrizeChangeReq true "增加团队成员激励"
// @router  /achieve/product/prize/compute [POST]
// @success 200 {object} response.JsonResponse "团队激励"
func (a *productMemberPrizeApi) Compute(r *ghttp.Request) {
	var input *model.ProductMemberPrizeComputeReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberPrize, err.Error())
	}

	err := service.ProductMemberPrize.Compute(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailProductMemberPrize, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// Modify SignUp @summary 编辑团队成员激励
// @tags    团队成员激励
// @produce json
// @param   entity  body model.ProductMemberPrizeChangeReq true "编辑团队成员激励"
// @router  /achieve/product/member/modify [PUT]
// @success 200 {object} response.JsonResponse "团队激励"
func (a *productMemberPrizeApi) Modify(r *ghttp.Request) {
	//var input model.ProductMemberPrizeChangeReq
	//if err := r.Parse(&input); err != nil {
	//	response.JsonExit(r, response.FormatFailProductMemberPrize, err.Error())
	//}
	//
	//if g.IsEmpty(input.ID) {
	//	response.JsonExit(r, response.FormatFailProductMemberPrize, "编辑对象数据丢失")
	//}
	//
	//err := service.ProductMemberPrize.Modify(r.Context(), input)
	//if err != nil {
	//	response.JsonExit(r, response.ModifyFailProductMemberPrize, err.Error())
	//} else {
	//	response.JsonExit(r, response.Success, "ok")
	//}
}

// GetList SignUp @summary 项目阶段激励清单
// @tags    项目激励
// @produce json
// @param   entity  body model.ProductStagePrizeApiGetListReq true "项目阶段激励清单"
// @router  /achieve/product/member/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberPrizeApi) GetList(r *ghttp.Request) {
	var input model.ProductMemberPrizeApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberPrize, err.Error())
	}

	res, err := service.ProductMemberPrize.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberPrize, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
