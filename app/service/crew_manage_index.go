package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewManageIndex 管理指数信息管理服务
var CrewManageIndex = crewManageIndexService{}

type crewManageIndexService struct{}

// GetAll 获取项目管理指数信息列表
func (s *crewManageIndexService) GetAll(ctx context.Context, input *model.CrewManageIndex) (*v1.GetAllCrewManageIndexRes, error) {
	res, err := boot.CrewManageIndexServer.GetAll(ctx, &v1.GetAllCrewManageIndexReq{
		CrewManageIndex: &v1.CrewManageIndexInfo{
			Remark:     input.Remark,
			ScoreIndex: gconv.Uint32(input.ScoreIndex),
		},
	})
	return res, err
}

// Create 创建项目管理指数信息
func (s *crewManageIndexService) Create(ctx context.Context, input *model.CrewManageIndexApiChangeReq) error {
	_, err := boot.CrewManageIndexServer.Create(ctx, &v1.CreateCrewManageIndexReq{
		ScoreIndex:    gconv.Uint32(input.ScoreIndex),
		ProductRoleId: gconv.Uint32(input.ProductRoleId),
		Remark:        input.Remark,
	})

	return err
}

// Modify 更新项目管理指数信息
func (s *crewManageIndexService) Modify(ctx context.Context, input *model.CrewManageIndexApiChangeReq) error {
	_, err := boot.CrewManageIndexServer.Modify(ctx, &v1.ModifyCrewManageIndexReq{
		Id:            gconv.Int32(input.ID),
		ScoreIndex:    gconv.Uint32(input.ScoreIndex),
		ProductRoleId: gconv.Uint32(input.ProductRoleId),
		Remark:        input.Remark,
	})

	return err
}

// Delete 删除项目管理指数信息
func (s *crewManageIndexService) Delete(ctx context.Context, input *model.CrewManageIndexApiDeleteReq) error {
	_, err := boot.CrewManageIndexServer.Delete(ctx, &v1.DeleteCrewManageIndexReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
