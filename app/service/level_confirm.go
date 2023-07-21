package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// LevelConfirm 项目优先级信息管理服务
var LevelConfirm = levelConfirmService{}

type levelConfirmService struct{}

// GetList 获取项目优先级信息列表
func (s *levelConfirmService) GetList(ctx context.Context, input *model.LevelConfirmApiGetListReq) (*v1.GetListLevelConfirmRes, error) {
	res := &v1.GetListLevelConfirmRes{}
	res, err := boot.LevelConfirmServer.GetList(ctx, &v1.GetListLevelConfirmReq{
		LevelConfirm: &v1.LevelConfirmInfo{
			Name:   input.LevelConfirm.Name,
			Remark: input.LevelConfirm.Remark,
		},
		Page: input.Page,
		Size: input.Size,
	})
	return res, err
}

// Create 创建项目优先级信息
func (s *levelConfirmService) Create(ctx context.Context, input *model.LevelConfirmApiChangeReq) error {
	_, err := boot.LevelConfirmServer.Create(ctx, &v1.CreateLevelConfirmReq{
		Name:          input.Name,
		ScoreMin:      gconv.Float32(input.ScoreMin),
		ScoreMax:      gconv.Float32(input.ScoreMax),
		ScoreRange:    v1.ScoreRangeEnum(input.ScoreRange),
		IsNeedPm:      v1.IsNeedPmEnum(input.IsNeedPm),
		PmDemand:      input.PmDemand,
		ProductDemand: input.ProductDemand,
		MonitDemand:   input.MonitDemand,
		IsNeedPml:     v1.IsNeedPmlEnum(input.IsNeedPml),
		Remark:        input.Remark,
	})

	return err
}

// Modify 更新项目优先级信息
func (s *levelConfirmService) Modify(ctx context.Context, input *model.LevelConfirmApiChangeReq) error {
	_, err := boot.LevelConfirmServer.Modify(ctx, &v1.ModifyLevelConfirmReq{
		Id:            gconv.Int32(input.ID),
		Name:          input.Name,
		ScoreMin:      gconv.Float32(input.ScoreMin),
		ScoreMax:      gconv.Float32(input.ScoreMax),
		ScoreRange:    v1.ScoreRangeEnum(input.ScoreRange),
		IsNeedPm:      v1.IsNeedPmEnum(input.IsNeedPm),
		PmDemand:      input.PmDemand,
		ProductDemand: input.ProductDemand,
		MonitDemand:   input.MonitDemand,
		IsNeedPml:     v1.IsNeedPmlEnum(input.IsNeedPml),
		Remark:        input.Remark,
	})

	return err
}

// Delete 删除项目优先级信息
func (s *levelConfirmService) Delete(ctx context.Context, input *model.LevelConfirmApiDeleteReq) error {
	_, err := boot.LevelConfirmServer.Delete(ctx, &v1.DeleteLevelConfirmReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
