package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewSolveRule 解决问题贡献信息管理服务
var CrewSolveRule = crewSolveRuleService{}

type crewSolveRuleService struct{}

// GetAll 获取项目解决问题贡献信息列表
func (s *crewSolveRuleService) GetAll(ctx context.Context, input *model.CrewSolveRule) (*v1.GetAllCrewSolveRuleRes, error) {
	res, err := boot.CrewSolveRuleServer.GetAll(ctx, &v1.GetAllCrewSolveRuleReq{
		CrewSolveRule: &v1.CrewSolveRuleInfo{
			Remark: input.Remark,
			Redio:  gconv.Float32(input.Redio),
			Demand: v1.DemandEnum(input.Demand),
		},
	})
	return res, err
}

// Create 创建项目解决问题贡献信息
func (s *crewSolveRuleService) Create(ctx context.Context, input *model.CrewSolveRuleApiChangeReq) error {
	_, err := boot.CrewSolveRuleServer.Create(ctx, &v1.CreateCrewSolveRuleReq{
		Remark: input.Remark,
		Redio:  gconv.Float32(input.Redio),
		Demand: v1.DemandEnum(input.Demand),
	})

	return err
}

// Modify 更新项目解决问题贡献信息
func (s *crewSolveRuleService) Modify(ctx context.Context, input *model.CrewSolveRuleApiChangeReq) error {
	_, err := boot.CrewSolveRuleServer.Modify(ctx, &v1.ModifyCrewSolveRuleReq{
		Id:     gconv.Int32(input.ID),
		Remark: input.Remark,
		Redio:  gconv.Float32(input.Redio),
		Demand: v1.DemandEnum(input.Demand),
	})

	return err
}

// Delete 删除项目解决问题贡献信息
func (s *crewSolveRuleService) Delete(ctx context.Context, input *model.CrewSolveRuleApiDeleteReq) error {
	_, err := boot.CrewSolveRuleServer.Delete(ctx, &v1.DeleteCrewSolveRuleReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
