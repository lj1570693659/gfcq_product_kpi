package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/do"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
)

// User 中间件管理服务
var User = userService{}

type userService struct{}

// SignUp 用户注册
func (s *userService) SignUp(ctx context.Context, r *model.UserServiceSignUpReq) error {
	// 账号唯一性数据检查
	if !s.CheckWorkNumber(ctx, r.WorkNumber) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.WorkNumber))
	}
	r.Password = util.Encrypt(r.Password)
	r.CreateTime = gtime.Now().String()
	r.UpdateTime = gtime.Now().String()
	if _, err := dao.User.Ctx(ctx).Data(r).Insert(); err != nil {
		return err
	}
	return nil
}

// IsSignedIn 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil && v.User.UserInfo != nil {
		return true
	}
	return false
}

// SignIn 用户登录，成功返回用户信息，否则返回nil; WorkNumber应当会md5值字符串
func (s *userService) SignIn(ctx context.Context, WorkNumber, password string) (model.Employee, error) {
	var user *entity.User
	userInfo := model.Employee{}
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().WorkNumber, WorkNumber).Where(dao.User.Columns().Password, util.Encrypt(password)).Scan(&user)
	if err != nil {
		return userInfo, err
	}
	if user == nil {
		return userInfo, errors.New("账号或密码错误")
	}
	sessionUser := &model.User{}
	sessionUserByte, _ := json.Marshal(user)
	json.Unmarshal(sessionUserByte, &sessionUser)
	if err := Session.SetUser(ctx, sessionUser); err != nil {
		return userInfo, err
	}
	Context.SetUserInfo(ctx, &model.UserInfo{
		Id:         gconv.Uint(user.Id),
		WorkNumber: user.WorkNumber,
		UserName:   user.UserName,
	})

	employeeInfo, err := Employee.GetOne(ctx, &model.EmployeeApiGetOneReq{model.Employee{WorkNumber: user.WorkNumber}})
	if err != nil {
		return userInfo, err
	}
	g.Log("login").Info(ctx, employeeInfo)
	Context.SetUserEmployee(ctx, employeeInfo.EmployeeInfo)
	return employeeInfo.EmployeeInfo, nil
}

// SignOut 用户注销
func (s *userService) SignOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// CheckWorkNumber 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckWorkNumber(ctx context.Context, WorkNumber string) bool {
	if i, err := dao.User.Ctx(ctx).Where(dao.Employee.Columns().WorkNumber, WorkNumber).Count(); err != nil {
		return false
	} else {
		return i == 0
	}
}

// GetProfile 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}

func (s *userService) ChangePwd(ctx context.Context, userInfo *model.UserApiChangePwdReq) error {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().WorkNumber, userInfo.WorkNumber).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("账号或密码错误")
	}

	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, user.Id).OmitEmpty().Data(do.User{Password: util.Encrypt(userInfo.Password)}).Update()
	if err != nil {
		return err
	}
	g.Log("login").Info(ctx, user)

	err = s.SignOut(ctx)
	return err
}
