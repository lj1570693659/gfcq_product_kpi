package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"net/http"
)

// ProductMember 项目成员信息API管理对象
var ProductMember = new(productMemberApi)

type productMemberApi struct{}

// GetList SignUp @summary 项目团队成员清单
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberGetListReq true "项目团队成员清单"
// @router  /product/member/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productMemberApi) GetList(r *ghttp.Request) {
	var input *model.ProductMemberGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.ProductMember.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetOne @summary 项目主信息详情
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiGetOneReq true "项目详情"
// @router  /product/member/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) GetOne(r *ghttp.Request) {
	var input *model.ProductMemberApiGetOneReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.ProductMember.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetOneFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create @summary 创建项目基础信息
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiChangeReq true "注册请求"
// @router  /product/member/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) Create(r *ghttp.Request) {
	var input *model.ProductMemberApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	// 验证PM权限
	if ok := service.Casbin.CheckProductAuth(r.Context(), input.ProId); !ok {
		response.JsonExit(r, http.StatusForbidden, "用户权限不足")
	}
	if out, err := service.ProductMember.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailProduct, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Modify @summary 更新项目基础信息
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiChangeReq true "注册请求"
// @router  /product/member/modify [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) Modify(r *ghttp.Request) {
	var input *model.ProductMemberApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	// 验证PM权限
	if ok := service.Casbin.CheckProductAuth(r.Context(), input.ProId); !ok {
		response.JsonExit(r, http.StatusForbidden, "用户权限不足")
	}

	if out, err := service.ProductMember.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Import @summary 导入团队成员信息
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiImportReq true "注册请求"
// @router  /product/member/import [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) Import(r *ghttp.Request) {
	var input *model.ProductMemberApiImportReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMember, err.Error())
	}

	// 验证PM权限
	if ok := service.Casbin.CheckProductAuth(r.Context(), input.ProId); !ok {
		response.JsonExit(r, http.StatusForbidden, "用户权限不足")
	}

	if out, err := service.ProductMember.Import(r.Context(), input); err != nil {
		response.JsonExit(r, response.FormatFailProductMember, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}
}

// WebImport @summary 网页导入团队成员信息
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiImportReq true "注册请求"
// @router  /product/member/import [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) WebImport(r *ghttp.Request) {
	var input *model.ProductMemberApiWebImportReq
	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductMember, err.Error())
	}

	// 验证PM权限
	if ok := service.Casbin.CheckProductAuth(r.Context(), input.ProId); !ok {
		response.JsonExit(r, http.StatusForbidden, "用户权限不足")
	}

	if out, err := service.ProductMember.WebImport(r.Context(), input); err != nil {
		response.JsonExit(r, response.FormatFailProductMember, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}
}

// Delete @summary 创建项目基础信息
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiChangeReq true "注册请求"
// @router  /product/member/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) Delete(r *ghttp.Request) {
	var input *model.ProductMemberApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	// 验证PM权限
	if ok := service.Casbin.CheckProductAuth(r.Context(), input.ProId); !ok {
		response.JsonExit(r, http.StatusForbidden, "用户权限不足")
	}
	if out, err := service.ProductMember.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailProduct, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Export @summary 导出项目组成员
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductMemberApiChangeReq true "注册请求"
// @router  /product/member/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productMemberApi) Export(r *ghttp.Request) {
	var input *model.ProductMemberWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	url, err := service.ProductMember.Export(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", url)
	}

}
