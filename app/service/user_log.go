package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	"strings"
)

// UserLog 用户操作记录服务
var UserLog = userLogService{}

type userLogService struct{}

//1: GET 2: POST 3:PUT 4:DELETE
var methodMap = map[string]uint{
	"GET":    consts.MethodGET,
	"POST":   consts.MethodPOST,
	"PUT":    consts.MethodPUT,
	"DELETE": consts.MethodDELETE,
}

func (s *userLogService) SaveLogData(ctx context.Context, cookie *ghttp.Cookie, requestData map[string]interface{}, method, requestURI string) error {
	data := entity.UserLog{RequestUri: requestURI}
	if v, ok := methodMap[method]; ok {
		if v == consts.MethodGET {
			return nil
		}
		data.MethodName = v
	}
	data.WorkNumber = cookie.Get(consts.TokenName)
	requestModuleLists := util.DeleteIntSlice(strings.Split(requestURI, "/"))
	if len(requestModuleLists) > 0 {
		data.RequestModule = requestModuleLists[0]
		data.RequestSecondModule = requestModuleLists[1]
	}

	data.ChangeTypeName = util.GetUserRequestTypeName(data.MethodName, requestModuleLists)
	data.RequestBody = gconv.String(requestData)
	data.CreateTime = gtime.Now()
	data.UpdateTime = gtime.Now()
	err := dao.UserLog.Create(ctx, data)
	return err
}

// GetList 日志管理清单
func (s *userLogService) GetList(ctx context.Context, in *model.UserLogApiReq) (res *response.GetListResponse, err error) {
	if len(Context.Get(ctx).User.UserInfo.WorkNumber) > 0 {
		in.WorkNumber = Context.Get(ctx).User.UserInfo.WorkNumber
	}

	res, _, err = dao.UserLog.GetList(ctx, in)
	return res, err
}
