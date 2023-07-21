package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	"mime/multipart"
	"time"
)

var ProductMemberKpi = productMemberKpiService{}

type productMemberKpiService struct{}

func (s *productMemberKpiService) GetList(ctx context.Context, in model.ProductMemberKpiApiGetListReq) (res *response.GetListResponse, err error) {
	res, err = dao.ProductMemberKpi.GetList(ctx, in.ProductMemberKpi, in.Page, in.Size)
	if err != nil {
		return res, err
	}
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

	if len(memberList) > 0 {
		for k, v := range memberList {
			memberInfo, err := ProductMember.GetMemberInfo(ctx, v)
			if err != nil {
				return "", err
			}

			excelData = append(excelData, map[string]interface{}{
				"A": k + 1,                        // 序号
				"B": v.WorkNumber,                 // 工号
				"C": memberInfo.Employee.UserName, // 姓名
				"D": v.JbName,                     // 职级
				"E": v.DutyIndex,                  // 责任指数
				"F": v.PrName,                     // 项目角色
				"G": v.ManageIndex,                // 管理指数
				"H": "",                           // 绩效等级
				"I": "",                           // 工时占比
				"J": "",                           // 浮动贡献
				"K": "",                           // 关键事件分类
				"L": "",                           // 事件性质
				"M": "",                           // 事件描述
				"N": "",                           // 处理结果
			})
		}

	}

	// 保存Excel文件
	return s.setCellValue(ctx, excelData, productInfo.Name)
}

// setCellValue 保存Excel文件
func (s *productMemberKpiService) setCellValue(ctx context.Context, data []map[string]interface{}, productName string) (string, error) {
	titleList := []string{"序号", "工号", "姓名", "职级", "责任指数", "项目角色", "管理指数", "绩效等级", "工时占比", "浮动贡献", "关键事件分类", "事件性质", "事件描述", "处理结果"}
	sheetName := "Sheet1"
	filepath := fmt.Sprintf("./public/excel/%s-%s.xlsx", productName, time.Now().Format("2006-01-02"))
	if err := util.ExportExcel(titleList, data, sheetName, filepath); err != nil {
		g.Log("excel").Error(ctx, err)
	}

	return filepath, nil
}

func (s *productMemberKpiService) Import(ctx context.Context, fileInfo multipart.File, in *model.ProductMemberKpiImportReq) error {
	// 1: 验证项目信息
	checkInput, _, err := s.checkInputData(ctx, model.ProductMemberKpiChangeReq{
		ProId:      in.ProId,
		ProStageId: in.StageId,
	})
	if err != nil || !checkInput {
		return err
	}
	// 2: 读取文件内容
	utilExcelDataFormat, err := util.ReadExcel(fileInfo)
	if err != nil {
		return err
	}
	// 3: 文件内容保存
	saveDataFormat := make([]model.ProductMemberKpiChangeReq, 0)
	for _, v := range utilExcelDataFormat {
		input := model.ProductMemberKpiChangeReq{
			ProId:         in.ProId,
			ProStageId:    in.StageId,
			WorkNumber:    gconv.String(v.B),
			OvertimeRadio: gconv.Float64(v.I),
			FloatRaio:     gconv.Float64(v.J),
			KpiLevel:      gconv.String(v.H),
			ProductMemberKey: model.ProductMemberKeyChangeReq{
				WorkNumber: gconv.String(v.B),
				KeyName:    gconv.String(v.M),
				Type:       gconv.String(v.K),
				Property:   gconv.String(v.L),
				Result:     gconv.String(v.N),
				HappenTime: gconv.String(v.O),
			},
		}

		saveDataFormat = append(saveDataFormat, input)
	}

	return s.saveProductMemberKpiFromExcel(ctx, saveDataFormat)
}

func (s *productMemberKpiService) Create(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	return s.saveProductMemberKpi(ctx, in)
}

func (s *productMemberKpiService) Modify(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	return s.saveProductMemberKpi(ctx, in)
}

func (s *productMemberKpiService) checkInputData(ctx context.Context, in model.ProductMemberKpiChangeReq) (bool, model.ProductMemberKpiChangeReq, error) {
	// 检查重复录入
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductStageKpi.Columns().ProId): in.ProId,
	}
	if in.ProStageId > 0 {
		condition["id = ?"] = in.ProStageId
	}
	getInfo, err := dao.ProductStageKpi.GetOneByCondition(ctx, condition)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, in, err
	}

	if g.IsNil(getInfo) || g.IsEmpty(getInfo.Id) {
		return false, in, errors.New("项目阶段绩效未录入，请确认输入信息是否正确")
	}

	if len(in.WorkNumber) > 0 {
		memberInfo, err := dao.ProductMember.GetOne(ctx, model.ProductMember{
			WorkNumber: in.WorkNumber,
			ProId:      in.ProId,
		})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, in, err
		}
		in.ProEmpId = memberInfo.Id
		in.PrId = memberInfo.PrId
		in.PrName = memberInfo.PrName
		in.JbId = memberInfo.JbId
		in.JbName = memberInfo.JbName
	}
	if len(in.KpiLevel) > 0 {
		kpiInfo, err := boot.CrewKpiRuleServer.GetOne(ctx, &v1.GetOneCrewKpiRuleReq{
			CrewKpiRule: &v1.CrewKpiRuleInfo{
				LevelName: in.KpiLevel,
			},
		})
		if err != nil && err.Error() != sql.ErrNoRows.Error() {
			return false, in, err
		}
		if g.IsNil(kpiInfo) || g.IsNil(kpiInfo.CrewKpiRule) || g.IsEmpty(kpiInfo.CrewKpiRule) {
			return false, in, errors.New("绩效等级信息错误，请核实")
		}
		in.KpiLevelId = kpiInfo.GetCrewKpiRule().GetId()
		in.KpiRadio = util.Decimal(gconv.Float64(kpiInfo.GetCrewKpiRule().GetRedio()))
	}

	return true, in, nil
}

func (s *productMemberKpiService) saveProductMemberKpiFromExcel(ctx context.Context, excelData []model.ProductMemberKpiChangeReq) error {
	if len(excelData) == 0 {
		return errors.New("文件内容为空，请先完善信息")
	}

	return dao.ProductMemberKpi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询项目优先级确认配置信息
		for _, v := range excelData {
			if err := s.saveProductMemberKpi(ctx, v); err != nil {
				return err
			}
		}
		return nil
	})
}
func (s *productMemberKpiService) saveProductMemberKpi(ctx context.Context, in model.ProductMemberKpiChangeReq) error {
	info := &model.ProductMemberKpi{}
	checkData, in, err := s.checkInputData(ctx, in)
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

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}

	info.Id = proMemKpi.Id
	if g.IsEmpty(proMemKpi.Id) {
		_, err = dao.ProductMemberKpi.Create(ctx, info)
	} else {
		_, err = dao.ProductMemberKpi.Modify(ctx, info)
	}

	// 更新项目成员基准指数
	go ProductMemberPrize.MemberBaseIndexChange(context.Background(), info)

	return err

}
