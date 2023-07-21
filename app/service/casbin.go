package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	"strings"
)

const (
	// ROLE 系统角色权限验证
	ROLE = "role"
	// BUSINESS_ROLE 项目权限验证
	BUSINESS_ROLE = "businessRole"
	// ACT_WRITE 读写权限
	ACT_WRITE = "write"
	// ACT_READ 只读权限
	ACT_READ = "read"
)

// Casbin 权限管理服务
var (
	Casbin = casbinService{}
	ok     bool
)

type casbinService struct{}

func (s *casbinService) CheckAuth(ctx context.Context, user *model.ContextUser, r *ghttp.Request, checkDimension string) (ok bool, err error) {
	if g.IsEmpty(user.EmployeeInfo) {
		return false, errors.New("请先完善员工信息")
	}

	// 1：路由-读写权限管理
	hasWrite := ACT_READ
	for _, v := range user.DepartmentInfo {
		if g.IsEmpty(v.Pid) {
			hasWrite = ACT_WRITE
		}
	}

	routerPath := []string{"system/account", "system/organize", "config/product", "config/inspirit", "achieve/product", "product"}
	for _, v := range routerPath {
		_, _ = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, v, hasWrite)
	}

	urlPath := util.DeleteIntSlice(strings.Split(r.URL.Path, "/"))
	ok, err = boot.Enforcer.Enforce(user.EmployeeInfo.WorkNumber, fmt.Sprintf("%s/%s", urlPath[0], urlPath[1]), hasWrite)

	// 2: 项目权限管理
	if checkDimension == BUSINESS_ROLE {
		if hasWrite == ACT_READ && len(user.ProductLists) > 0 {
			for _, v := range user.ProductLists {
				// 关联项目&&项目角色 TODO
				_, _ = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, v.Id, hasWrite)
			}
		}
	}

	return ok, err
}
