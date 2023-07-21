package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	v1 "github.com/lj1570693659/gfcq_protoc/config/product/v1"
)

// ProductRoles 项目角色信息管理服务
var ProductRoles = productRolesService{}

type productRolesService struct{}

// GetList 获取项目角色信息列表
func (s *productRolesService) GetList(ctx context.Context, input *model.ProductRolesApiGetListReq) ([]model.ProductRolesApiGetList, error) {
	res := make([]model.ProductRolesApiGetList, 0)

	// 查询一级项目角色信息
	getAll, err := boot.RolesServer.GetAll(ctx, &v1.GetAllRolesReq{
		Roles: &v1.RolesInfo{
			Id:   gconv.Int32(input.Roles.Id),
			Name: input.Roles.Name,
			Pid:  -1,
		},
	})
	if err != nil {
		return res, err
	}

	if len(getAll.GetData()) > 0 {
		res = make([]model.ProductRolesApiGetList, 0)
		gconv.Scan(getAll.GetData(), &res)
		_, res = s.getProductRolesTreeNode(ctx, res)
	}
	return res, nil
}

// Create 创建项目角色信息
func (s *productRolesService) Create(ctx context.Context, input *model.ProductRolesApiChangeReq) error {
	_, err := boot.RolesServer.Create(ctx, &v1.CreateRolesReq{
		Pid:     gconv.Int32(input.Pid),
		Name:    input.Name,
		Explain: input.Explain,
		Remark:  input.Remark,
	})

	return err
}

// Modify 更新项目角色信息
func (s *productRolesService) Modify(ctx context.Context, input *model.ProductRolesApiChangeReq) error {
	_, err := boot.RolesServer.Modify(ctx, &v1.ModifyRolesReq{
		Id:      gconv.Int32(input.ID),
		Pid:     gconv.Int32(input.Pid),
		Name:    input.Name,
		Explain: input.Explain,
		Remark:  input.Remark,
	})

	return err
}

// Delete 删除项目角色信息
func (s *productRolesService) Delete(ctx context.Context, input *model.ProductRolesApiDeleteReq) error {
	_, err := boot.RolesServer.Delete(ctx, &v1.DeleteRolesReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}

// getProductRolesTreeNode 递归获取子节点
func (s *productRolesService) getProductRolesTreeNode(ctx context.Context, perms []model.ProductRolesApiGetList) (context.Context, []model.ProductRolesApiGetList) {
	//定义子节点
	for k, v := range perms {
		// 计算子角色
		getChild, err := boot.RolesServer.GetAll(ctx, &v1.GetAllRolesReq{
			Roles: &v1.RolesInfo{
				Pid: gconv.Int32(v.ID),
			},
		})
		if err != nil {
			return ctx, perms
		}
		info := make([]model.ProductRolesApiGetList, 0)
		gconv.Scan(getChild.GetData(), &info)
		perms[k].ChildLevel = info

		s.getProductRolesTreeNode(ctx, info)
	}
	return ctx, perms
}
