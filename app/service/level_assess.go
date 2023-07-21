package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// LevelAssess 项目等级评估信息管理服务
var LevelAssess = levelAssessService{}

type levelAssessService struct{}

// GetList 获取项目等级评估信息列表
func (s *levelAssessService) GetList(ctx context.Context, input *model.LevelAssessApiGetListReq) (res []model.LevelAssessApiGetList, err error) {
	res = make([]model.LevelAssessApiGetList, 0)
	// 查询一级部门信息
	getAll, err := boot.LevelAssessServer.GetListWithoutPage(ctx, &v1.GetListWithoutLevelAssessReq{
		LevelAssess: &v1.LevelAssessInfo{
			Id:                 gconv.Int32(input.Id),
			EvaluateDimensions: input.EvaluateDimensions,
			EvaluateId:         -1,
		},
	})
	if err != nil {
		return res, err
	}

	if len(getAll.GetData()) > 0 {
		gconv.Scan(getAll.GetData(), &res)
		_, res = s.getLevelAccessTreeNode(ctx, res)
	}
	return res, nil
}

// GetOne 获取项目等级评估信息详情
func (s *levelAssessService) GetOne(ctx context.Context, input *model.LevelAssess) (res *v1.GetOneLevelAssessRes, err error) {
	info, err := boot.LevelAssessServer.GetOne(ctx, &v1.GetOneLevelAssessReq{
		LevelAssess: &v1.LevelAssessInfo{
			Id:                 gconv.Int32(input.Id),
			EvaluateDimensions: input.EvaluateDimensions,
			EvaluateCriteria:   input.EvaluateCriteria,
			ScoreCriteria:      input.ScoreCriteria,
			EvaluateId:         gconv.Int32(input.EvaluateId),
		},
	})

	return info, err
}

// Create 创建项目等级评估信息
func (s *levelAssessService) Create(ctx context.Context, input *model.LevelAssessApiChangeReq) error {
	_, err := boot.LevelAssessServer.Create(ctx, &v1.CreateLevelAssessReq{
		EvaluateDimensions: input.EvaluateDimensions,
		EvaluateCriteria:   input.EvaluateCriteria,
		ScoreCriteria:      input.ScoreCriteria,
		EvaluateId:         gconv.Int32(input.EvaluateId),
		Weight:             gconv.Float32(input.Weight),
		Remark:             input.Remark,
	})

	return err
}

// Modify 更新项目等级评估信息
func (s *levelAssessService) Modify(ctx context.Context, input *model.LevelAssessApiChangeReq) error {
	_, err := boot.LevelAssessServer.Modify(ctx, &v1.ModifyLevelAssessReq{
		Id:                 gconv.Int32(input.ID),
		EvaluateDimensions: input.EvaluateDimensions,
		EvaluateCriteria:   input.EvaluateCriteria,
		ScoreCriteria:      input.ScoreCriteria,
		EvaluateId:         gconv.Int32(input.EvaluateId),
		Weight:             gconv.Float32(input.Weight),
		Remark:             input.Remark,
	})

	return err
}

// Delete 删除项目等级评估信息
func (s *levelAssessService) Delete(ctx context.Context, input *model.LevelAssessApiDeleteReq) error {
	_, err := boot.LevelAssessServer.Delete(ctx, &v1.DeleteLevelAssessReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}

// getLevelAccessTreeNode 递归获取子节点
func (s *levelAssessService) getLevelAccessTreeNode(ctx context.Context, perms []model.LevelAssessApiGetList) (context.Context, []model.LevelAssessApiGetList) {
	//定义子节点
	for k, v := range perms {
		// 计算下级部门
		getChild, err := boot.LevelAssessServer.GetListWithoutPage(ctx, &v1.GetListWithoutLevelAssessReq{
			LevelAssess: &v1.LevelAssessInfo{
				EvaluateId: gconv.Int32(v.ID),
			},
		})
		if err != nil {
			return ctx, perms
		}
		info := make([]model.LevelAssessApiGetList, 0)
		gconv.Scan(getChild.GetData(), &info)
		perms[k].ChildLevel = info

		s.getLevelAccessTreeNode(ctx, info)
	}
	return ctx, perms
}
