package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	common "github.com/lj1570693659/gfcq_protoc/common/v1"
	v1 "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
)

// CrewDutyIndex 责任指数信息责任服务
var CrewDutyIndex = crewDutyIndexService{}

type crewDutyIndexService struct{}

// GetAll 获取项目责任指数信息列表
func (s *crewDutyIndexService) GetAll(ctx context.Context, input *model.CrewDutyIndex) ([]model.CrewDutyIndexApiAll, error) {
	res := make([]model.CrewDutyIndexApiAll, 0)
	resDuty, err := boot.CrewDutyIndexServer.GetAll(ctx, &v1.GetAllCrewDutyIndexReq{
		CrewDutyIndex: &v1.CrewDutyIndexInfo{
			Remark:     input.Remark,
			ScoreIndex: gconv.Uint32(input.ScoreIndex),
			JobLevelId: gconv.Uint32(input.JobLevelId),
		},
	})
	if len(resDuty.GetData()) > 0 {
		for _, v := range resDuty.GetData() {
			info, err := boot.JobLevelServer.GetOne(ctx, &common.GetOneJobLevelReq{
				Id: gconv.Int32(v.GetJobLevelId()),
			})
			if err != nil {
				return res, err
			}
			res = append(res, model.CrewDutyIndexApiAll{
				CrewDutyInfo: model.CrewDutyIndex{
					Id:         gconv.Uint(v.Id),
					ScoreIndex: gconv.Uint(v.ScoreIndex),
					JobLevelId: gconv.Uint(v.JobLevelId),
					Arith:      v.Arith.String(),
					Remark:     v.Remark,
				},
				JobLevelName: info.GetJobLevel().Name,
			})
		}
	}
	return res, err
}

// Create 创建项目责任指数信息
func (s *crewDutyIndexService) Create(ctx context.Context, input *model.CrewDutyIndexApiChangeReq) error {
	_, err := boot.CrewDutyIndexServer.Create(ctx, &v1.CreateCrewDutyIndexReq{
		Remark:     input.Remark,
		ScoreIndex: gconv.Uint32(input.ScoreIndex),
		JobLevelId: gconv.Uint32(input.JobLevelId),
		Arith:      util.GetArith(input.Arith),
	})

	return err
}

// Modify 更新项目责任指数信息
func (s *crewDutyIndexService) Modify(ctx context.Context, input *model.CrewDutyIndexApiChangeReq) error {
	_, err := boot.CrewDutyIndexServer.Modify(ctx, &v1.ModifyCrewDutyIndexReq{
		Id:         gconv.Int32(input.ID),
		Remark:     input.Remark,
		ScoreIndex: gconv.Uint32(input.ScoreIndex),
		JobLevelId: gconv.Uint32(input.JobLevelId),
		Arith:      util.GetArith(input.Arith),
	})

	return err
}

// Delete 删除项目责任指数信息
func (s *crewDutyIndexService) Delete(ctx context.Context, input *model.CrewDutyIndexApiDeleteReq) error {
	_, err := boot.CrewDutyIndexServer.Delete(ctx, &v1.DeleteCrewDutyIndexReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
