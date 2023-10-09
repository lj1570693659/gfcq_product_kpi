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
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	"github.com/lj1570693659/gfcq_product_kpi/library/response"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

var ProductStageKpi = productStageKpiService{}
var KpiRuleLists []*inspirit.CrewKpiRuleInfo

type productStageKpiService struct{}

// 初始化数据
func init() {
	ctx := context.Background()
	res, err := boot.CrewKpiRuleServer.GetAll(ctx, &inspirit.GetAllCrewKpiRuleReq{})
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(res.GetData()) {
		panic("绩效等级未配置，请先完善数据")
	}
	KpiRuleLists = res.GetData()
}

func (s *productStageKpiService) Create(ctx context.Context, in *model.ProductStageKpiApiChangeReq) (*model.ProductStageKpi, error) {
	res := &model.ProductStageKpi{}
	check, err := s.checkInputData(ctx, in)
	if !check && err != nil {
		return res, err
	}

	input, err := s.completeInputData(ctx, in)
	if err != nil {
		return res, err
	}

	input, err = dao.ProductStageKpi.Create(ctx, input)
	return input, err
}

func (s *productStageKpiService) Modify(ctx context.Context, in *model.ProductStageKpiApiChangeReq) error {
	if g.IsEmpty(in.ID) {
		return errors.New("缺少编辑对象")
	}
	if ok, err := s.checkInputData(ctx, in); err != nil && !ok {
		return err
	}

	input, err := s.completeInputData(ctx, in)
	if err != nil {
		return err
	}

	// 更新绩效主表
	_, err = dao.ProductStageKpi.Modify(ctx, input)
	if err != nil {
		return err
	}

	err = dao.ProductStageKpi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新团队成员绩效表
		err = ProductMemberKpi.SyncPmKpi(ctx, in.ProId, in.StageId)

		// 更新团队成员奖金分配
		ProductMemberPrize.Compute(ctx, &model.ProductMemberPrizeComputeReq{
			ProId:   input.ProId,
			StageId: input.StageId,
		})
		return err
	})
	return err
}

// GetList 项目清单
func (s *productStageKpiService) GetList(ctx context.Context, in *model.ProductStageKpiApiGetListReq) (res *response.GetListResponse, err error) {
	resData := make([]model.ProductStageKpiList, 0)
	res, productEntity, err := dao.ProductStageKpi.GetList(ctx, in)
	if err != nil {
		return res, err
	}

	productList, err := dao.Product.GetAll(ctx, model.ProductWhere{})
	if err != nil {
		return res, err
	}

	if res.TotalSize > 0 {
		for _, v := range productEntity {
			info := model.ProductStageKpiList{
				ProductStageKpi: v,
			}
			// 项目信息
			info.ProductInfo, err = s.getProductInfo(v.ProId, productList)
			if err != nil {
				return res, err
			}
			info.StageInfo, err = dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: v.StageId})
			if err != nil {
				return res, err
			}
			resData = append(resData, info)
		}
	}
	res.Data = resData
	return res, nil
}

// GetOne 项目绩效详情
func (s *productStageKpiService) GetOne(ctx context.Context, in *model.ProductStageKpi) (res model.ProductStageKpiInfo, err error) {
	res = model.ProductStageKpiInfo{
		PmInfo:          &model.Employee{},
		PmlInfo:         &model.Employee{},
		ProductStageKpi: &model.ProductStageKpi{},
		ProductInfo: &model.ProductInfo{
			Product: &model.Product{},
		},
	}
	// 项目阶段绩效信息
	res.ProductStageKpi, err = dao.ProductStageKpi.GetOne(ctx, in)
	if err != nil {
		return res, err
	}
	// 项目信息
	productInfo, err := dao.Product.GetOne(ctx, model.Product{Id: res.ProductStageKpi.ProId})
	if err != nil {
		return res, err
	}
	infoByte, _ := json.Marshal(productInfo)
	json.Unmarshal(infoByte, res.ProductInfo.Product)

	//项目经理信息
	pmInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(productInfo.PmId)})
	if err != nil {
		return res, err
	}
	pmInfoByte, _ := json.Marshal(pmInfo.GetEmployee())
	json.Unmarshal(pmInfoByte, res.PmInfo)

	//项目负责人信息
	if productInfo.PmlId > 0 {
		pmlInfo, err := boot.EmployeeServer.GetOne(ctx, &common.GetOneEmployeeReq{Id: gconv.Int32(productInfo.PmlId)})
		if err != nil {
			return res, err
		}
		pmlInfoByte, _ := json.Marshal(pmlInfo.GetEmployee())
		json.Unmarshal(pmlInfoByte, res.PmlInfo)
	}

	// 项目类型
	typeInfo, err := boot.TypeServer.GetOne(ctx, &v1.GetOneTypeReq{Type: &v1.TypeInfo{Id: gconv.Int32(productInfo.Tid)}})
	if err != nil {
		return res, err
	}
	gconv.Struct(typeInfo.GetType(), &res.ProductInfo.ProductType)

	// 研发模式
	modeInfo, err := boot.ModeServer.GetOne(ctx, &v1.GetOneModeReq{Mode: &v1.ModeInfo{Id: gconv.Int32(productInfo.ModeId)}})
	if err != nil {
		return res, err
	}
	gconv.Struct(modeInfo.GetMode(), &res.ProductInfo.ProductMode)

	// 当前阶段
	res.StageInfo, err = ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: res.ProductStageKpi.StageId})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *productStageKpiService) checkInputData(ctx context.Context, in *model.ProductStageKpiApiChangeReq) (bool, error) {
	// 检查重复录入
	condition := g.Map{
		fmt.Sprintf("%s = ?", dao.ProductStageKpi.Columns().ProId): in.ProId,
	}
	if in.ID > 0 {
		condition["id != ?"] = in.ID
	}
	if in.StageId > 0 {
		condition["stage_id = ?"] = in.StageId
	}
	getInfo, err := dao.ProductStageKpi.GetOneByCondition(ctx, condition)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return false, err
	}

	if !g.IsNil(getInfo) && !g.IsEmpty(getInfo.Id) {
		return false, errors.New("当前阶段绩效已录入，请确认输入信息是否正确")
	}
	return true, nil
}
func (s *productStageKpiService) completeInputData(ctx context.Context, in *model.ProductStageKpiApiChangeReq) (*model.ProductStageKpi, error) {
	res := &model.ProductStageKpi{}
	gconv.Struct(in, res)

	// 已有数据：阶段、阶段得分、PM分配比例、浮动贡献、绩效等级
	// 1: 阶段比例  项目总激励预算（product表中incentive_budget字段值）
	// 1.1: 查询项目主表，获取需要的数据
	productInfo, err := dao.Product.GetOne(ctx, model.Product{Id: in.ProId})
	if err != nil {
		return res, err
	}
	// 1.2: 根绝项目类型，查询当前阶段所占激励预算比例
	stageInfo, err := ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: in.StageId})
	if err != nil {
		return res, err
	}

	if g.IsEmpty(stageInfo.ModeStage) {
		return res, errors.New("项目阶段不存在，请确认数据是否正确")
	}
	res.StageRadio = util.Decimal(gconv.Float64(stageInfo.Single.QuotaRadio))
	// 1.3 阶段预算 = 阶段比例 * 项目激励预算
	res.StageBudget = res.StageRadio * productInfo.IncentiveBudget
	// 1.4 阶段应发比例
	getRadio, err := boot.StageRadioServer.GetQuotaRadioByScore(ctx, &inspirit.GetQuotaRadioByScoreReq{Score: gconv.Uint32(in.StageScore)})
	if err != nil {
		return res, err
	}
	res.ShouldSentRadio = util.Decimal(gconv.Float64(getRadio.GetQuotaRadio()))
	// 1.5 阶段额度 = 阶段预算 * 应发比例
	res.StageQuota = util.Decimal(res.StageBudget * res.ShouldSentRadio)

	// 1.6 团队额度 = 阶段额度 * 团队占比
	teamRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.TeamRadio)
	if err != nil {
		return res, err
	}
	if g.IsEmpty(teamRadio) {
		return res, errors.New("请先完善项目团队激励比例配置信息")
	}
	res.CrewQuota = util.Decimal(res.StageQuota * gconv.Float64(teamRadio))

	// 1.7 团建额度 = 阶段额度 * 团建占比
	teamBuildingRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.TeamBuildingRadio)
	if err != nil {
		return res, err
	}
	if g.IsEmpty(teamBuildingRadio) {
		return res, errors.New("请先完善项目团建激励比例配置信息")
	}
	res.TeamBuildQuota = util.Decimal(res.StageQuota * gconv.Float64(teamBuildingRadio))

	// 1.8 业务支持额度 = 阶段额度 * 业务支持占比
	businessRadio, err := dao.Config.GetKeyValueByKeyName(ctx, consts.BusinessSupportRadio)
	if err != nil {
		return res, err
	}
	if g.IsEmpty(businessRadio) {
		return res, errors.New("请先完善项目业务支持激励比例配置信息")
	}
	res.SupportQuota = util.Decimal(res.StageQuota * gconv.Float64(businessRadio))

	// PM发放基础 = 团队额度 * PM分配比例
	res.PmBase = res.CrewQuota * res.PmRadio
	// PM绩效等级比例

	// 根据PM绩效得分获取绩效等级
	fmt.Println("KpiRuleLists---------------", KpiRuleLists)
	res.PmKpiLevelId = util.GetKpiRuleByScore(KpiRuleLists, res.PmKpiLevelScore)
	fmt.Println("KpiRuleLists---------------", res.PmKpiLevelId)
	kpiLevel, err := dao.CrewKpiRule.GetOne(ctx, model.CrewKpiRule{Id: res.PmKpiLevelId})
	if err != nil {
		return res, err
	}
	res.PmKpiLevelRadio = kpiLevel.Redio
	// PM实际应发额度 = (绩效等级对应比例 + 浮动比例) * 发放基数
	res.PmIncentiveQuota = util.Decimal((res.PmKpiLevelRadio + res.PmFloatRadio) * res.PmBase)
	// PM绩效等级名称
	res.PmKpiLevelName = kpiLevel.LevelName
	//PM实际应发额度
	return res, nil
}

func (s *productStageKpiService) getProductInfo(proId uint, productList []model.Product) (info model.Product, err error) {
	if len(productList) == 0 {
		return info, errors.New("项目清单数据为空，请先完善项目信息")
	}
	for _, v := range productList {
		if proId == v.Id {
			return v, nil
		}
	}
	return
}
