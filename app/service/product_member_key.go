package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
)

var ProductMemberKey = productMemberKeyService{}

type productMemberKeyService struct{}

func (s *productMemberKeyService) GetList(ctx context.Context, in model.ProductMemberKeyListsReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductMemberKeyList, 0)
	res, entity, err := dao.ProductMemberKey.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range entity {
			info := model.ProductMemberKeyList{
				ProductMemberKey: v,
			}
			// 成员绩效信息
			info.ProductMemberKpi, err = dao.ProductMemberKpi.GetOne(ctx, model.ProductMemberKpi{Id: v.StageKpiId})
			if err != nil {
				return res, err
			}

			resData = append(resData, info)

		}
	}
	res.Data = resData
	return res, nil
}

func (s *productMemberKeyService) Create(ctx context.Context, in model.ProductMemberKeyApiChangeReq) error {
	info, err := dao.ProductMemberKey.GetOne(ctx, model.ProductMemberKey{
		StageKpiId: in.StageKpiId,
	})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	if info.Id > 0 {
		return errors.New("当前成员关键事件已存在，请勿重复创建")
	}

	data := &model.ProductMemberKey{
		StageKpiId: in.StageKpiId,
		ProId:      in.ProId,
		ProEmpId:   in.ProEmpId,
		ProStageId: in.ProStageId,
		WorkNumber: in.WorkNumber,
		Username:   in.Username,
		KeyName:    in.KeyName,
		HappenTime: in.HappenTime,
		Type:       in.Type,
		Property:   in.Property,
		Result:     in.Result,
		Remark:     in.Remark,
	}
	_, err = dao.ProductMemberKey.Create(ctx, data)

	return err
}

func (s *productMemberKeyService) Modify(ctx context.Context, in model.ProductMemberKeyApiChangeReq) error {
	info, err := dao.ProductMemberKey.GetOne(ctx, model.ProductMemberKey{
		StageKpiId: in.StageKpiId,
	})
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	if g.IsEmpty(info.Id) {
		return errors.New("当前成员关键事件不存在，请先创建")
	}
	data := &model.ProductMemberKey{
		Id:         info.Id,
		StageKpiId: info.StageKpiId,
		ProId:      info.ProId,
		ProEmpId:   info.ProEmpId,
		ProStageId: info.ProStageId,
		WorkNumber: info.WorkNumber,
		Username:   info.Username,
		KeyName:    in.KeyName,
		HappenTime: in.HappenTime,
		Type:       in.Type,
		Property:   in.Property,
		Result:     in.Result,
		Remark:     in.Remark,
	}

	_, err = dao.ProductMemberKey.Modify(ctx, data)

	return err
}

// Delete 删除项目优先级信息
func (s *productMemberKeyService) Delete(ctx context.Context, input model.ProductMemberKeyApiDeleteReq) error {
	_, err := dao.ProductMemberKey.Delete(ctx, input.ID)

	return err
}

// Export 项目绩效详情
//func (s *productMemberKeyService) Export(ctx context.Context, in *model.ProductMemberExport) (string, error) {
//	excelData := make([]map[string]interface{}, 0)
//
//	// 项目成员
//	memberList, err := ProductMember.GetAll(ctx, &model.ProductMemberWhere{
//		ProId: in.ProId,
//	})
//	if err != nil {
//		return "", err
//	}
//	// 项目信息
//	productInfo, err := Product.GetOne(ctx, &model.ProductApiGetOneReq{model.Product{Id: in.ProId}})
//	if err != nil {
//		return "", err
//	}
//
//	departmentList, err := boot.DepertmentServer.GetListWithoutPage(ctx, &common.GetListWithoutDepartmentReq{})
//	if err != nil {
//		return "", err
//	}
//
//	if len(memberList) > 0 {
//		for k, v := range memberList {
//			memberInfo, err := ProductMember.GetMemberInfo(ctx, v)
//			if err != nil {
//				return "", err
//			}
//
//			// 职级信息
//			jobLevel, err := boot.JobLevelServer.GetOne(ctx, &common.GetOneJobLevelReq{Id: gconv.Int32(memberInfo.Employee.JobLevel)})
//			if err != nil {
//				return "", err
//			}
//
//			excelData = append(excelData, map[string]interface{}{
//				"A": k + 1,                                                                                // 序号
//				"B": v.PrName,                                                                             // 项目角色
//				"C": v.WorkNumber,                                                                         // 工号
//				"D": memberInfo.Employee.UserName,                                                         // 姓名
//				"E": v.Type,                                                                               // 分类
//				"F": Department.GetDepartmentName(memberInfo.Employee.DepartId, departmentList.GetData()), // 部门
//				"G": v.PutInto,                                                                            // 投入占比
//				"H": v.DutyIndex,                                                                          // 责任指数
//				"I": jobLevel.GetJobLevel().GetName(),                                                     // 职级
//				"J": v.SpecificDuty,                                                                       // 责任和职务
//				"K": v.WorkAddress,                                                                        // 工作地
//				"L": util.GetIsGuide(v.IsGuide),                                                           // 主导方
//				"M": util.GetIsGuide(v.IsSupport),                                                         // 支持方
//				"N": v.Remark,                                                                             // 备注
//				"O": v.ManageIndex,                                                                        // 管理指数
//				"P": "",                                                                                   // 工时占比
//				"Q": "",                                                                                   // 浮动贡献
//				"R": "",                                                                                   // 绩效等级
//				"S": "",                                                                                   // 关键事件分类
//				"T": "",                                                                                   // 事件描述
//				"U": "",                                                                                   // 处理结果
//			})
//		}
//
//	}
//
//	// 保存Excel文件
//	return s.setCellValue(ctx, excelData, productInfo.Name)
//}
//
//// setCellValue 保存Excel文件
//func (s *productMemberKeyService) setCellValue(ctx context.Context, data []map[string]interface{}, productName string) (string, error) {
//	titleList := []string{"序号", "项目角色", "工号", "姓名", "分类", "部门", "投入占比", "责任指数", "职级", "责任和职务", "工作地",
//		"主导方", "支持方", "备注", "管理指数", "工时占比", "浮动贡献", "绩效等级", "关键事件分类", "事件描述", "处理结果"}
//	sheetName := "Sheet1"
//	fileName := fmt.Sprintf("/excel/%s-%s.xlsx", productName, time.Now().Format("2006-01-02"))
//	filepath := fmt.Sprintf("./public/%s", fileName)
//	if err := util.ExportExcel(titleList, data, sheetName, filepath); err != nil {
//		g.Log("excel").Error(ctx, err)
//	}
//
//	return fileName, nil
//}
//
//func (s *productMemberKeyService) Import(ctx context.Context, in *model.ProductMemberKeyImportReq) error {
//	// 1: 验证项目信息
//	checkInput, _, err := s.checkInputData(ctx, model.ProductMemberKeyChangeReq{
//		ProId:      in.ProId,
//		ProStageId: in.StageId,
//	})
//	if err != nil || !checkInput {
//		return err
//	}
//	// 2: 读取文件内容
//	saveDataFormat, err := s.makeProductMemberKeyExcelData(in.TableHeader, in.TableData, in.ProId, in.StageId)
//	if err != nil {
//		return err
//	}
//
//	return s.saveProductMemberKeyFromExcel(ctx, saveDataFormat)
//}
//
//func (s *productMemberKeyService) Create(ctx context.Context, in model.ProductMemberKeyChangeReq) error {
//	return s.saveProductMemberKey(ctx, in)
//}
//
//func (s *productMemberKeyService) Modify(ctx context.Context, in model.ProductMemberKeyChangeReq) error {
//	return s.saveProductMemberKey(ctx, in)
//}
//
//func (s *productMemberKeyService) checkInputData(ctx context.Context, in model.ProductMemberKeyChangeReq) (bool, model.ProductMemberKeyChangeReq, error) {
//	// 检查重复录入
//	condition := g.Map{
//		fmt.Sprintf("%s = ?", dao.ProductStageKpi.Columns().ProId): in.ProId,
//	}
//	if in.ProStageId > 0 {
//		condition["stage_id = ?"] = in.ProStageId
//	}
//	getInfo, err := dao.ProductStageKpi.GetOneByCondition(ctx, condition)
//	if err != nil && err.Error() != sql.ErrNoRows.Error() {
//		return false, in, err
//	}
//
//	if g.IsNil(getInfo) || g.IsEmpty(getInfo.Id) {
//		return false, in, errors.New("项目阶段绩效未录入，请确认输入信息是否正确")
//	}
//
//	if len(in.WorkNumber) > 0 {
//		memberInfo, err := dao.ProductMember.GetOne(ctx, model.ProductMember{
//			WorkNumber: in.WorkNumber,
//			ProId:      in.ProId,
//		})
//		if err != nil && err.Error() != sql.ErrNoRows.Error() {
//			return false, in, err
//		}
//		in.ProEmpId = memberInfo.Id
//		in.PrId = memberInfo.PrId
//		in.PrName = memberInfo.PrName
//		in.JbId = memberInfo.JbId
//		in.JbName = memberInfo.JbName
//	}
//	if len(in.KpiLevel) > 0 {
//		kpiInfo, err := boot.CrewKpiRuleServer.GetOne(ctx, &v1.GetOneCrewKpiRuleReq{
//			CrewKpiRule: &v1.CrewKpiRuleInfo{
//				LevelName: in.KpiLevel,
//			},
//		})
//		if err != nil && err.Error() != sql.ErrNoRows.Error() {
//			return false, in, err
//		}
//		if g.IsNil(kpiInfo) || g.IsNil(kpiInfo.CrewKpiRule) || g.IsEmpty(kpiInfo.CrewKpiRule) {
//			return false, in, errors.New("绩效等级信息错误，请核实")
//		}
//		in.KpiLevelId = kpiInfo.GetCrewKpiRule().GetId()
//		in.KpiRadio = util.Decimal(gconv.Float64(kpiInfo.GetCrewKpiRule().GetRedio()))
//	}
//
//	return true, in, nil
//}
//
//func (s *productMemberKeyService) saveProductMemberKeyFromExcel(ctx context.Context, excelData []model.ProductMemberKeyChangeReq) error {
//	if len(excelData) == 0 {
//		return errors.New("文件内容为空，请先完善信息")
//	}
//
//	return dao.ProductMemberKey.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
//		// 查询项目优先级确认配置信息
//		for _, v := range excelData {
//			if err := s.saveProductMemberKey(ctx, v); err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//}
//func (s *productMemberKeyService) saveProductMemberKey(ctx context.Context, in model.ProductMemberKeyChangeReq) error {
//	info := &model.ProductMemberKey{}
//	checkData, in, err := s.checkInputData(ctx, in)
//	if err != nil || !checkData {
//		return err
//	}
//	gconv.Struct(in, info)
//	proMemKpi, err := dao.ProductMemberKey.GetOne(ctx, model.ProductMemberKey{
//		Id:         in.ID,
//		ProId:      in.ProId,
//		ProEmpId:   info.ProEmpId,
//		ProStageId: in.ProStageId,
//	})
//
//	if err != nil && err.Error() != sql.ErrNoRows.Error() {
//		return err
//	}
//
//	info.Id = proMemKpi.Id
//	result := &model.ProductMemberKey{}
//	if g.IsEmpty(proMemKpi.Id) {
//		result, err = dao.ProductMemberKey.Create(ctx, info)
//	} else {
//		result, err = dao.ProductMemberKey.Modify(ctx, info)
//	}
//
//	// 更新项目成员基准指数
//	go ProductMemberPrize.MemberBaseIndexChange(context.Background(), info)
//
//	// 更新项目成员关键事件清单
//	if len(in.ProductMemberKey.KeyName) > 0 {
//		_, err = dao.ProductMemberKey.Create(ctx, &model.ProductMemberKey{
//			StageKpiId: result.Id,
//			ProId:      in.ProId,
//			ProEmpId:   info.ProEmpId,
//			ProStageId: in.ProStageId,
//			Type:       util.GetFloatKeyType(in.ProductMemberKey.Type),
//			Property:   util.GetFloatKeyProperty(in.FloatRaio),
//			KeyName:    in.ProductMemberKey.KeyName,
//			HappenTime: gtime.NewFromStr(in.ProductMemberKey.HappenTime),
//			Result:     in.ProductMemberKey.Result,
//		})
//	}
//
//	return err
//
//}
//
//func (s *productMemberKeyService) makeProductMemberKeyExcelData(tableHeader []string, tableData []map[string]interface{}, proId, stageId uint) (saveDataFormat []model.ProductMemberKeyChangeReq, err error) {
//	saveDataFormat = make([]model.ProductMemberKeyChangeReq, 0)
//	if len(tableData) == 0 {
//		return saveDataFormat, errors.New("表格数据为空，请先完善数据")
//	}
//
//	for _, v := range tableData {
//		info := model.ProductMemberKeyChangeReq{ProId: proId, ProStageId: stageId, ProductMemberKey: model.ProductMemberKeyChangeReq{}}
//		for vk, vv := range v {
//			switch vk {
//			case "工时占比":
//				info.OvertimeRadio = gconv.Float64(vv)
//			case "工号":
//				info.WorkNumber = gconv.String(vv)
//			case "浮动贡献":
//				info.FloatRaio = gconv.Float64(vv)
//			case "绩效等级":
//				info.KpiLevel = gconv.String(vv)
//			case "关键事件分类":
//				info.ProductMemberKey.Type = gconv.String(vv)
//			case "事件描述":
//				info.ProductMemberKey.KeyName = gconv.String(vv)
//			case "发生时间":
//				info.ProductMemberKey.HappenTime = gconv.String(vv)
//			case "处理结果":
//				info.ProductMemberKey.Result = gconv.String(vv)
//			}
//		}
//		saveDataFormat = append(saveDataFormat, info)
//	}
//	return saveDataFormat, err
//}
