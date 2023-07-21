package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewHoursIndex 工时指数
var CrewHoursIndex = new(crewHoursIndexApi)

type crewHoursIndexApi struct{}

// GetAll SignUp @summary 获取激励预算信息列表
// @tags    工时指数
// @produce json
// @param   entity  body model.CrewHoursIndex true "注册请求"
// @router  /config/inspirit/hours/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewHoursIndexApi) GetAll(r *ghttp.Request) {
	var input *model.CrewHoursIndex

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	res, err := service.CrewHoursIndex.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    工时指数
// @produce json
// @param   entity  body model.CrewHoursIndexApiChangeReq true "创建激励预算"
// @router  /config/inspirit/hours/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewHoursIndexApi) Create(r *ghttp.Request) {
	var input *model.CrewHoursIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.CrewHoursIndex.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    工时指数
// @produce json
// @param   entity  body model.CrewHoursIndexApiChangeReq true "注册请求"
// @router  /config/inspirit/hours/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewHoursIndexApi) Modify(r *ghttp.Request) {
	var input *model.CrewHoursIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailStageRadio, "编辑对象数据丢失")
	}

	if err := service.CrewHoursIndex.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    工时指数
// @produce json
// @param   entity  body model.CrewHoursIndexApiDeleteReq true "删除请求"
// @router  /config/inspirit/hours/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewHoursIndexApi) Delete(r *ghttp.Request) {
	var input *model.CrewHoursIndexApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.CrewHoursIndex.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
