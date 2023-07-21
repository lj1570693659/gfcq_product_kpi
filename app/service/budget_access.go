package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// BudgetAccess 激励预算信息管理服务
var BudgetAccess = budgetAccessService{}

type budgetAccessService struct{}

// GetAll 获取项目激励预算信息列表
func (s *budgetAccessService) GetAll(ctx context.Context, input *model.BudgetAccess) (*v1.GetAllBudgetAssessRes, error) {
	res, err := boot.BudgetAssessServer.GetAll(ctx, &v1.GetAllBudgetAssessReq{
		BudgetAssess: &v1.BudgetAssessInfo{
			Remark: input.Remark,
		},
	})
	return res, err
}

// GetList 获取项目激励预算信息列表
func (s *budgetAccessService) GetList(ctx context.Context, input *model.BudgetAccessApiGetListReq) (*v1.GetListBudgetAssessRes, error) {
	res := &v1.GetListBudgetAssessRes{}
	res, err := boot.BudgetAssessServer.GetList(ctx, &v1.GetListBudgetAssessReq{
		BudgetAssess: &v1.BudgetAssessInfo{
			Remark: input.BudgetAccess.Remark,
		},
		Page: input.Page,
		Size: input.Size,
	})
	return res, err
}

// Create 创建项目激励预算信息
func (s *budgetAccessService) Create(ctx context.Context, input *model.BudgetAccessApiChangeReq) error {
	_, err := boot.BudgetAssessServer.Create(ctx, &v1.CreateBudgetAssessReq{
		ScoreMin:    gconv.Uint32(input.ScoreMin),
		ScoreMax:    gconv.Uint32(input.ScoreMax),
		ScoreRange:  productV1.ScoreRangeEnum(input.ScoreRange),
		BudgetMin:   gconv.Float32(input.BudgetMin),
		BudgetMax:   gconv.Float32(input.BudgetMax),
		BudgetRange: productV1.ScoreRangeEnum(input.BudgetRange),
		Remark:      input.Remark,
	})

	return err
}

// Modify 更新项目激励预算信息
func (s *budgetAccessService) Modify(ctx context.Context, input *model.BudgetAccessApiChangeReq) error {
	_, err := boot.BudgetAssessServer.Modify(ctx, &v1.ModifyBudgetAssessReq{
		Id:          gconv.Int32(input.ID),
		ScoreMin:    gconv.Uint32(input.ScoreMin),
		ScoreMax:    gconv.Uint32(input.ScoreMax),
		ScoreRange:  productV1.ScoreRangeEnum(input.ScoreRange),
		BudgetMin:   gconv.Float32(input.BudgetMin),
		BudgetMax:   gconv.Float32(input.BudgetMax),
		BudgetRange: productV1.ScoreRangeEnum(input.BudgetRange),
		Remark:      input.Remark,
	})

	return err
}

// Delete 删除项目激励预算信息
func (s *budgetAccessService) Delete(ctx context.Context, input *model.BudgetAccessApiDeleteReq) error {
	_, err := boot.BudgetAssessServer.Delete(ctx, &v1.DeleteBudgetAssessReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
