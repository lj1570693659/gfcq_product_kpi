package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Product 员工信息API管理对象
var Product = new(productApi)

type productApi struct{}

// GetList SignUp @summary 项目清单
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productApi) GetList(r *ghttp.Request) {
	var input *model.ProductApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	ctx := r.Context()
	if service.Context.Get(ctx).User.RoleLevel == service.LevelLow {
		input.ProductWhere.Ids = service.Context.Get(ctx).User.ProductIds
	}

	res, err := service.Product.GetList(ctx, input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetAll SignUp @summary 项目筛选清单
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/all [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productApi) GetAll(r *ghttp.Request) {
	var input model.ProductWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.Product.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetOne @summary 项目主信息详情
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetOneReq true "项目详情"
// @router  /product/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productApi) GetOne(r *ghttp.Request) {
	var input *model.ProductApiGetOneReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.Product.GetOne(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetOneFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetDetail @summary 项目详情
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetOneReq true "项目详情"
// @router  /product/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productApi) GetDetail(r *ghttp.Request) {
	var input *model.ProductApiGetOneReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.Product.GetDetail(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetOneFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create @summary 创建项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiChangeReq true "注册请求"
// @router  /product/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productApi) Create(r *ghttp.Request) {
	var input *model.ProductApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	if out, err := service.Product.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailProduct, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Modify @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /product/modify [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productApi) Modify(r *ghttp.Request) {
	var input *model.ProductApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if out, err := service.Product.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}

// Delete @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /product/delete [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productApi) Delete(r *ghttp.Request) {
	var input *model.Product

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailEmployee, err.Error())
	}

	if out, err := service.Product.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error(), out)
	} else {
		response.JsonExit(r, response.Success, "ok", out)
	}

}
