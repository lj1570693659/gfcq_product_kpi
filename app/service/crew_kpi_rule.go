package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	product "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// CrewKpiRule 解决问题贡献信息管理服务
var CrewKpiRule = crewKpiRuleService{}

type crewKpiRuleService struct{}

// GetAll 获取项目解决问题贡献信息列表
func (s *crewKpiRuleService) GetAll(ctx context.Context, input *model.CrewKpiRule) (*v1.GetAllCrewKpiRuleRes, error) {
	res, err := boot.CrewKpiRuleServer.GetAll(ctx, &v1.GetAllCrewKpiRuleReq{
		CrewKpiRule: &v1.CrewKpiRuleInfo{},
	})
	return res, err
}

// Create 创建项目解决问题贡献信息
func (s *crewKpiRuleService) Create(ctx context.Context, input *model.CrewKpiRuleApiChangeReq) error {
	_, err := boot.CrewKpiRuleServer.Create(ctx, &v1.CreateCrewKpiRuleReq{
		LevelName:  input.LevelName,
		ScoreMin:   input.ScoreMin,
		ScoreMax:   input.ScoreMax,
		ScoreRange: product.ScoreRangeEnum(input.ScoreRange),
		Remark:     input.Remark,
		Redio:      gconv.Float32(input.Redio),
	})

	return err
}

// Modify 更新项目解决问题贡献信息
func (s *crewKpiRuleService) Modify(ctx context.Context, input *model.CrewKpiRuleApiChangeReq) error {
	_, err := boot.CrewKpiRuleServer.Modify(ctx, &v1.ModifyCrewKpiRuleReq{
		Id:         gconv.Int32(input.ID),
		LevelName:  input.LevelName,
		ScoreMin:   input.ScoreMin,
		ScoreMax:   input.ScoreMax,
		ScoreRange: product.ScoreRangeEnum(input.ScoreRange),
		Remark:     input.Remark,
		Redio:      gconv.Float32(input.Redio),
	})

	return err
}

// Delete 删除项目解决问题贡献信息
func (s *crewKpiRuleService) Delete(ctx context.Context, input *model.CrewKpiRuleApiDeleteReq) error {
	_, err := boot.CrewKpiRuleServer.Delete(ctx, &v1.DeleteCrewKpiRuleReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
