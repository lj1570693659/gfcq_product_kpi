package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

// JobLevel 职级信息管理服务
var JobLevel = jobLevelService{}

type jobLevelService struct{}

// GetList 获取职级信息列表
func (s *jobLevelService) GetList(ctx context.Context, input *model.JobLevelApiGetListReq) (*v1.GetListJobLevelRes, error) {
	res := &v1.GetListJobLevelRes{}
	res, err := boot.JobLevelServer.GetList(ctx, &v1.GetListJobLevelReq{
		JobLevel: &v1.JobLevelInfo{
			Name:   input.JobLevel.Name,
			Remark: input.JobLevel.Remark,
		},
		Page: input.Page,
		Size: input.Size,
	})
	return res, err
}

// GetOne 获取职级信息详情
func (s *jobLevelService) GetOne(ctx context.Context, input *model.JobLevel) (res *v1.GetOneJobLevelRes, err error) {
	info, err := boot.JobLevelServer.GetOne(ctx, &v1.GetOneJobLevelReq{
		Id:   gconv.Int32(input.Id),
		Name: input.Name,
	})

	return info, err
}

// Create 创建职级信息
func (s *jobLevelService) Create(ctx context.Context, input *model.JobLevelApiChangeReq) error {
	_, err := boot.JobLevelServer.Create(ctx, &v1.CreateJobLevelReq{
		Name:   input.Name,
		Remark: input.Remark,
	})

	return err
}

// Modify 更新职级信息
func (s *jobLevelService) Modify(ctx context.Context, input *model.JobLevelApiChangeReq) error {
	_, err := boot.JobLevelServer.Modify(ctx, &v1.ModifyJobLevelReq{
		Id:     gconv.Int32(input.ID),
		Name:   input.Name,
		Remark: input.Remark,
	})

	return err
}

// Delete 删除职级信息
func (s *jobLevelService) Delete(ctx context.Context, input *model.JobLevelApiDeleteReq) error {
	_, err := boot.JobLevelServer.Delete(ctx, &v1.DeleteJobLevelReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}
