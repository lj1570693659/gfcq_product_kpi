package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
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

	// LevelHigh 最高权限
	LevelHigh = 1
	// LevelMiddle 部门主管权限
	LevelMiddle = 2
	// LevelLow 部门员工权限
	LevelLow = 3
)

// Casbin 权限管理服务
var (
	Casbin = casbinService{}
	ok     bool
)

type casbinService struct{}

func (s *casbinService) CheckAuth(ctx context.Context, user *model.ContextUser, r *ghttp.Request, checkDimension string) (ok bool, err error) {
	if g.IsNil(user) || g.IsEmpty(user.EmployeeInfo) {
		return false, errors.New("请先完善员工信息")
	}

	// 1：路由-读写权限管理
	authLevel := LevelLow
	for _, v := range user.DepartmentInfo {
		if v.Level < gconv.Uint(authLevel) {
			authLevel = gconv.Int(v.Level)
		}
	}

	// 2: 项目权限管理
	isPm := consts.IsNotPm
	//if checkDimension == BUSINESS_ROLE {
	if authLevel == LevelLow {
		if len(user.ProductMemberList) > 0 {
			proIds := make([]uint, 0)
			for _, v := range user.ProductMemberList {
				proIds = append(proIds, v.ProId)
				if v.IsSpecial == consts.IsPm {
					isPm = consts.IsPm
				}
			}
			Context.SetUserProductIds(ctx, proIds)
			Context.SetUserProductRole(ctx, isPm)
		}
	}

	//}

	g.Log("auth").Info(ctx, fmt.Sprintf("用户：%s，权限级别为：%d, 所在部门信息：%v", user.EmployeeInfo.WorkNumber, authLevel, user.DepartmentInfo))
	Context.SetUserRoleLevel(ctx, authLevel)
	routerPath := []string{
		"system/account",
		"system/organize",
		"config/product",
		"config/inspirit",
		"achieve/product",
		"product/lists",
		"product/info",
		"product/all",
		"product/detail",
		"product/create",
		"product/delete",
		"product/modify",
		"product/member",
		"product/stage",
		"statistics/summation",
		"statistics/product",
		"statistics/level",
	}
	for _, v := range routerPath {
		if authLevel < LevelLow {
			ok, err = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, gconv.String(v), ACT_WRITE)
			ok, err = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, gconv.String(v), ACT_READ)
		} else {
			ok, err = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, gconv.String(v), ACT_READ)
			if Context.Get(ctx).User.ProductRole == consts.IsPm {
				ok, err = boot.Enforcer.AddPolicy(user.EmployeeInfo.WorkNumber, gconv.String(v), ACT_WRITE)
			}
		}
		if err != nil {
			g.Log("auth").Error(ctx, errors.New(fmt.Sprintf("员工：%s, 注册权限失败,失败原因：%v", user.EmployeeInfo.WorkNumber, err)))
			return ok, err
		}
	}

	urlPath := util.DeleteIntSlice(strings.Split(r.URL.Path, "/"))
	if r.Method == "GET" {
		ok, err = boot.Enforcer.Enforce(user.EmployeeInfo.WorkNumber, fmt.Sprintf("%s/%s", urlPath[0], urlPath[1]), ACT_READ)
	} else {
		ok, err = boot.Enforcer.Enforce(user.EmployeeInfo.WorkNumber, fmt.Sprintf("%s/%s", urlPath[0], urlPath[1]), ACT_WRITE)
	}
	if !ok {
		g.Log("auth").Info(ctx, errors.New(fmt.Sprintf("员工：%s, 权限校验失败, 访问路径：%v, 失败原因：%v", user.EmployeeInfo.WorkNumber, fmt.Sprintf("%s/%s", urlPath[0], urlPath[1]), err)))
		return ok, err
	}

	return ok, err
}

func (s *casbinService) CheckProductAuth(ctx context.Context, id uint) bool {
	if Context.Get(ctx).User.ProductRole == consts.IsPm {
		if util.CheckIn(Context.Get(ctx).User.ProductIds, id) {
			return true
		}
	} else {
		if Context.Get(ctx).User.RoleLevel < LevelLow {
			return true
		}
	}
	return false
}
