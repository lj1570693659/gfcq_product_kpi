package service

import (
	"context"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

// Job 职级信息管理服务
var Job = jobService{}

type jobService struct{}

// GetList 获取职级信息列表
//func (s *jobService) GetList(ctx context.Context, input *model.JobApiGetListReq) (*v1.GetListJobRes, error) {
//	res := &v1.GetListJobRes{}
//	res, err := boot.JobServer.GetList(ctx, &v1.GetListJobReq{
//		Job: &v1.JobInfo{
//			Name:   input.Job.Name,
//			Remark: input.Job.Remark,
//		},
//		Page: input.Page,
//		Size: input.Size,
//	})
//	return res, err
//}

// GetAll 获取职级信息列表
func (s *jobService) GetAll(ctx context.Context, input *model.Job) (*v1.GetAllJobRes, error) {
	res := &v1.GetAllJobRes{}
	res, err := boot.JobServer.GetAll(ctx, &v1.GetAllJobReq{
		Job: &v1.JobInfo{
			Name:   input.Name,
			Remark: input.Remark,
		},
	})
	return res, err
}

//// GetOne 获取职级信息详情
//func (s *JobService) GetOne(ctx context.Context, input *model.Job) (res *v1.GetOneJobRes, err error) {
//	info, err := boot.JobServer.GetOne(ctx, &v1.GetOneJobReq{
//		Id:   gconv.Int32(input.Id),
//		Name: input.Name,
//	})
//
//	return info, err
//}

// Create 创建职级信息
//func (s *jobService) Create(ctx context.Context, input *model.JobApiChangeReq) error {
//	_, err := boot.JobServer.Create(ctx, &v1.CreateJobReq{
//		Name:   input.Name,
//		Remark: input.Remark,
//	})
//
//	return err
//}
//
//// Modify 更新职级信息
//func (s *jobService) Modify(ctx context.Context, input *model.JobApiChangeReq) error {
//	_, err := boot.JobServer.Modify(ctx, &v1.ModifyJobReq{
//		Id:     gconv.Int32(input.ID),
//		Name:   input.Name,
//		Remark: input.Remark,
//	})
//
//	return err
//}

// Delete 删除职级信息
//func (s *jobService) Delete(ctx context.Context, input *model.JobApiDeleteReq) error {
//	_, err := boot.JobServer.Delete(ctx, &v1.DeleteJobReq{
//		Id: gconv.Int32(input.ID),
//	})
//
//	return err
//}
