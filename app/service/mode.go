package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// Mode 研发模式信息管理服务
var Mode = modeService{}

type modeService struct{}

// GetAll 获取研发模式信息列表
func (s *modeService) GetAll(ctx context.Context, input *model.Mode) (*v1.GetAllModeRes, error) {
	res, err := boot.ModeServer.GetAll(ctx, &v1.GetAllModeReq{
		Mode: &v1.ModeInfo{
			Name:   input.Name,
			Remark: input.Remark,
		},
	})
	return res, err
}

// Create 创建研发模式信息
func (s *modeService) Create(ctx context.Context, input *model.ModeApiChangeReq) error {
	_, err := boot.ModeServer.Create(ctx, &v1.CreateModeReq{
		Name:   input.Name,
		Remark: input.Remark,
	})

	return err
}

// Modify 更新研发模式信息
func (s *modeService) Modify(ctx context.Context, input *model.ModeApiChangeReq) error {
	_, err := boot.ModeServer.Modify(ctx, &v1.ModifyModeReq{
		Id:     gconv.Int32(input.ID),
		Name:   input.Name,
		Remark: input.Remark,
	})

	return err
}

// Delete 删除研发模式信息
func (s *modeService) Delete(ctx context.Context, input *model.ModeApiDeleteReq) error {
	_, err := boot.ModeServer.Delete(ctx, &v1.DeleteModeReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
