package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	wechat "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"strings"
	"time"
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
		UserName:   input.UserName,
		WorkNumber: input.WorkNumber,
		Phone:      input.Phone,
		Email:      input.Email,
		DepartId:   input.DepartId,
		JobId:      input.JobId,
		JobLevel:   gconv.Int32(input.JobLevel),
		Status:     v1.StatusEnum(input.Status),
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
			departmentIds := util.DeleteInt32Slice(gconv.Int32s(strings.Split(v.DepartId, ",")))
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
		res.JobIds = gconv.Ints(strings.Split(res.EmployeeInfo.JobId, ","))

		// 员工岗位信息
		jobList := make([]entity.Job, 0)
		jobName := make([]string, 0)
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
				jobName = append(jobName, job.Job.Name)
			}
		}
		res.JobInfo = jobList
		res.JobName = strings.Join(jobName, ",")

		// 员工所在部门信息
		departmentList := make([]entity.Department, 0)
		departmentName := make([]string, 0)
		res.DepartmentIds = util.DeleteInt32Slice(gconv.Int32s(strings.Split(employeeInfo.Employee.DepartId, ",")))
		if len(res.DepartmentIds) > 0 {
			for _, departId := range res.DepartmentIds {
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
					Level:  gconv.Uint(departmentInfo.Department.Level),
					Remark: departmentInfo.Department.Remark,
				})
				departmentName = append(departmentName, departmentInfo.Department.Name)
			}
		}

		res.DepartmentInfo = departmentList
		res.DepartmentName = strings.Join(departmentName, ",")

		// 职级信息
		if employeeInfo.GetEmployee().GetJobLevel() > 0 {
			jobLevel, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{Id: employeeInfo.GetEmployee().GetJobLevel()})
			if err != nil {
				return res, err
			}
			gconv.Struct(jobLevel.GetJobLevel(), &res.LevelInfo)
		}

		// 参与项目信息
		proMemberList, err := dao.ProductMember.GetAll(ctx, &model.ProductMemberWhere{EmpId: []uint{gconv.Uint(employeeInfo.GetEmployee().Id)}})
		if err != nil {
			return res, err
		}
		res.ProductMemberList = proMemberList
	}

	return res, err
}

// Create 创建员工信息
func (s *employeeService) Create(ctx context.Context, input *model.EmployeeApiCreateReq) error {
	if len(input.WorkNumber) == 0 {
		input.WorkNumber = Context.Get(ctx).User.UserInfo.WorkNumber
	}
	employeeInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{
		WorkNumber: input.WorkNumber,
	})

	if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
		return err
	}

	if !g.IsNil(employeeInfo) && !g.IsNil(employeeInfo.Employee) {
		return errors.New("员工信息已同步，请勿重复添加")
	}

	res, err := boot.EmployeeServer.Create(ctx, &v1.CreateEmployeeReq{
		Remark:       input.Remark,
		UserName:     input.UserName,
		WorkNumber:   input.WorkNumber,
		Sex:          v1.SexEnum(input.Sex),
		Phone:        input.Phone,
		Email:        input.Email,
		JobLevel:     gconv.Int32(input.JobLevel),
		JobId:        gconv.Int32s(input.JobId),
		InstructorId: gconv.Int32(input.InstructorId),
		Status:       v1.StatusEnum(input.Status),
	})

	// 默认同步登录系统账号信息
	err, isSignUp := User.IsSignUp(ctx, &model.UserServiceSignUpReq{
		WorkNumber: input.WorkNumber,
	})

	if !isSignUp && g.IsNil(err) {
		defaultUserPassword, err := dao.Config.GetKeyValueByKeyName(ctx, consts.DefaultUserPassword)
		if err != nil {
			g.Log("config").Error(ctx, err)
		}
		err = User.SignUp(ctx, &model.UserServiceSignUpReq{
			EmployeeId: res.GetEmployee().Id,
			WorkNumber: input.WorkNumber,
			UserName:   input.WorkNumber,
			Password:   defaultUserPassword,
		})
	}
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
		Id:           gconv.Int32(input.ID),
		Remark:       input.Remark,
		UserName:     input.UserName,
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

func (s *employeeService) GetEmployeeCount(ctx context.Context, departId int32) (*v1.GetCountEmployeeJobRes, error) {
	return boot.EmployeeJobServer.GetCount(ctx, &v1.GetCountEmployeeJobReq{
		EmployeeJob: &v1.EmployeeJobInfo{
			DepartId: departId,
		},
		GroupBy:           dao.EmployeeJob.Columns().EmployeeId,
		GetFiledNameCount: dao.EmployeeJob.Columns().EmployeeId,
	})
}

func (s *employeeService) GetLeader(ctx context.Context, departmentList []*v1.DepartmentInfo, departmentIds string) (map[string]string, error) {
	leader := map[string]string{}
	if len(departmentList) == 0 {
		return leader, nil
	}
	departIds := gconv.Int32s(util.DeleteIntSlice(strings.Split(departmentIds, ",")))
	for _, v := range departmentList {
		for _, empDv := range departIds {
			if empDv == v.Id {
				if v.Pid > 0 {
					// 下级部门
					leaderInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{DepartId: []int32{v.Pid}})
					if err != nil {
						return leader, nil
					}
					leader[v.GetName()] = fmt.Sprintf("%s-%s", leaderInfo.GetEmployee().UserName, leaderInfo.GetEmployee().WorkNumber)
				} else {
					// 上级部门
					leader[v.GetName()] = "-"
				}
			}
		}
	}
	return leader, nil
}

// GetCheckIn 获取员工信息列表
func (s *employeeService) GetCheckIn(ctx context.Context, input *model.GetCheckIn) (*wechat.GetUserCheckInDayDataRes, error) {
	res := &wechat.GetUserCheckInDayDataRes{}
	ts := time.Now().AddDate(0, 0, -1)
	tn := time.Now().AddDate(0, 0, 0)
	if len(input.StartTime) > 0 {
		ts, _ = time.Parse("2006-01-02", input.StartTime)
	}
	if len(input.StartTime) > 0 {
		tn, _ = time.Parse("2006-01-02", input.EndTime)
	}
	if input.ProId > 0 {
		productMemberList, _ := ProductMember.GetAll(ctx, &model.ProductMemberWhere{ProId: gconv.Uint(input.ProId)})
		if len(productMemberList) > 0 {
			for _, v := range productMemberList {
				input.UseridList = append(input.UseridList, v.WorkNumber)
			}
		}
	}
	fmt.Println("UseridList-------------------------", input.UseridList)
	fmt.Println("-------------------------", len(input.UseridList))
	res, err := boot.WechatCheckIn.GetUserCheckInDayData(ctx, &wechat.GetUserCheckInDayDataReq{
		DepartId:   input.DepartId,
		WorkNumber: input.UseridList,
		StartTime:  gconv.Int32(time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix()),
		EndTime:    gconv.Int32(time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location()).Unix()),
	})

	if err != nil {
		return res, err
	}
	return res, nil
}
