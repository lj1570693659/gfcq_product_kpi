package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"strings"
)

// Employee 员工信息管理服务
var Employee = employeeService{}

type employeeService struct{}

// IsSyncEmployee 判断用户是否已经登录
func (s *employeeService) IsSyncEmployee(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User.UserInfo.EmployeeId > 0 {
		return true
	}
	return false
}

// GetList 获取员工信息列表
func (s *employeeService) GetList(ctx context.Context, input *model.EmployeeApiGetListReq) (*model.GetListEmployeeRes, error) {
	apiRes := &model.GetListEmployeeRes{}
	res := &v1.GetListEmployeeRes{}
	data := make([]model.EmployeeApiGetOneRes, 0)
	// 员工主体信息
	where := &v1.EmployeeInfo{
		Id:         gconv.Int32(input.Employee.Id),
		UserName:   input.Employee.UserName,
		WorkNumber: input.Employee.WorkNumber,
		Phone:      input.Employee.Phone,
		Email:      input.Employee.Email,
		DepartId:   input.Employee.DepartId,
		JobId:      input.Employee.JobId,
		JobLevel:   gconv.Int32(input.Employee.JobLevel),
		Status:     v1.StatusEnum(input.Employee.Status),
	}

	res, err := boot.EmployeeServer.GetList(ctx, &v1.GetListEmployeeReq{
		Employee: where,
		Page:     input.Page,
		Size:     input.Size,
	})
	if err != nil {
		return apiRes, err
	}

	if res.GetTotalSize() > 0 {
		for _, v := range res.GetData() {
			info := model.EmployeeApiGetOneRes{}
			// 性别
			info.SexName = util.GetEmploySex(v.GetSex())
			// 部门
			departmentList := make([]entity.Department, 0)
			departmentIds := gconv.Int32s(strings.Split(v.DepartId, ","))
			departmentName := make([]string, 0)
			if len(departmentIds) > 0 {
				for _, departId := range departmentIds {
					departmentInfo, err := boot.DepertmentServer.GetOne(ctx, &v1.GetOneDepartmentReq{
						Id: departId,
					})
					if err != nil {
						return apiRes, err
					}

					departmentList = append(departmentList, entity.Department{
						Id:     gconv.Int(departmentInfo.Department.Id),
						Name:   departmentInfo.Department.Name,
						Pid:    gconv.Int(departmentInfo.Department.Pid),
						Remark: departmentInfo.Department.Remark,
					})
					departmentName = append(departmentName, departmentInfo.Department.Name)
				}
				info.DepartmentName = strings.Join(departmentName, ",")
			}
			info.DepartmentInfo = departmentList
			// 岗位
			jobList := make([]entity.Job, 0)
			jobName := make([]string, 0)
			jobIds := gconv.Int32s(strings.Split(v.JobId, ","))
			if len(jobIds) > 0 {
				for _, jobId := range jobIds {
					job, err := boot.JobServer.GetOne(ctx, &v1.GetOneJobReq{
						Id: jobId,
					})
					if err != nil {
						return apiRes, err
					}

					jobList = append(jobList, entity.Job{
						Id:       gconv.Int(job.Job.Id),
						Name:     job.Job.Name,
						DepartId: gconv.Int(job.Job.DepartId),
						Remark:   job.Job.Remark,
					})
					jobName = append(jobName, job.Job.Name)
				}
			}
			info.JobInfo = jobList
			info.JobName = strings.Join(jobName, ",")
			// 职级
			if v.JobLevel > 0 {
				jobLevel, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{Id: v.JobLevel})
				if err != nil {
					return apiRes, err
				}
				fmt.Println("jobLevel=============", jobLevel)
				info.LevelInfo = model.JobLevel{
					Name: jobLevel.GetJobLevel().GetName(),
				}
			}
			// 指导老师
			if v.InstructorId > 0 {
				instructorInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{Id: v.InstructorId})
				if err != nil {
					return apiRes, err
				}
				info.InstructorInfo = model.Employee{
					UserName:   instructorInfo.GetEmployee().GetUserName(),
					WorkNumber: instructorInfo.GetEmployee().GetWorkNumber(),
				}
			}
			// 在职状态
			info.StatusName = util.GetEmployStatus(v.GetStatus())
			gconv.Struct(v, &info.EmployeeInfo)
			data = append(data, info)
		}
	}

	apiRes.Page = res.Page
	apiRes.Size = res.Size
	apiRes.TotalSize = res.TotalSize
	apiRes.Data = data
	return apiRes, nil
}

// GetAll 获取员工信息列表
func (s *employeeService) GetAll(ctx context.Context, input *model.Employee) (*v1.GetAllEmployeeRes, error) {
	res := &v1.GetAllEmployeeRes{}
	// 员工主体信息
	where := &v1.EmployeeInfo{
		Id:         gconv.Int32(input.Id),
		UserName:   input.UserName,
		WorkNumber: input.WorkNumber,
		Phone:      input.Phone,
		Email:      input.Email,
		DepartId:   input.DepartId,
		JobId:      input.JobId,
		JobLevel:   gconv.Int32(input.JobLevel),
		Status:     v1.StatusEnum(input.Status),
	}

	res, err := boot.EmployeeServer.GetAll(ctx, &v1.GetAllEmployeeReq{
		Employee: where,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetOne 获取员工信息详情
func (s *employeeService) GetOne(ctx context.Context, input *model.EmployeeApiGetOneReq) (res *model.EmployeeApiGetOneRes, err error) {
	res = &model.EmployeeApiGetOneRes{}

	// 员工主体信息
	where := &v1.GetOneEmployeeReq{
		Id:         gconv.Int32(input.Id),
		UserName:   input.UserName,
		WorkNumber: input.WorkNumber,
		Phone:      input.Phone,
		Email:      input.Email,
		JobLevel:   []int32{int32(input.JobLevel)},
		Status:     v1.StatusEnum(input.Status),
	}

	employeeInfo, err := boot.EmployeeServer.GetOne(ctx, where)
	if err != nil {
		return res, err
	}

	if !g.IsEmpty(employeeInfo) {
		employeeInfoByte, _ := json.Marshal(employeeInfo.Employee)
		json.Unmarshal(employeeInfoByte, &res.EmployeeInfo)

		// 员工岗位信息
		jobList := make([]entity.Job, 0)
		jobIds := gconv.Int32s(strings.Split(employeeInfo.Employee.JobId, ","))
		if len(jobIds) > 0 {
			for _, jobId := range jobIds {
				job, err := boot.JobServer.GetOne(ctx, &v1.GetOneJobReq{
					Id: jobId,
				})
				if err != nil {
					return res, err
				}

				jobList = append(jobList, entity.Job{
					Id:       gconv.Int(job.Job.Id),
					Name:     job.Job.Name,
					DepartId: gconv.Int(job.Job.DepartId),
					Remark:   job.Job.Remark,
				})
			}
		}
		res.JobInfo = jobList

		// 员工所在部门信息
		departmentList := make([]entity.Department, 0)
		departmentIds := gconv.Int32s(strings.Split(employeeInfo.Employee.DepartId, ","))
		if len(departmentIds) > 0 {
			for _, departId := range departmentIds {
				departmentInfo, err := boot.DepertmentServer.GetOne(ctx, &v1.GetOneDepartmentReq{
					Id: departId,
				})
				if err != nil {
					return res, err
				}

				departmentList = append(departmentList, entity.Department{
					Id:     gconv.Int(departmentInfo.Department.Id),
					Name:   departmentInfo.Department.Name,
					Pid:    gconv.Int(departmentInfo.Department.Pid),
					Remark: departmentInfo.Department.Remark,
				})
			}
		}

		res.DepartmentInfo = departmentList
	}

	return res, err
}

// Create 创建员工信息
func (s *employeeService) Create(ctx context.Context, input *model.EmployeeApiCreateReq) error {
	employeeInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{
		//WorkNumber: Context.Get(ctx).User.UserInfo.WorkNumber,
		WorkNumber: input.WorkNumber,
	})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	if !g.IsNil(employeeInfo) && !g.IsNil(employeeInfo.Employee) {
		return errors.New("员工信息已同步，请勿重复添加")
	}

	_, err = boot.EmployeeServer.Create(ctx, &v1.CreateEmployeeReq{
		Remark:   input.Remark,
		UserName: input.UserName,
		//WorkNumber:   Context.Get(ctx).User.UserInfo.WorkNumber,
		WorkNumber:   input.WorkNumber,
		Sex:          v1.SexEnum(input.Sex),
		Phone:        input.Phone,
		Email:        input.Email,
		JobLevel:     gconv.Int32(input.JobLevel),
		JobId:        gconv.Int32s(input.JobId),
		InstructorId: gconv.Int32(input.InstructorId),
		Status:       v1.StatusEnum(input.Status),
	})

	return err
}

// Modify 更新员工信息
func (s *employeeService) Modify(ctx context.Context, input *model.EmployeeApiModifyReq) error {
	employeeInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{
		WorkNumber: Context.Get(ctx).User.UserInfo.WorkNumber,
	})
	if err != nil {
		return err
	}
	if g.IsNil(employeeInfo) {
		return errors.New("员工信息未同步，请先同步信息")
	}

	_, err = boot.EmployeeServer.Modify(ctx, &v1.ModifyEmployeeReq{
		Id:       gconv.Int32(input.ID),
		Remark:   input.Remark,
		UserName: input.UserName,
		//WorkNumber:   Context.Get(ctx).User.UserInfo.WorkNumber,
		WorkNumber:   input.WorkNumber,
		Sex:          v1.SexEnum(input.Sex),
		Phone:        input.Phone,
		Email:        input.Email,
		JobLevel:     gconv.Int32(input.JobLevel),
		JobId:        gconv.Int32s(input.JobId),
		InstructorId: gconv.Int32(input.InstructorId),
		Status:       v1.StatusEnum(input.Status),
	})

	return err
}