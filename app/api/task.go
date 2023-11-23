package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// Task 员工信息API管理对象
var Task = new(taskApi)

type taskApi struct{}

// AutoRemindNotEnd SignUp @summary 自动提醒
// @tags    项目管理
// @produce json
// @param   entity  body model.TaskApiGetListReq true "项目清单"
// @router  /task/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *taskApi) AutoRemindNotEnd(r *ghttp.Request) {
	service.Task.AutoRemindNotEnd(r.Context())
	response.JsonExit(r, response.Success, "ok")
}

// AutoRemindNotStart SignUp @summary 项目筛选清单
// @tags    项目管理
// @produce json
// @param   entity  body model.TaskApiGetListReq true "项目清单"
// @router  /Task/all [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *taskApi) AutoRemindNotStart(r *ghttp.Request) {
	service.Task.AutoRemindNotStart(r.Context())

	response.JsonExit(r, response.Success, "ok")

}

// AutoUpgrade @summary 项目主信息详情
// @tags    项目管理
// @produce json
// @param   entity  body model.TaskApiGetOneReq true "项目详情"
// @router  /Task/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) AutoUpgrade(r *ghttp.Request) {
	service.Task.AutoUpgrade(r.Context())
	response.JsonExit(r, response.Success, "ok")

}

// AutoCheck @summary 项目详情
// @tags    项目管理
// @produce json
// @param   entity  body model.TaskApiGetOneReq true "项目详情"
// @router  /Task/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) AutoCheck(r *ghttp.Request) {
	go service.Timer.RunTask(r.Context())

	response.JsonExit(r, response.Success, "ok")

}

// GetList SignUp @summary 获取部门信息列表
// @tags    部门信息服务
// @produce json
// @param   entity  body model.DepartmentApiGetListReq true "注册请求"
// @router  /system/organize/department/lists [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) GetList(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}

	res, err := service.Task.GetProductList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

func (a *taskApi) ProductModify(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailTask, "编辑对象数据丢失")
	}

	res, err := service.Task.ProductModify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetNpiList @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /Task/modify [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) GetNpiList(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}

	res, err := service.TaskNpi.GetNpiList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

func (a *taskApi) NpiModify(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailTask, "编辑对象数据丢失")
	}

	res, err := service.TaskNpi.NpiModify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetQualityList @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /Task/delete [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) GetQualityList(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}

	res, err := service.TaskQuality.GetQualityList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
func (a *taskApi) QualityModify(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailTask, "编辑对象数据丢失")
	}

	res, err := service.TaskQuality.QualityModify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProduceList @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /Task/delete [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) GetProduceList(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}

	res, err := service.TaskProduce.GetProduceList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
func (a *taskApi) ProduceModify(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailTask, "编辑对象数据丢失")
	}

	res, err := service.TaskProduce.ProduceModify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetDevelopList @summary 更新项目基础信息
// @tags    项目管理
// @produce json
// @param   entity  body model.EmployeeApiModifyReq true "注册请求"
// @router  /Task/delete [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *taskApi) GetDevelopList(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}

	res, err := service.TaskDevelop.GetDevelopList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
func (a *taskApi) DevelopModify(r *ghttp.Request) {
	var input *model.ProductTaskWhere

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailTask, err.Error())
	}
	if g.IsEmpty(input.ID) {
		response.JsonExit(r, response.FormatFailTask, "编辑对象数据丢失")
	}

	res, err := service.TaskDevelop.DevelopModify(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.CreateFailEmployee, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}
