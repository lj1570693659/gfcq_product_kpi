package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// StatisticsSummation 总量统计API
var StatisticsSummation = statisticsSummationService{}

type statisticsSummationService struct{}

func (a *statisticsSummationService) GetInspire(ctx context.Context) (inspireData model.Inspire, err error) {
	inspireData = model.Inspire{}
	// 项目数量
	if inspireData.ProductCount, err = dao.Product.Ctx(ctx).Count(dao.Product.Columns().Id); err != nil {
		return inspireData, err
	}
	// 总激励预算
	if inspireData.IncentiveBudget, err = dao.Product.Ctx(ctx).Sum(dao.Product.Columns().IncentiveBudget); err != nil {
		return inspireData, err
	}
	// 应发激励汇总
	if inspireData.StageBudget, err = dao.ProductStageKpi.Ctx(ctx).Sum(dao.ProductStageKpi.Columns().StageBudget); err != nil {
		return inspireData, err
	}
	// 实发激励汇总
	if inspireData.StageQuota, err = dao.ProductStageKpi.Ctx(ctx).Sum(dao.ProductStageKpi.Columns().StageQuota); err != nil {
		return inspireData, err
	}

	inspireData.IncentiveBudget = inspireData.IncentiveBudget / 10000
	inspireData.StageBudget = inspireData.StageBudget / 10000
	inspireData.StageQuota = inspireData.StageQuota / 10000

	return inspireData, nil
}

func (a *statisticsSummationService) GetStage(ctx context.Context) (stageData []model.StageStatic, err error) {
	stageData = make([]model.StageStatic, 0)
	// 阀点名称
	stageLists, err := boot.ModeStageServer.GetAll(ctx, &v1.GetAllModeStageReq{ModeStage: &v1.ModeStageInfo{Tid: 1}})
	if err != nil {
		return stageData, err
	}
	for _, v := range stageLists.GetData() {
		info := model.StageStatic{
			StageName:   v.GetName(),
			StageBudget: 0,
			StageQuota:  0,
		}
		// 应发激励汇总
		proStageIds, err := dao.ProductStageRule.Ctx(ctx).Where(dao.ProductStageRule.Columns().ProStageId, v.Id).Array(dao.ProductStageRule.Columns().Id)
		if err != nil {
			return stageData, err
		}

		if len(proStageIds) > 0 {
			// 应发激励汇总
			query := dao.ProductStageKpi.Ctx(ctx).WhereIn(dao.ProductStageKpi.Columns().StageId, proStageIds)
			info.StageBudget, err = query.Sum(dao.ProductStageKpi.Columns().StageBudget)
			if err != nil {
				return stageData, err
			}
			// 实发激励汇总
			info.StageQuota, err = query.Sum(dao.ProductStageKpi.Columns().StageQuota)
			if err != nil {
				return stageData, err
			}

			info.StageBudget = info.StageBudget / 10000
			info.StageQuota = info.StageQuota / 10000
		}
		stageData = append(stageData, info)
	}

	return stageData, nil
}

func (a *statisticsSummationService) GetProductStage(ctx context.Context) (stageData model.ProductStageLint, err error) {
	stageData = model.ProductStageLint{
		StageName:   make([]string, 0),
		StageQuota:  make([]float64, 0),
		StageBudget: make([]float64, 0),
	}
	stageKpiLists, err := dao.ProductStageKpi.GetAll(ctx, &model.ProductStageKpiApiGetListReq{}, model.GetDataOrder{KeyName: dao.ProductStageKpi.Columns().ProId, OrderAsc: true}, 0)
	if err != nil || len(stageKpiLists) == 0 {
		return stageData, err
	}
	for _, v := range stageKpiLists {
		// 阀点名称
		stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: v.StageId})
		if err != nil {
			return stageData, err
		}
		productInfo, err := dao.Product.GetOne(ctx, model.Product{Id: v.ProId})
		if err != nil {
			return stageData, err
		}

		stageData.StageName = append(stageData.StageName, fmt.Sprintf("%s-%s", productInfo.SubName, stageInfo.Name))
		// 应发激励汇总
		stageData.StageBudget = append(stageData.StageBudget, v.StageBudget/10000)
		// 实发激励汇总
		stageData.StageQuota = append(stageData.StageQuota, v.StageQuota/10000)
	}

	return stageData, nil
}

func (a *statisticsSummationService) GetProductStageScore(ctx context.Context) (stageData []model.ProductStageScore, err error) {
	stageData = make([]model.ProductStageScore, 0)
	stageKpiLists, err := dao.ProductStageKpi.GetAll(ctx, &model.ProductStageKpiApiGetListReq{}, model.GetDataOrder{KeyName: dao.ProductStageKpi.Columns().StageScore, OrderDesc: true}, 0)
	if err != nil || len(stageKpiLists) == 0 {
		return stageData, err
	}
	for _, v := range stageKpiLists {
		info := model.ProductStageScore{}
		// 阀点名称
		stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: v.StageId})
		if err != nil {
			return stageData, err
		}
		info.StageName = stageInfo.Name
		productInfo, err := dao.Product.GetOne(ctx, model.Product{Id: v.ProId})
		if err != nil {
			return stageData, err
		}
		info.ProductName = productInfo.Name
		info.ProductSubName = productInfo.SubName
		info.StageScore = v.StageScore
		stageData = append(stageData, info)
	}

	return stageData, nil
}

func (a *statisticsSummationService) GetProductStageTop(ctx context.Context) (stageData []model.ProductStageTop, err error) {
	stageData = make([]model.ProductStageTop, 0)

	topNumber, err := dao.Config.GetKeyValueByKeyName(ctx, consts.StageTopNumber)
	if err != nil {
		g.Log("config").Error(ctx, err)
	}
	if g.IsEmpty(topNumber) {
		panic("首页阀点绩效TOP排名未配置，请先完善数据")
	}

	stageKpiLists, err := dao.ProductStageKpi.GetAll(ctx, &model.ProductStageKpiApiGetListReq{}, model.GetDataOrder{KeyName: dao.ProductStageKpi.Columns().StageScore, OrderDesc: true}, gconv.Int(topNumber))
	if err != nil || len(stageKpiLists) == 0 {
		return stageData, err
	}
	for _, v := range stageKpiLists {
		info := model.ProductStageTop{}
		// 阀点名称
		stageInfo, err := dao.ProductStageRule.GetOne(ctx, &model.ProductStageRule{Id: v.StageId})
		if err != nil {
			return stageData, err
		}
		info.StageName = stageInfo.Name
		productInfo, err := dao.Product.GetOne(ctx, model.Product{Id: v.ProId})
		if err != nil {
			return stageData, err
		}
		// 团队人数
		info.TeamNumber, err = dao.ProductMember.Ctx(ctx).Where(dao.ProductMember.Columns().ProId, v.ProId).Count(dao.ProductMember.Columns().EmpId)
		if err != nil {
			return stageData, err
		}
		info.ProductName = productInfo.Name
		info.ProductSubName = productInfo.SubName
		info.StageScore = v.StageScore
		info.ShouldSentRadio = v.ShouldSentRadio
		info.StageQuota = v.StageQuota
		info.StageBudget = v.StageBudget
		stageData = append(stageData, info)
	}

	return stageData, nil
}
