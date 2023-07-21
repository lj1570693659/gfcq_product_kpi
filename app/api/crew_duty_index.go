package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewDutyIndex 责任指数
var CrewDutyIndex = new(crewDutyIndexApi)

type crewDutyIndexApi struct{}

// GetAll SignUp @summary 获取责任指数信息列表
// @tags    责任指数
// @produce json
// @param   entity  body model.CrewDutyIndex true "注册请求"
// @router  /config/inspirit/duty/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewDutyIndexApi) GetAll(r *ghttp.Request) {
	var input *model.CrewDutyIndex

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailDutyIndex, err.Error())
	}

	res, err := service.CrewDutyIndex.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善责任指数信息
// @tags    责任指数
// @produce json
// @param   entity  body model.CrewDutyIndexApiChangeReq true "创建责任指数"
// @router  /config/inspirit/duty/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewDutyIndexApi) Create(r *ghttp.Request) {
	var input *model.CrewDutyIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailDutyIndex, err.Error())
	}

	if err := service.CrewDutyIndex.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailDutyIndex, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新责任指数信息
// @tags    责任指数
// @produce json
// @param   entity  body model.CrewDutyIndexApiChangeReq true "注册请求"
// @router  /config/inspirit/duty/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewDutyIndexApi) Modify(r *ghttp.Request) {
	var input *model.CrewDutyIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailDutyIndex, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailDutyIndex, "编辑对象数据丢失")
	}

	if err := service.CrewDutyIndex.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailDutyIndex, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除责任指数信息
// @tags    责任指数
// @produce json
// @param   entity  body model.CrewDutyIndexApiDeleteReq true "删除请求"
// @router  /config/inspirit/duty/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewDutyIndexApi) Delete(r *ghttp.Request) {
	var input *model.CrewDutyIndexApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailDutyIndex, err.Error())
	}

	if err := service.CrewDutyIndex.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailDutyIndex, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
