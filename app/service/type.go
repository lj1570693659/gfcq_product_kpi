package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
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

// Create 新增项目阶段信息列表
func (s *typeService) Create(ctx context.Context, input *model.TypeApiChangeReq) (*v1.CreateTypeRes, error) {
	res, err := boot.TypeServer.Create(ctx, &v1.CreateTypeReq{
		Name:   input.Name,
		Remark: input.Remark,
	})
	return res, err
}

// Modify 新增项目阶段信息列表
func (s *typeService) Modify(ctx context.Context, input *model.TypeApiChangeReq) (*v1.ModifyTypeRes, error) {
	res, err := boot.TypeServer.Modify(ctx, &v1.ModifyTypeReq{
		Name:   input.Name,
		Remark: input.Remark,
		Id:     gconv.Int32(input.ID),
	})
	return res, err
}

// Delete 新增项目阶段信息列表
func (s *typeService) Delete(ctx context.Context, input *model.TypeApiDeleteReq) (*v1.DeleteTypeRes, error) {
	res, err := boot.TypeServer.Delete(ctx, &v1.DeleteTypeReq{
		Id: gconv.Int32(input.ID),
	})
	return res, err
}

// GetStageAll 获取项目阶段信息列表
func (s *typeService) GetStageAll(ctx context.Context, input *model.ProductModeStage) ([]model.GetStage, error) {
	resData := make([]model.GetStage, 0)
	res, err := boot.ModeStageServer.GetAll(ctx, &v1.GetAllModeStageReq{
		ModeStage: &v1.ModeStageInfo{
			Name:       input.Name,
			Remark:     input.Remark,
			Tid:        gconv.Int32(input.Tid),
			QuotaRadio: gconv.Float32(input.QuotaRadio),
		},
	})

	typeLists, err := boot.TypeServer.GetAll(ctx, &v1.GetAllTypeReq{})
	if err != nil {
		return resData, err
	}
	if len(res.Data) > 0 {
		for _, v := range res.Data {
			info := model.GetStage{
				ProductModeStage: model.ProductModeStage{
					Id:         gconv.Uint(v.Id),
					Tid:        gconv.Uint(v.Tid),
					Name:       v.Name,
					QuotaRadio: util.Decimal(gconv.Float64(v.QuotaRadio)),
					Remark:     v.Remark,
				},
			}
			for _, tv := range typeLists.GetData() {
				if v.GetTid() == tv.GetId() {
					info.ProductType = model.ProductType{
						Id:     gconv.Uint(tv.Id),
						Name:   tv.Name,
						Remark: tv.Remark,
					}
				}
			}
			resData = append(resData, info)
		}
	}
	return resData, err
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
