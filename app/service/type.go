package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// Type 项目类型信息管理服务
var Type = typeService{}

type typeService struct{}

// GetAll 获取研发模式信息列表
func (s *typeService) GetAll(ctx context.Context, input *model.ProductType) (*v1.GetAllTypeRes, error) {
	res, err := boot.TypeServer.GetAll(ctx, &v1.GetAllTypeReq{
		Type: &v1.TypeInfo{
			Name:   input.Name,
			Remark: input.Remark,
		},
	})
	return res, err
}

// GetStageAll 获取项目阶段信息列表
func (s *typeService) GetStageAll(ctx context.Context, input *model.ProductModeStage) (*v1.GetAllModeStageRes, error) {
	res, err := boot.ModeStageServer.GetAll(ctx, &v1.GetAllModeStageReq{
		ModeStage: &v1.ModeStageInfo{
			Name:       input.Name,
			Remark:     input.Remark,
			Tid:        gconv.Int32(input.Tid),
			QuotaRadio: gconv.Float32(input.QuotaRadio),
		},
	})
	return res, err
}

// CreateModeStage 新增项目阶段信息列表
func (s *typeService) CreateModeStage(ctx context.Context, input *model.ProductModeStage) (*v1.CreateModeStageRes, error) {
	res, err := boot.ModeStageServer.Create(ctx, &v1.CreateModeStageReq{
		Name:       input.Name,
		Remark:     input.Remark,
		Tid:        gconv.Int32(input.Tid),
		QuotaRadio: gconv.Float32(input.QuotaRadio),
	})
	return res, err
}

// ModifyModeStage 编辑项目阶段信息列表
func (s *typeService) ModifyModeStage(ctx context.Context, input *model.ProductModeStage) (*v1.ModifyModeStageRes, error) {
	res, err := boot.ModeStageServer.Modify(ctx, &v1.ModifyModeStageReq{
		Name:       input.Name,
		Remark:     input.Remark,
		Tid:        gconv.Int32(input.Tid),
		Id:         gconv.Int32(input.Id),
		QuotaRadio: gconv.Float32(input.QuotaRadio),
	})
	return res, err
}

// DeleteModeStage 删除项目阶段信息列表
func (s *typeService) DeleteModeStage(ctx context.Context, input *model.ProductModeStage) (*v1.DeleteModeStageRes, error) {
	res, err := boot.ModeStageServer.Delete(ctx, &v1.DeleteModeStageReq{
		Id: gconv.Int32(input.Id),
	})
	return res, err
}
