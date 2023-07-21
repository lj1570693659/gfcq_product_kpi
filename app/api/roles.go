package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductRoles 项目角色信息API管理对象
var ProductRoles = new(productRolesApi)

type productRolesApi struct{}

// GetList SignUp @summary 获取项目角色信息列表
// @tags    项目角色信息服务
// @produce json
// @param   entity  body model.ProductRolesApiGetListReq true "注册请求"
// @router   /config/product/roles/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productRolesApi) GetList(r *ghttp.Request) {
	var input *model.ProductRolesApiGetListReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductRoles, err.Error())
	}

	res, err := service.ProductRoles.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善项目角色信息
// @tags    项目角色信息服务
// @produce json
// @param   entity  body model.ProductRolesApiChangeReq true "创建项目角色"
// @router   /config/product/roles/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productRolesApi) Create(r *ghttp.Request) {
	var input *model.ProductRolesApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductRoles, err.Error())
	}

	if err := service.ProductRoles.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailProductRoles, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除项目角色信息
// @tags    项目角色信息服务
// @produce json
// @param   entity  body model.ProductRolesApiDeleteReq true "删除请求"
// @router   /config/product/roles/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productRolesApi) Delete(r *ghttp.Request) {
	var input *model.ProductRolesApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductRoles, err.Error())
	}

	if err := service.ProductRoles.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailProductRoles, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新项目角色信息
// @tags    项目角色信息服务
// @produce json
// @param   entity  body model.ProductRolesApiChangeReq true "注册请求"
// @router   /config/product/roles/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *productRolesApi) Modify(r *ghttp.Request) {
	var input *model.ProductRolesApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProductRoles, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailProductRoles, "编辑对象数据丢失")
	}

	if err := service.ProductRoles.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailProductRoles, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
