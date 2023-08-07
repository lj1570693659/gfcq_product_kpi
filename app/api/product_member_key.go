package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductMemberKey 项目小组成员绩效录入
var ProductMemberKey = new(productMemberKeyApi)

type productMemberKeyApi struct{}

//
//// Export SignUp @summary 团队成员信息导出
//// @tags    团队成员绩效
//// @produce json
//// @param   entity  body model.ProductMemberExport true "团队成员信息导出"
//// @router  /achieve/product/member/export [POST]
//// @success 200 {object} response.JsonResponse "团队绩效"
//func (a *productMemberKeyApi) Export(r *ghttp.Request) {
//	var input *model.ProductMemberExport
//
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
//	}
//
//	filepath, err := service.ProductMemberKey.Export(r.Context(), input)
//
//	if err != nil {
//		response.JsonExit(r, response.CreateFailProductStageKpi, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok", filepath)
//	}
//}
//
//// Import SignUp @summary 团队成员绩效导入
//// @tags    团队成员绩效
//// @produce json
//// @param   entity  body model.ProductMemberKeyImportReq true "团队成员绩效导入"
//// @router  /achieve/product/member/import [POST]
//// @success 200 {object} response.JsonResponse "团队绩效"
//func (a *ProductMemberKeyApi) Import(r *ghttp.Request) {
//	var input *model.ProductMemberKeyImportReq
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
//	}
//
//	err := service.ProductMemberKey.Import(r.Context(), input)
//	if err != nil {
//		response.JsonExit(r, response.CreateFailProductMemberKey, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//}
//
//// Create SignUp @summary 增加团队成员绩效
//// @tags    团队成员绩效
//// @produce json
//// @param   entity  body model.ProductMemberKeyChangeReq true "增加团队成员绩效"
//// @router  /achieve/product/member/create [POST]
//// @success 200 {object} response.JsonResponse "团队绩效"
//func (a *ProductMemberKeyApi) Create(r *ghttp.Request) {
//	var input model.ProductMemberKeyChangeReq
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
//	}
//
//	err := service.ProductMemberKey.Create(r.Context(), input)
//	if err != nil {
//		response.JsonExit(r, response.CreateFailProductMemberKey, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//}
//
//// Modify SignUp @summary 编辑团队成员绩效
//// @tags    团队成员绩效
//// @produce json
//// @param   entity  body model.ProductMemberKeyChangeReq true "编辑团队成员绩效"
//// @router  /achieve/product/member/modify [PUT]
//// @success 200 {object} response.JsonResponse "团队绩效"
//func (a *ProductMemberKeyApi) Modify(r *ghttp.Request) {
//	var input model.ProductMemberKeyChangeReq
//	if err := r.Parse(&input); err != nil {
//		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
//	}
//
//	if g.IsEmpty(input.ID) {
//		response.JsonExit(r, response.FormatFailProductMemberKey, "编辑对象数据丢失")
//	}
//
//	err := service.ProductMemberKey.Modify(r.Context(), input)
//	if err != nil {
//		response.JsonExit(r, response.ModifyFailProductMemberKey, err.Error())
//	} else {
//		response.JsonExit(r, response.Success, "ok")
//	}
//}

// GetList SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiGetListReq true "项目阶段绩效清单"
// @router  /achieve/product/crucial/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberKeyApi) GetList(r *ghttp.Request) {
	var input model.ProductMemberKeyListsReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
	}

	res, err := service.ProductMemberKey.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberKey, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiGetListReq true "项目阶段绩效清单"
// @router  /achieve/product/crucial/create [POST]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberKeyApi) Create(r *ghttp.Request) {
	var input model.ProductMemberKeyApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
	}

	err := service.ProductMemberKey.Create(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberKey, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductStageKpiApiGetListReq true "项目阶段绩效清单"
// @router  /achieve/product/crucial/modify [POST]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberKeyApi) Modify(r *ghttp.Request) {
	var input model.ProductMemberKeyApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMemberKey, err.Error())
	}

	if g.IsEmpty(input.Id) {
		response.JsonExit(r, response.FormatFailStageRadio, "编辑对象数据丢失")
	}

	err := service.ProductMemberKey.Modify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberKey, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete SignUp @summary 项目阶段绩效清单
// @tags    项目绩效
// @produce json
// @param   entity  body model.ProductMemberKeyApiDeleteReq true "项目阶段绩效清单"
// @router  /achieve/product/crucial/delete [POST]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberKeyApi) Delete(r *ghttp.Request) {
	var input model.ProductMemberKeyApiDeleteReq

	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailStageRadio, "编辑对象数据丢失")
	}

	err := service.ProductMemberKey.Delete(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProductMemberKey, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
