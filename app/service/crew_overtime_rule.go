package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	productV1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// CrewOvertimeRule 加班贡献管理服务
var CrewOvertimeRule = crewOvertimeRuleService{}

type crewOvertimeRuleService struct{}

// GetAll 获取项目工时指数信息列表
func (s *crewOvertimeRuleService) GetAll(ctx context.Context, input *model.CrewOvertimeRule) (*v1.GetAllCrewOvertimeRuleRes, error) {
	res, err := boot.CrewOvertimeRuleServer.GetAll(ctx, &v1.GetAllCrewOvertimeRuleReq{
		CrewOvertimeRule: &v1.CrewOvertimeRuleInfo{
			Remark:     input.Remark,
			ScoreMin:   gconv.Float32(input.ScoreMin),
			ScoreMax:   gconv.Float32(input.ScoreMax),
			ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
			Redio:      gconv.Float32(input.Redio),
		},
	})
	return res, err
}

// Create 创建项目工时指数信息
func (s *crewOvertimeRuleService) Create(ctx context.Context, input *model.CrewOvertimeRuleApiChangeReq) error {
	_, err := boot.CrewOvertimeRuleServer.Create(ctx, &v1.CreateCrewOvertimeRuleReq{
		Remark:     input.Remark,
		ScoreMin:   gconv.Float32(input.ScoreMin),
		ScoreMax:   gconv.Float32(input.ScoreMax),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		Redio:      gconv.Float32(input.Redio),
	})

	return err
}

// Modify 更新项目工时指数信息
func (s *crewOvertimeRuleService) Modify(ctx context.Context, input *model.CrewOvertimeRuleApiChangeReq) error {
	_, err := boot.CrewOvertimeRuleServer.Modify(ctx, &v1.ModifyCrewOvertimeRuleReq{
		Id:         gconv.Int32(input.ID),
		Remark:     input.Remark,
		ScoreMin:   gconv.Float32(input.ScoreMin),
		ScoreMax:   gconv.Float32(input.ScoreMax),
		ScoreRange: productV1.ScoreRangeEnum(input.ScoreRange),
		Redio:      gconv.Float32(input.Redio),
	})

	return err
}

// Delete 删除项目工时指数信息
func (s *crewOvertimeRuleService) Delete(ctx context.Context, input *model.CrewOvertimeRuleApiDeleteReq) error {
	_, err := boot.CrewOvertimeRuleServer.Delete(ctx, &v1.DeleteCrewOvertimeRuleReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
