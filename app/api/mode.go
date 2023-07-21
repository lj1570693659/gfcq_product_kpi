package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Mode 项目类型信息API管理对象
var Mode = new(modeApi)

type modeApi struct{}

// GetAll GetList SignUp @summary 获取部门信息列表
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.Mode true "注册请求"
// @router  /config/product/mode/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *modeApi) GetAll(r *ghttp.Request) {
	var input *model.Mode

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}

	res, err := service.Mode.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiChangeReq true "创建部门"
// @router  /config/product/mode/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *modeApi) Create(r *ghttp.Request) {
	var input *model.ModeApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}

	if err := service.Mode.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiChangeReq true "注册请求"
// @router  /config/product/mode/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *modeApi) Modify(r *ghttp.Request) {
	var input *model.ModeApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailMode, "编辑对象数据丢失")
	}

	if err := service.Mode.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除部门信息
// @tags    项目类型信息服务
// @produce json
// @param   entity  body model.ModeApiDeleteReq true "删除请求"
// @router  /config/product/mode/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *modeApi) Delete(r *ghttp.Request) {
	var input *model.ModeApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailMode, err.Error())
	}

	if err := service.Mode.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailMode, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
