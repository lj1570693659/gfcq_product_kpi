package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/service"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

// ProductStage 项目成员信息API管理对象
var ProductStage = new(productStageApi)

type productStageApi struct{}

// GetList SignUp @summary 项目团队成员清单
// @tags    项目团队管理
// @produce json
// @param   entity  body model.ProductStageGetListReq true "项目团队成员清单"
// @router  /product/member/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productStageApi) GetList(r *ghttp.Request) {
	var input *model.ProductStageRule

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.ProductStageRule.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}
