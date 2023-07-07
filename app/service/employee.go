package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

// 员工信息管理服务
var Employee = employeeService{}

type employeeService struct{}

// IsSyncEmployee 判断用户是否已经登录
func (s *employeeService) IsSyncEmployee(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User.EmployeeId > 0 {
		return true
	}
	return false
}

// IsSyncEmployee 判断用户是否已经登录
func (s *employeeService) Create(ctx context.Context, input *model.EmployeeApiCreateReq) error {
	employeeInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{
		WorkNumber: Context.Get(ctx).User.WorkNumber,
	})
	if err != nil {
		return err
	}
	if !g.IsNil(employeeInfo) {
		return errors.New("员工信息已同步，请勿重复添加")
	}

	createRes, err := boot.EmployeeServer.Create(ctx, &v1.CreateEmployeeReq{
		Remark:     input.Remark,
		UserName:   input.UserName,
		WorkNumber: Context.Get(ctx).User.WorkNumber,
		Sex:        v1.SexEnum(input.Sex),
		Phone:      input.Phone,
		Email:      input.Email,
		DepartId:   gconv.Int32(input.DepartId),
		JobLevel:   gconv.Int32(input.JobLevel),
		//JobId:        input.JobId,
		InstructorId: gconv.Int32(input.InstructorId),
		Status:       v1.StatusEnum(input.Status),
	})
	fmt.Println("Context.Get(ctx).User.WorkNumber------------------", Context.Get(ctx).User.WorkNumber)
	fmt.Println("employeeInfo------------------", employeeInfo)
	return nil

}
