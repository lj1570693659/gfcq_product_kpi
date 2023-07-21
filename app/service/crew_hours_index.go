package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// CrewHoursIndex 工时指数信息管理服务
var CrewHoursIndex = crewHoursIndexService{}

type crewHoursIndexService struct{}

// GetAll 获取项目工时指数信息列表
func (s *crewHoursIndexService) GetAll(ctx context.Context, input *model.CrewHoursIndex) (*v1.GetAllCrewHoursIndexRes, error) {
	res, err := boot.CrewHoursIndexServer.GetAll(ctx, &v1.GetAllCrewHoursIndexReq{
		CrewHoursIndex: &v1.CrewHoursIndexInfo{
			Remark:     input.Remark,
			ScoreMin:   gconv.Float32(input.ScoreMin),
			ScoreMax:   gconv.Float32(input.ScoreMax),
			ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
			ScoreIndex: gconv.Uint32(input.ScoreIndex),
		},
	})
	return res, err
}

// Create 创建项目工时指数信息
func (s *crewHoursIndexService) Create(ctx context.Context, input *model.CrewHoursIndexApiChangeReq) error {
	_, err := boot.CrewHoursIndexServer.Create(ctx, &v1.CreateCrewHoursIndexReq{
		Remark:     input.Remark,
		ScoreMin:   gconv.Float32(input.ScoreMin),
		ScoreMax:   gconv.Float32(input.ScoreMax),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		ScoreIndex: gconv.Uint32(input.ScoreIndex),
	})

	return err
}

// Modify 更新项目工时指数信息
func (s *crewHoursIndexService) Modify(ctx context.Context, input *model.CrewHoursIndexApiChangeReq) error {
	_, err := boot.CrewHoursIndexServer.Modify(ctx, &v1.ModifyCrewHoursIndexReq{
		Id:         gconv.Int32(input.ID),
		Remark:     input.Remark,
		ScoreMin:   gconv.Float32(input.ScoreMin),
		ScoreMax:   gconv.Float32(input.ScoreMax),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		ScoreIndex: gconv.Uint32(input.ScoreIndex),
	})

	return err
}

// Delete 删除项目工时指数信息
func (s *crewHoursIndexService) Delete(ctx context.Context, input *model.CrewHoursIndexApiDeleteReq) error {
	_, err := boot.CrewHoursIndexServer.Delete(ctx, &v1.DeleteCrewHoursIndexReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
