package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/dao"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

var ProductStageRule = productStageRuleService{}

type productStageRuleService struct{}

// GetOne 项目绩效详情
func (s *productStageRuleService) GetOne(ctx context.Context, in *model.ProductStageRule) (res *model.StageInfo, err error) {
	res = &model.StageInfo{
		Single:    &model.ProductStageRule{},
		ModeStage: &model.ModeStage{},
	}
	res.Single, err = dao.ProductStageRule.GetOne(ctx, in)
	if err != nil {
		return res, err
	}

	stageInfo, err := boot.ModeStageServer.GetOne(ctx, &v1.GetOneModeStageReq{ModeStage: &v1.ModeStageInfo{Id: gconv.Int32(res.Single.ProStageId)}})
	if err != nil {
		return res, err
	}

	gconv.Struct(stageInfo.GetModeStage(), &res.ModeStage)
	return res, nil
}

// CreateDefault 创建默认数据
func (s *productStageRuleService) CreateDefault(ctx context.Context, typeId, productId uint) error {
	stage, err := boot.ModeStageServer.GetAll(ctx, &v1.GetAllModeStageReq{ModeStage: &v1.ModeStageInfo{Tid: gconv.Int32(typeId)}})
	if err != nil {
		return err
	}
	if len(stage.GetData()) == 0 {
		return errors.New("项目阶段配置信息为空，请先完善")
	}

	for _, v := range stage.GetData() {
		if err = dao.ProductStageRule.Create(ctx, model.ProductStageRule{
			Name:       v.GetName(),
			ProId:      productId,
			ProStageId: gconv.Uint(v.GetId()),
			QuotaRadio: gconv.Float64(v.GetQuotaRadio()),
			Remark:     "",
		}); err != nil {
			return err
		}
	}
	return nil
}
