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
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
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
	checkProduct, info, err := s.checkInputData(ctx, &model.ProductMemberApiChangeReq{
		ProId: in.ProId,
	})
	if err != nil || !checkProduct {
		return res, err
	}
	// 2: 读取文件内容
	saveDataFormat, err := s.makeProductMemberExcelData(in.TableData, in.ProId)
	if err != nil {
		return res, err
	}

	// 3: 查询项目经理信息
	pmInfo, err := s.GetPmInfo(ctx, in.ProId)
	if err != nil {
		return res, err
	}

	// 3: 文件内容保存
	err = s.SaveProductMemberFromExcel(ctx, saveDataFormat, pmInfo.GetEmployee())
	if err != nil {
		return res, err
	}

	// 4： 同步项目经理信息
	err = s.SyncPmInfo(ctx, info)

	return res, err
}

func (s *productMemberService) WebImport(ctx context.Context, in *model.ProductMemberApiWebImportReq) (*entity.Product, error) {
	res := &entity.Product{}
	// 1: 验证项目信息
	checkProduct, info, err := s.checkInputData(ctx, &model.ProductMemberApiChangeReq{
		ProId: in.ProId,
	})
	if err != nil || !checkProduct {
		return res, err
	}
	// 2: 读取文件内容
	saveDataFormat, err := s.makeProductMemberWebData(in.UseridList, in.ProId)
	if err != nil {
		return res, err
	}

	// 3: 查询项目经理信息
	pmInfo, err := s.GetPmInfo(ctx, in.ProId)
	if err != nil {
		return res, err
	}

	// 3: 文件内容保存
	err = s.SaveProductMemberFromExcel(ctx, saveDataFormat, pmInfo.GetEmployee())
	if err != nil {
		return res, err
	}

	// 4： 同步项目经理信息
	err = s.SyncPmInfo(ctx, info)

	return res, err
}

func (s *productMemberService) GetPmInfo(ctx context.Context, proId uint) (pmInfo *v1.GetOneEmployeeRes, err error) {

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return pmInfo, err
	}

	proInfo, err := dao.Product.GetOne(ctx, model.Product{Id: proId})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return pmInfo, err
	}
	pmInfo, err = boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{Id: gconv.Int32(proInfo.PmId)})
	return pmInfo, err
}

// SyncPmInfo 同步项目经理数据
func (s *productMemberService) SyncPmInfo(ctx context.Context, in *model.ProductMemberApiChangeReq) error {
	pmInfo, err := s.GetPmInfo(ctx, in.ProId)
	checkData, in, err := s.checkInputData(ctx, &model.ProductMemberApiChangeReq{
		ProId:         in.ProId,
		EmpId:         in.EmpId,
		IsSpecial:     consts.IsPm,
		WorkNumber:    pmInfo.GetEmployee().GetWorkNumber(),
		AttributeName: "兼职",
		PrName:        "项目经理",
	})
	if err != nil || !checkData {
		return err
	}

	modelMember := &model.ProductMember{}
	gconv.Struct(in, modelMember)

	proMem, err := dao.ProductMember.GetOne(ctx, model.ProductMember{
		ProId:     in.ProId,
		IsSpecial: consts.IsPm,
	})

	modelMember.Id = proMem.Id

	if g.IsEmpty(proMem.Id) {
		_, err = dao.ProductMember.Create(ctx, modelMember)
	} else {
		_, err = dao.ProductMember.Modify(ctx, modelMember)
	}

	return err
}

// GetList 团队成员列表
func (s *productMemberService) GetList(ctx context.Context, in *model.ProductMemberGetListReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductMemberGetListRes, 0)
	res, dataEntity, err := dao.ProductMember.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	departList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &v1.GetListWithoutDepartmentReq{})
	if err != nil {
		return res, err
	}

	// 项目清单
	productList, err := dao.Product.GetAll(ctx, model.ProductWhere{})
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range dataEntity {
			info := model.ProductMemberGetListRes{}
			info.ProductMemberInfo = v
			for _, p := range productList {
				if p.Id == v.ProId {
					info.ProductInfo = p
				}
			}

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

			// 直接上级
			info.LeaderInfo, err = Employee.GetLeader(ctx, departList.GetData(), employ.GetEmployee().GetDepartId())
			if err != nil {
				return res, err
			}

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

// Delete 删除项目组成员
func (s *productMemberService) Delete(ctx context.Context, in *model.ProductMemberApiDeleteReq) (*model.ProductMember, error) {
	res := &model.ProductMember{}

	res, err := dao.ProductMember.GetOne(ctx, model.ProductMember{Id: in.Id})
	if err != nil {
		return res, err
	}

	productMemberKpi, err := dao.ProductMemberKpi.GetOne(ctx, model.ProductMemberKpi{ProId: in.ProId, ProEmpId: in.Id})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return res, err
	}
	if !g.IsEmpty(productMemberKpi.Id) {
		return res, errors.New("请先移除项目组成员绩效信息")
	}

	_, err = dao.ProductMember.Delete(ctx, in.Id)

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
		if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
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
	if len(in.PrName) > 0 {
		roles, err := boot.RolesServer.GetOne(ctx, &productV1.GetOneRolesReq{
			Roles: &productV1.RolesInfo{
				Name: in.PrName,
			},
		})

		if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
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
	} else if in.PrId > 0 {
		roles, err := boot.RolesServer.GetOne(ctx, &productV1.GetOneRolesReq{
			Roles: &productV1.RolesInfo{
				Id: gconv.Int32(in.PrId),
			},
		})

		if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
			return false, in, err
		}
		if g.IsNil(roles.GetRoles()) || g.IsEmpty(roles.GetRoles().GetId()) {
			return false, in, errors.New(fmt.Sprintf("项目角色：%s 信息未录入，请先录入", in.PrName))
		}
		in.PrName = roles.GetRoles().GetName()
		in.IsSpecial = gconv.Uint(roles.GetRoles().GetIsSpecial())
		// 管理指数
		in.ManageIndex, err = s.GetManageIndexByJobLevel(ctx, roles.GetRoles().GetId(), roles.GetRoles().GetPid())
		if err != nil {
			return false, in, err
		}
	}

	return true, in, err
}

func (s *productMemberService) SaveProductMemberFromExcel(ctx context.Context, excelData []*model.ProductMemberApiChangeReq, pmInfo *v1.EmployeeInfo) (err error) {
	if len(excelData) == 0 {
		return errors.New("文件内容识别失败，请先完善信息")
	}
	return dao.ProductMember.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询项目优先级确认配置信息
		for _, v := range excelData {
			if v.WorkNumber != pmInfo.GetWorkNumber() {
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
		}
		return nil
	})
}
func (s *productMemberService) GetDutyIndexByJobLevel(ctx context.Context, jbId uint32, jbName string) (scoreIndex uint32, err error) {
	dutyAll, err := boot.CrewDutyIndexServer.GetOne(ctx, &inspirit.GetOneCrewDutyIndexReq{
		CrewDutyIndex: &inspirit.CrewDutyIndexInfo{
			JobLevelId: jbId,
		},
	})
	if err != nil {
		return 0, err
	}
	if g.IsNil(dutyAll) || g.IsEmpty(dutyAll) {
		return 0, errors.New("责任指数配置数据未录入，请先录入")
	}

	return dutyAll.GetCrewDutyIndex().ScoreIndex, nil
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

func (s *productMemberService) makeProductMemberExcelData(tableData []map[string]interface{}, proId uint) (data []*model.ProductMemberApiChangeReq, err error) {
	data = make([]*model.ProductMemberApiChangeReq, 0)
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
				info.IsGuide = consts.IsNotGuide
				vvStr := gconv.String(vv)
				if len(vvStr) > 0 && strings.Contains(vvStr, "是") {
					info.IsGuide = consts.IsGuide
				}
			case "支持方":
				info.IsSupport = consts.IsNotSupport
				vvStr := gconv.String(vv)
				if len(vvStr) > 0 && strings.Contains(vvStr, "是") {
					info.IsSupport = consts.IsSupport
				}
			case "备注":
				info.Remark = gconv.String(vv)
			}
		}
		if len(info.PrName) != 0 && len(info.WorkNumber) != 0 {
			data = append(data, info)
		}
	}
	return data, nil
}

func (s *productMemberService) makeProductMemberWebData(tableData map[string][]string, proId uint) (data []*model.ProductMemberApiChangeReq, err error) {
	data = make([]*model.ProductMemberApiChangeReq, 0)
	if len(tableData) == 0 {
		return data, errors.New("数据为空，请选择项目组成员数据")
	}
	fmt.Println("tableData--------------", tableData)
	for key, value := range tableData {
		for _, v := range value {
			info := &model.ProductMemberApiChangeReq{
				ProId:      proId,
				WorkNumber: v,
				PrId:       gconv.Uint(key),
			}
			if len(info.WorkNumber) > 0 && info.PrId > 0 {
				data = append(data, info)
			}
		}
	}
	fmt.Println("data--------------", data)
	return data, nil
}

func (s *productMemberService) Export(ctx context.Context, in *model.ProductMemberWhere) (string, error) {
	excelData := make([]map[string]interface{}, 0)

	resData := make([]model.ProductMemberGetListRes, 0)
	dataEntity, err := dao.ProductMember.GetAll(ctx, in)
	if err != nil {
		return "", err
	}
	fmt.Println("dataEntity-----------------", len(dataEntity))

	departList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &v1.GetListWithoutDepartmentReq{})
	if err != nil {
		return "", err
	}

	// 项目清单
	productList, err := dao.Product.GetAll(ctx, model.ProductWhere{})
	if err != nil {
		return "", err
	}

	if len(dataEntity) > 0 {
		for _, v := range dataEntity {
			info := model.ProductMemberGetListRes{}
			info.ProductMemberInfo = *v
			for _, p := range productList {
				if p.Id == v.ProId {
					info.ProductInfo = p
				}
			}

			// 姓名
			employ, err := boot.EmployeeServer.GetOne(ctx, &v1.GetOneEmployeeReq{Id: gconv.Int32(v.EmpId)})
			if err != nil {
				return "", err
			}
			info.EmployeeInfo = employ.GetEmployee()
			// 职级
			jobLevel, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{Id: gconv.Int32(employ.GetEmployee().GetJobLevel())})
			if err != nil {
				return "", err
			}
			info.JobLevelInfo = jobLevel.GetJobLevel()

			// 直接上级
			info.LeaderInfo, err = Employee.GetLeader(ctx, departList.GetData(), employ.GetEmployee().GetDepartId())
			if err != nil {
				return "", err
			}

			info.DepartmentInfo, err = Employee.GetDepartment(ctx, departList.GetData(), employ.GetEmployee().GetDepartId())
			resData = append(resData, info)
		}
	}
	fmt.Println("resData-----------------", len(resData))
	fileName := "项目组成员"
	if len(resData) > 0 {
		for k, v := range resData {
			excelData = append(excelData, map[string]interface{}{
				"A": k + 1,                                          // 序号
				"B": v.ProductMemberInfo.PrName,                     // 项目角色
				"C": v.ProductMemberInfo.WorkNumber,                 // 工号
				"D": v.EmployeeInfo.UserName,                        // 姓名
				"E": v.ProductMemberInfo.Type,                       // 分类
				"F": v.DepartmentInfo.Name,                          // 部门
				"G": v.ProductMemberInfo.PutInto,                    // 投入水平
				"H": v.JobLevelInfo.Name,                            // 职级
				"I": v.ProductMemberInfo.WorkAddress,                // 工作地
				"J": v.ProductMemberInfo.SpecificDuty,               // 职责和任务
				"K": v.ProductMemberInfo.Remark,                     // 备注
				"L": util.GetIsGuide(v.ProductMemberInfo.IsGuide),   // 主导方
				"M": util.GetIsGuide(v.ProductMemberInfo.IsSupport), // 支持方
			})
		}

	}

	titleList := []string{"序号", "项目角色", "工号", "姓名", "分类", "部门", "投入占比", "职级", "工作地", "责任和职务", "备注", "主导方", "支持方"}

	// 保存Excel文件
	return util.SetCellValue(ctx, excelData, fileName, titleList)
}
