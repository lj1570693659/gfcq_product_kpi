package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// CrewManageIndex 管理指数
var CrewManageIndex = new(crewManageIndexApi)

type crewManageIndexApi struct{}

// GetAll SignUp @summary 获取激励预算信息列表
// @tags    管理指数
// @produce json
// @param   entity  body model.BudgetAccessApiGetListReq true "注册请求"
// @router  /config/inspirit/manage/all [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewManageIndexApi) GetAll(r *ghttp.Request) {
	var input *model.CrewManageIndex

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	res, err := service.CrewManageIndex.GetAll(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// Create SignUp @summary 完善激励预算信息
// @tags    管理指数
// @produce json
// @param   entity  body model.CrewManageIndexApiChangeReq true "创建激励预算"
// @router  /config/inspirit/manage/create [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewManageIndexApi) Create(r *ghttp.Request) {
	var input *model.CrewManageIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.CrewManageIndex.Create(r.Context(), input); err != nil {
		response.JsonExit(r, response.CreateFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Modify SignUp @summary 更新激励预算信息
// @tags    管理指数
// @produce json
// @param   entity  body model.CrewManageIndexApiChangeReq true "注册请求"
// @router  /config/inspirit/manage/modify [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewManageIndexApi) Modify(r *ghttp.Request) {
	var input *model.CrewManageIndexApiChangeReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailStageRadio, "编辑对象数据丢失")
	}

	if err := service.CrewManageIndex.Modify(r.Context(), input); err != nil {
		response.JsonExit(r, response.ModifyFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}

// Delete Modify SignUp @summary 删除激励预算信息
// @tags    管理指数
// @produce json
// @param   entity  body model.CrewManageIndexApiDeleteReq true "删除请求"
// @router  /config/inspirit/manage/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *crewManageIndexApi) Delete(r *ghttp.Request) {
	var input *model.CrewManageIndexApiDeleteReq

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailStageRadio, err.Error())
	}

	if err := service.CrewManageIndex.Delete(r.Context(), input); err != nil {
		response.JsonExit(r, response.DeleteFailStageRadio, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
