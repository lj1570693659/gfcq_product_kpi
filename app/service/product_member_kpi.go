package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"time"
)

var ProductMemberKpi = productMemberKpiService{}

type productMemberKpiService struct{}

func (s *productMemberKpiService) GetList(ctx context.Context, in model.ProductMemberKpiApiGetListReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductMemberKpiInfo, 0)
	res, entity, err := dao.ProductMemberKpi.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	// 部门清单
	departmentList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &common.GetListWithoutDepartmentReq{})
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range entity {
			info := model.ProductMemberKpiInfo{
				ProductMemberKpi: v,
			}
			// 员工信息
			info.ProductMember, err = ProductMember.GetOne(ctx, &model.ProductMemberApiGetOneReq{model.ProductMember{Id: v.ProEmpId}})
			if err != nil {
				return res, err
			}

			empInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(info.ProductMember.EmpId)})
			if err != nil {
				return res, err
			}
			info.UserName = empInfo.GetEmployee().GetUserName()
			info.DepartmentName = Department.GetDepartmentName(empInfo.GetEmployee().GetDepartId(), departmentList.GetData())
			info.ProductMember.PutInto = util.Decimal(info.ProductMember.PutInto)
			resData = append(resData, info)

		}
	}
	res.Data = resData
	return res, nil
}

// Export 项目绩效详情
func (s *productMemberKpiService) Export(ctx context.Context, in *model.ProductMemberExport) (string, error) {
	excelData := make([]map[string]interface{}, 0)

	// 项目成员
	memberList, err := ProductMember.GetAll(ctx, &model.ProductMemberWhere{
		ProId: in.ProId,
	})
	if err != nil {
		return "", err
	}
	// 项目信息
	productInfo, err := Product.GetOne(ctx, &model.ProductApiGetOneReq{model.Product{Id: in.ProId}})
	if err != nil {
		return "", err
	}

	departmentList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &common.GetListWithoutDepartmentReq{})
	if err != nil {
		return "", err
	}

	if len(memberList) > 0 {
		for k, v := range memberList {
			memberInfo, err := ProductMember.GetMemberInfo(ctx, v)
			if err != nil {
				return "", err
			}

			// 职级信息
			jobLevel, err := boot.JobLevelServer.GetOne(ctx, &common.GetOneJobLevelReq{Id: gconv.Int32(memberInfo.Employee.JobLevel)})
			if err != nil {
				return "", err
			}

			excelData = append(excelData, map[string]interface{}{
				"A": k + 1,                                                                                // 序号
				"B": v.PrName,                                                                             // 项目角色
				"C": v.WorkNumber,                                                                         // 工号
				"D": memberInfo.Employee.UserName,                                                         // 姓名
				"E": v.Type,                                                                               // 分类
				"F": Department.GetDepartmentName(memberInfo.Employee.DepartId, departmentList.GetData()), // 部门
				"G": v.PutInto,                                                                            // 投入占比
				"H": v.DutyIndex,                                                                          // 责任指数
				"I": jobLevel.GetJobLevel().GetName(),                                                     // 职级
				"J": v.SpecificDuty,                                                                       // 责任和职务
				"K": v.WorkAddress,                                                                        // 工作地
				"L": util.GetIsGuide(v.IsGuide),                                                           // 主导方
				"M": util.GetIsGuide(v.IsSupport),                                                         // 支持方
				"N": v.Remark,                                                                             // 备注
				"O": v.ManageIndex,                                                                        // 管理指数
				"P": "",                                                                                   // 工时占比
				"Q": "",                                                                                   // 浮动贡献
				"R": "",                                                                                   // 绩效等级
				"S": "",                                                                                   // 关键事件分类
				"T": "",                                                                                   // 事件描述
				"U": "",                                                                                   // 处理结果
				"V": "",                                                                                   // 发生时间
			})
		}

	}

	// 保存Excel文件
	return s.setCellValue(ctx, excelData, productInfo.SubName)
}

// setCellValue 保存Excel文件
func (s *productMemberKpiService) setCellValue(ctx context.Context, data []map[string]interface{}, productName string) (string, error) {
	titleList := []string{"序号", "项目角色", "工号", "姓名", "分类", "部门", "投入占比", "责任指数", "职级", "责任和职务", "工作地",
		"主导方", "支持方", "备注", "管理指数", "工时占比", "浮动贡献", "绩效得分", "关键事件分类", "事件描述", "处理结果", "发生时间"}
	sheetName := "Sheet1"
	fileName := fmt.Sprintf("/excel/%s-%s.xlsx", productName, time.Now().Format("2006-01-02"))
	filepath := fmt.Sprintf("./public/%s", fileName)
	if err := util.ExportExcel(titleList, data, sheetName, filepath); err != nil {
		g.Log("excel").Error(ctx, err)
	}

	return fileName, nil
}

func (s *productMemberKpiService) Import(ctx context.Context, in *model.ProductMemberKpiImportReq) error {
	// 1: 验证项目信息
	checkInput, _, _, err := s.checkInputData(ctx, model.ProductMemberKpiChangeReq{
		ProId:      in.ProId,
		ProStageId: in.StageId,
	})
	if err != nil || !checkInput {
		return err
	}
	// 2: 读取文件内容
	saveDataFormat, err := s.makeProductMemberKpiExcelData(in.TableHeader, in.TableData, in.ProId, in.StageId)
	if err != nil {
		return err
	}

	// 3: 查询项目经理信息
	pmInfo, err := ProductMember.GetPmInfo(ctx, in.ProId)
	if err != nil {
		return err
	}

	err = s.saveProductMemberKpiFromExcel(ctx, saveDataFormat, pmInfo.GetEmployee())

	// 同步项目经理数据
	s.SyncPmKpi(ctx, in.ProId, in.StageId)

	return err
}

// SyncPmKpi 同步项目经理绩效数据
func (s *productMemberKpiService) SyncPmKpi(ctx context.Context, proId, stageId uint) error {
	proPmKpi, err := dao.ProductMemberKpi.GetOne(ctx, model.ProductMemberKpi{
		ProId:      proId,
		ProStageId: stageId,
		IsPm:       consts.IsPm,
	})

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	checkInput, pmInput, getStageKpiInfo, err := s.checkInputData(ctx, model.ProductMemberKpiChangeReq{
		ProId:      proId,
		ProStageId: stageId,
		IsPm:       consts.IsPm,
	})
	if err != nil || !checkInput {
		return err
	}
	err = s.saveProductMemberKpi(ctx, model.ProductMemberKpiChangeReq{
		ID:            proPmKpi.Id,
		ProId:         proId,
		ProStageId:    stageId,
		IsPm:          consts.IsPm,
		WorkNumber:    pmInput.WorkNumber,
		OvertimeRadio: 0,
		FloatRaio:     getStageKpiInfo.PmFloatRadio,
		KpiLevel:      getStageKpiInfo.PmKpiLevelName,
	})
	return err
}

func (s *productMemberKpiService) Create(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	return s.saveProductMemberKpi(ctx, in)
}

func (s *productMemberKpiService) Modify(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	return s.saveProductMemberKpi(ctx, in)
}

func (s *productMemberKpiService) checkInputData(ctx context.Context, in model.ProductMemberKpiChangeReq) (bool, model.ProductMemberKpiChangeReq, *model.ProductStageKpi, error) {
	// 检查重复录入
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductStageKpi.Columns().ProId): in.ProId,
	}
	if in.ProStageId > 0 {
		condition["stage_id = ?"] = in.ProStageId
	}
	getInfo, err := dao.ProductStageKpi.GetOneByCondition(ctx, condition)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, in, getInfo, err
	}

	if g.IsNil(getInfo) || g.IsEmpty(getInfo.Id) {
		return false, in, getInfo, errors.New("项目阶段绩效未录入，请确认输入信息是否正确")
	}

	memberInfo := &model.ProductMember{}

	if len(in.WorkNumber) > 0 || in.IsPm == consts.IsPm {
		where := model.ProductMember{ProId: in.ProId}
		if len(in.WorkNumber) > 0 {
			where.WorkNumber = in.WorkNumber
		}
		if in.IsPm == consts.IsPm {
			where.IsSpecial = consts.IsPm
		}
		memberInfo, err = dao.ProductMember.GetOne(ctx, where)
		if err != nil {
			return false, in, getInfo, err
		}
		in.WorkNumber = memberInfo.WorkNumber
		in.IsPm = memberInfo.IsSpecial
		in.ProEmpId = memberInfo.Id
		in.PrId = memberInfo.PrId
		in.PrName = memberInfo.PrName
		in.JbId = memberInfo.JbId
		in.JbName = memberInfo.JbName
	}

	if in.KpiLevelScore > 0 {
		in.KpiLevelId = util.GetKpiRuleByScore(KpiRuleLists, in.KpiLevelScore)

		kpiInfo, err := boot.CrewKpiRuleServer.GetOne(ctx, &v1.GetOneCrewKpiRuleReq{
			CrewKpiRule: &v1.CrewKpiRuleInfo{
				Id: gconv.Int32(in.KpiLevelId),
			},
		})
		if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
			return false, in, getInfo, err
		}
		if g.IsNil(kpiInfo) || g.IsNil(kpiInfo.CrewKpiRule) || g.IsEmpty(kpiInfo.CrewKpiRule) {
			return false, in, getInfo, errors.New("绩效等级信息错误，请核实")
		}
		in.KpiLevel = kpiInfo.GetCrewKpiRule().GetLevelName()
		in.KpiRadio = util.Decimal(gconv.Float64(kpiInfo.GetCrewKpiRule().GetRedio()))
	}

	return true, in, getInfo, nil
}

func (s *productMemberKpiService) saveProductMemberKpiFromExcel(ctx context.Context, excelData []model.ProductMemberKpiChangeReq, pmInfo *common.EmployeeInfo) error {
	if len(excelData) == 0 {
		return errors.New("文件内容为空，请先完善信息")
	}
	g.Log("memberKpi").Info(ctx, excelData)
	return dao.ProductMemberKpi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询项目优先级确认配置信息
		for _, v := range excelData {
			fmt.Println("excelData---------------------", v)
			fmt.Println("excelData.WorkNumber---------------------", v.WorkNumber)
			fmt.Println("excelData.IsPm---------------------", v.IsPm)
			fmt.Println("pmInfo---------------------", pmInfo)
			if v.WorkNumber != pmInfo.GetWorkNumber() {
				if err := s.saveProductMemberKpi(ctx, v); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
func (s *productMemberKpiService) saveProductMemberKpi(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	info := &model.ProductMemberKpi{}
	checkData, in, _, err := s.checkInputData(ctx, in)
	if err != nil || !checkData {
		return err
	}

	gconv.Struct(in, info)
	proMemKpi, err := dao.ProductMemberKpi.GetOne(ctx, model.ProductMemberKpi{
		Id:         in.ID,
		ProId:      in.ProId,
		ProEmpId:   info.ProEmpId,
		ProStageId: in.ProStageId,
	})

	fmt.Println("saveProductMemberKpi.info-----------------", info)
	fmt.Println("saveProductMemberKpi.in-----------------", in)
	fmt.Println("saveProductMemberKpi-----------------", proMemKpi)

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	info.Id = proMemKpi.Id
	result := &model.ProductMemberKpi{}
	if g.IsEmpty(proMemKpi.Id) {
		result, err = dao.ProductMemberKpi.Create(ctx, info)
	} else {
		result, err = dao.ProductMemberKpi.Modify(ctx, info)
	}

	// 更新项目成员基准指数
	if info.IsPm == consts.IsPm {
		ProductMemberPrize.PmBaseIndexChange(context.Background(), info)
	} else {
		ProductMemberPrize.MemberBaseIndexChange(context.Background(), info)
	}

	// 更新项目成员关键事件清单
	if len(in.ProductMemberKey.KeyName) > 0 {
		_, err = dao.ProductMemberKey.Create(ctx, &model.ProductMemberKey{
			StageKpiId: result.Id,
			ProId:      in.ProId,
			ProEmpId:   info.ProEmpId,
			ProStageId: in.ProStageId,
			Type:       util.GetFloatKeyType(in.ProductMemberKey.Type),
			Property:   util.GetFloatKeyProperty(in.FloatRaio),
			KeyName:    in.ProductMemberKey.KeyName,
			HappenTime: gtime.NewFromStr(in.ProductMemberKey.HappenTime),
			Result:     in.ProductMemberKey.Result,
		})
	}

	return err

}

func (s *productMemberKpiService) makeProductMemberKpiExcelData(tableHeader []string, tableData []map[string]interface{}, proId, stageId uint) (saveDataFormat []model.ProductMemberKpiChangeReq, err error) {
	saveDataFormat = make([]model.ProductMemberKpiChangeReq, 0)
	if len(tableData) == 0 {
		return saveDataFormat, errors.New("表格数据为空，请先完善数据")
	}

	for _, v := range tableData {
		info := model.ProductMemberKpiChangeReq{ProId: proId, ProStageId: stageId, ProductMemberKey: model.ProductMemberKeyChangeReq{}}
		for vk, vv := range v {
			switch vk {
			case "工时占比":
				info.OvertimeRadio = gconv.Float64(vv)
			case "工号":
				info.WorkNumber = gconv.String(vv)
			case "浮动贡献":
				info.FloatRaio = gconv.Float64(vv)
			case "绩效得分":
				info.KpiLevelScore = gconv.Uint(vv)
			case "关键事件分类":
				info.ProductMemberKey.Type = gconv.String(vv)
			case "事件描述":
				info.ProductMemberKey.KeyName = gconv.String(vv)
			case "发生时间":
				info.ProductMemberKey.HappenTime = gconv.String(vv)
			case "处理结果":
				info.ProductMemberKey.Result = gconv.String(vv)
			}
		}
		saveDataFormat = append(saveDataFormat, info)
	}
	return saveDataFormat, err
}
