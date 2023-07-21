package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// BudgetRadio 激励激励应发信息管理服务
var BudgetRadio = budgetRadioService{}

type budgetRadioService struct{}

// GetAll 获取项目激励激励应发信息列表
func (s *budgetRadioService) GetAll(ctx context.Context, input *model.BudgetRadio) (*v1.GetAllStageRadioRes, error) {
	res, err := boot.StageRadioServer.GetAll(ctx, &v1.GetAllStageRadioReq{
		StageRadio: &v1.StageRadioInfo{
			Remark: input.Remark,
		},
	})
	return res, err
}

// Create 创建项目激励激励应发信息
func (s *budgetRadioService) Create(ctx context.Context, input *model.BudgetRadioApiChangeReq) error {
	_, err := boot.StageRadioServer.Create(ctx, &v1.CreateStageRadioReq{
		ScoreMin:   gconv.Uint32(input.ScoreMin),
		ScoreMax:   gconv.Uint32(input.ScoreMax),
		Sid:        gconv.Uint32(input.Sid),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		QuotaRadio: gconv.Float32(input.QuotaRadio),
		Remark:     input.Remark,
	})

	return err
}

// Modify 更新项目激励激励应发信息
func (s *budgetRadioService) Modify(ctx context.Context, input *model.BudgetRadioApiChangeReq) error {
	_, err := boot.StageRadioServer.Modify(ctx, &v1.ModifyStageRadioReq{
		Id:         gconv.Int32(input.ID),
		ScoreMin:   gconv.Uint32(input.ScoreMin),
		ScoreMax:   gconv.Uint32(input.ScoreMax),
		Sid:        gconv.Uint32(input.Sid),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		QuotaRadio: gconv.Float32(input.QuotaRadio),
		Remark:     input.Remark,
	})

	return err
}

// Delete 删除项目激励激励应发信息
func (s *budgetRadioService) Delete(ctx context.Context, input *model.BudgetRadioApiDeleteReq) error {
	_, err := boot.StageRadioServer.Delete(ctx, &v1.DeleteStageRadioReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
