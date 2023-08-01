package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"strings"
)

var ProductMember = productMemberService{}

type productMemberService struct{}

// ProductMemberExcel Excel导出数据格式
type ProductMemberExcel struct {
	ProId      uint   // 项目ID
	EmpId      uint   // 项目成员ID
	WorkNumber string // 员工工号
	Attribute  uint   // 属性（1：全职，2：兼职）
	PrId       uint   // 项目角色ID
	PrName     string // 项目角色名称
	JbId       uint   // 职级ID
	JbName     string // 职级名称
}

func (s *productMemberService) Import(ctx context.Context, in *model.ProductMemberApiImportReq) (*entity.Product, error) {
	res := &entity.Product{}
	// 1: 验证项目信息
	checkProduct, _, err := s.checkInputData(ctx, &model.ProductMemberApiChangeReq{
		ProId: in.ProId,
	})
	if err != nil || !checkProduct {
		return res, err
	}
	// 2: 读取文件内容
	saveDataFormat, err := s.makeProductMemberExcelData(in.TableHeader, in.TableData, in.ProId)

	if err != nil {
		return res, err
	}
	// 3: 文件内容保存
	err = s.SaveProductMemberFromExcel(ctx, saveDataFormat)

	return res, err
}

// GetList 团队成员列表
func (s *productMemberService) GetList(ctx context.Context, in *model.ProductMemberGetListReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductMemberGetListRes, 0)
	res, dataEntity, err := dao.ProductMember.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range dataEntity {
			info := model.ProductMemberGetListRes{}
			info.ProductMemberInfo = v

			// 姓名
			employ, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{Id: gconv.Int32(v.EmpId)})
			if err != nil {
				return res, err
			}
			info.EmployeeInfo = employ.GetEmployee()
			// 职级
			jobLevel, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{Id: gconv.Int32(employ.GetEmployee().GetJobLevel())})
			if err != nil {
				return res, err
			}
			info.JobLevelInfo = jobLevel.GetJobLevel()

			resData = append(resData, info)
		}
	}

	res.Data = resData
	return res, nil
}

// GetAll 团队成员清单 - 不分页
func (s *productMemberService) GetAll(ctx context.Context, in *model.ProductMemberWhere) (res []*model.ProductMember, err error) {
	res, err = dao.ProductMember.GetAll(ctx, in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *productMemberService) GetOne(ctx context.Context, in *model.ProductMemberApiGetOneReq) (res *model.ProductMember, err error) {
	res, err = dao.ProductMember.GetOne(ctx, in.ProductMember)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *productMemberService) GetMemberInfo(ctx context.Context, in *model.ProductMember) (info model.ProductMemberInfo, err error) {
	info = model.ProductMemberInfo{}
	// 查询姓名
	employ, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{
		WorkNumber: in.WorkNumber,
	})
	if err != nil {
		return info, err
	}
	info.Employee = model.Employee{
		UserName: employ.GetEmployee().UserName,
		DepartId: employ.GetEmployee().DepartId,
		JobLevel: gconv.Uint(employ.GetEmployee().JobLevel),
	}

	return info, err
}

func (s *productMemberService) Create(ctx context.Context, in *model.ProductMemberApiChangeReq) (*model.ProductMember, error) {
	res := &model.ProductMember{}
	checkInput, in, err := s.checkInputData(ctx, in)
	if err != nil || !checkInput {
		return res, err
	}

	data := &model.ProductMember{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	res, err = dao.ProductMember.Create(ctx, data)
	return res, err
}

func (s *productMemberService) Modify(ctx context.Context, in *model.ProductMemberApiChangeReq) (*model.ProductMember, error) {
	res := &model.ProductMember{}
	if g.IsEmpty(in.Id) {
		return res, errors.New("缺少编辑对象")
	}

	checkInput, in, err := s.checkInputData(ctx, in)
	if err != nil || !checkInput {
		return res, err
	}

	data := &model.ProductMember{}
	input, _ := json.Marshal(in)
	err = json.Unmarshal(input, &data)
	if err != nil {
		return res, err
	}

	res, err = dao.ProductMember.Modify(ctx, data)
	return res, err
}

func (s *productMemberService) checkInputData(ctx context.Context, in *model.ProductMemberApiChangeReq) (bool, *model.ProductMemberApiChangeReq, error) {
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.Product.Columns().Id): in.ProId,
	}

	// 1: 项目基础信息是否已录入
	productInfo, err := dao.Product.GetOneByCondition(ctx, condition)
	if err != nil && err != sql.ErrNoRows {
		return false, in, err
	}
	if g.IsNil(productInfo) || g.IsEmpty(productInfo.Id) {
		return false, in, errors.New("当前项目不存在，请确认输入信息是否正确")
	}

	if !g.IsEmpty(in.WorkNumber) {
		employeeInfo, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{WorkNumber: in.WorkNumber})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, in, err
		}
		if g.IsNil(employeeInfo.GetEmployee()) || g.IsEmpty(employeeInfo.GetEmployee().GetId()) {
			return false, in, errors.New(fmt.Sprintf("工号：%s 员工信息未录入，请先录入", in.WorkNumber))
		}
		in.EmpId = gconv.Uint(employeeInfo.Employee.Id)
		in.JbId = gconv.Uint(employeeInfo.Employee.JobLevel)

		// 查询员工职级
		jobLevel, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{Id: gconv.Int32(in.JbId)})
		if err != nil {
			return false, in, err
		}
		if g.IsNil(jobLevel.GetJobLevel()) || g.IsEmpty(jobLevel.GetJobLevel().GetId()) {
			return false, in, errors.New(fmt.Sprintf("职级：%d 信息错误，请联系管理员", in.JbId))
		}
		in.JbName = jobLevel.GetJobLevel().Name
		// 责任指数
		in.DutyIndex, err = s.GetDutyIndexByJobLevel(ctx, gconv.Uint32(jobLevel.GetJobLevel().GetId()), jobLevel.GetJobLevel().GetName())
		if err != nil {
			return false, in, err
		}
	}

	// 员工投入属性ID值
	if !g.IsEmpty(in.AttributeName) {
		in.Attribute = util.GetEmployAttribute(in.AttributeName)
	}

	// 角色ID
	if !g.IsEmpty(in.PrName) {
		roles, err := boot.RolesServer.GetOne(ctx, &productV1.GetOneRolesReq{
			Roles: &productV1.RolesInfo{
				Name: in.PrName,
				Pid:  0,
			},
		})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, in, err
		}
		if g.IsNil(roles.GetRoles()) || g.IsEmpty(roles.GetRoles().GetId()) {
			return false, in, errors.New(fmt.Sprintf("项目角色：%s 信息未录入，请先录入", in.PrName))
		}
		in.PrId = gconv.Uint(roles.GetRoles().GetId())
		in.IsSpecial = gconv.Uint(roles.GetRoles().GetIsSpecial())
		// 管理指数
		in.ManageIndex, err = s.GetManageIndexByJobLevel(ctx, roles.GetRoles().GetId(), roles.GetRoles().GetPid())
		if err != nil {
			return false, in, err
		}
	}

	return true, in, err
}

func (s *productMemberService) SaveProductMemberFromExcel(ctx context.Context, excelData []*model.ProductMemberApiChangeReq) (err error) {
	if len(excelData) == 0 {
		return errors.New("文件内容为空，请先完善信息")
	}

	return dao.ProductMember.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询项目优先级确认配置信息
		for _, v := range excelData {
			proMem, err := dao.ProductMember.GetOne(ctx, model.ProductMember{
				ProId:      v.ProId,
				WorkNumber: v.WorkNumber,
			})
			if err != nil && err != sql.ErrNoRows {
				return err
			}

			checkData, in, err := s.checkInputData(ctx, v)
			if err != nil || !checkData {
				return err
			}
			model := &model.ProductMember{}
			gconv.Struct(in, model)
			model.Id = proMem.Id
			if g.IsEmpty(proMem.Id) {
				_, err = dao.ProductMember.Create(ctx, model)
			} else {
				_, err = dao.ProductMember.Modify(ctx, model)
			}

			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (s *productMemberService) GetDutyIndexByJobLevel(ctx context.Context, jbId uint32, jbName string) (scoreIndex uint32, err error) {
	dutyAll, err := boot.CrewDutyIndexServer.GetAll(ctx, &inspirit.GetAllCrewDutyIndexReq{
		CrewDutyIndex: &inspirit.CrewDutyIndexInfo{},
	})
	if err != nil {
		return 0, err
	}
	if len(dutyAll.GetData()) == 0 {
		return 0, errors.New("责任指数配置数据未录入，请先录入")
	}

	// 倒叙获取列表值
	jobLevel, err := boot.JobLevelServer.GetAll(ctx, &v1.GetAllJobLevelReq{
		JobLevel: &v1.JobLevelInfo{},
		Sort:     v1.OrderEnum_desc,
	})
	if err != nil {
		return 0, err
	}
	if len(jobLevel.GetData()) == 0 {
		return 0, errors.New("职级配置数据未录入，请先录入")
	}
	jobMap := make(map[uint32]uint32, 0)
	for _, v := range jobLevel.GetData() {
		jobMap[gconv.Uint32(v.GetName())] = gconv.Uint32(v.GetId())
	}

	for _, v := range dutyAll.GetData() {
		switch v.GetArith() {
		case inspirit.ArithEnum_eq:
			if v.JobLevelId == jbId {
				return v.GetScoreIndex(), nil
			}
		case inspirit.ArithEnum_gt:
			for kj, _ := range jobMap {
				if gconv.Uint32(jbName) < kj {
					return v.GetScoreIndex(), nil
				}
			}
		case inspirit.ArithEnum_lt:
			for kj, _ := range jobMap {
				if gconv.Uint32(jbName) > kj {
					return v.GetScoreIndex(), nil
				}
			}
		case inspirit.ArithEnum_egt:
			for kj, _ := range jobMap {
				if gconv.Uint32(jbName) <= kj {
					return v.GetScoreIndex(), nil
				}
			}
		case inspirit.ArithEnum_elt:
			for kj, _ := range jobMap {
				if gconv.Uint32(jbName) >= kj {
					return v.GetScoreIndex(), nil
				}
			}
		}
	}
	return 0, nil
}

// GetManageIndexByJobLevel 管理指数
func (s *productMemberService) GetManageIndexByJobLevel(ctx context.Context, id, pid int32) (scoreIndex uint, err error) {
	if pid > 0 {
		roles, err := boot.RolesServer.GetOne(ctx, &productV1.GetOneRolesReq{
			Roles: &productV1.RolesInfo{
				Id: pid,
			},
		})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return 0, err
		}
		manageInfo, err := boot.CrewManageIndexServer.GetOne(ctx, &inspirit.GetOneCrewManageIndexReq{
			CrewManageIndex: &inspirit.CrewManageIndexInfo{ProductRoleId: gconv.Uint32(roles.GetRoles().GetId())},
		})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return 0, err
		}
		return gconv.Uint(manageInfo.GetCrewManageIndex().GetScoreIndex()), nil
	}
	return 0, nil
}

func (s *productMemberService) makeProductMemberExcelData(tableHeader []string, tableData []map[string]interface{}, proId uint) ([]*model.ProductMemberApiChangeReq, error) {
	data := make([]*model.ProductMemberApiChangeReq, 0)
	if len(tableData) == 0 {
		return data, errors.New("表格数据为空，请先完善数据")
	}

	for _, v := range tableData {
		info := &model.ProductMemberApiChangeReq{ProId: proId}
		for vk, vv := range v {
			switch vk {
			case "项目角色":
				info.PrName = gconv.String(vv)
			case "工号":
				info.WorkNumber = gconv.String(vv)
			case "分类":
				info.Type = gconv.String(vv)
			case "投入占比":
				vvs := gconv.String(vv)
				if strings.Contains(vvs, "%") {
					putInto := strings.Split(vvs, "%")
					info.PutInto = util.Decimal(gconv.Float64(putInto[0]) / 100)
				} else {
					info.PutInto = util.Decimal(gconv.Float64(vvs))
				}
			case "责任和职务":
				info.SpecificDuty = gconv.String(vv)
			case "工作地":
				info.WorkAddress = gconv.String(vv)
			case "主导方":
				info.IsGuide = gconv.Uint(gconv.Bool(vv))
			case "支持方":
				info.IsSupport = gconv.Uint(gconv.Bool(vv))
			case "备注":
				info.Remark = gconv.String(vv)
			}
		}
		data = append(data, info)
	}
	return data, nil
}