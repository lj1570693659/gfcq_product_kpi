package service

import (
	"context"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_product_kpi/boot"
	"github.com/lj1570693659/gfcq_product_kpi/library/util"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	"strings"
)

// Department 部门信息管理服务
var Department = departmentService{}

type departmentService struct{}

// GetList 获取部门信息列表
func (s *departmentService) GetList(ctx context.Context, input *model.DepartmentApiGetListReq) ([]model.DepartmentApiGetList, error) {
	res := make([]model.DepartmentApiGetList, 0)

	// 查询一级部门信息
	getAll, err := boot.DepertmentServer.GetListWithoutPage(ctx, &v1.GetListWithoutDepartmentReq{
		Department: &v1.DepartmentInfo{
			Id:   gconv.Int32(input.Department.Id),
			Name: input.Department.Name,
			//Pid:  -1,
		},
	})
	if err != nil {
		return res, err
	}

	if len(getAll.GetData()) > 0 {
		res = make([]model.DepartmentApiGetList, 0)
		gconv.Scan(getAll.GetData(), &res)
		for k, v := range res {
			// 计算直属上级部门员工数量
			var childCountSum int32
			getCount, err := Employee.GetEmployeeCount(ctx, gconv.Int32(v.ID))
			if err != nil {
				return res, err
			}
			res[k].EmployeeCount = getCount.GetCount() + childCountSum
		}
		res = s.getDepartTreeNode(res, 0)
	}
	return res, nil
}

// getDepartTreeNode 递归获取部门子节点
func (s *departmentService) getDepartTreeNode(tree []model.DepartmentApiGetList, pid int) []model.DepartmentApiGetList {
	// 获得列出商品分类
	var goodArr []model.DepartmentApiGetList
	for _, v := range tree {
		if v.Pid == pid {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := s.getDepartTreeNode(tree, v.ID)
			node := model.DepartmentApiGetList{
				ID:            v.ID,
				Pid:           v.Pid,
				Name:          v.Name,
				EmployeeCount: v.EmployeeCount,
				Children:      child,
			}
			goodArr = append(goodArr, node)
		}
	}
	return goodArr

}

// GetOne 获取部门信息详情
func (s *departmentService) GetOne(ctx context.Context, input *model.DepartmentApiGetListReq) (res *model.DepartmentApiGetOneRes, err error) {
	res = &model.DepartmentApiGetOneRes{}

	info, err := boot.DepertmentServer.GetOne(ctx, &v1.GetOneDepartmentReq{
		Id: gconv.Int32(input.Id),
	})
	if err != nil {
		return res, err
	}

	gconv.Struct(info.GetDepartment(), &res.Department)

	// 查询部门员工信息
	employList, err := boot.EmployeeServer.GetList(ctx, &v1.GetListEmployeeReq{
		Employee: &v1.EmployeeInfo{
			DepartId: gconv.String(res.Department.Id),
		},
	})
	if err != nil {
		return res, err
	}
	gconv.Struct(employList.GetData(), &res.EmployeeList)
	return res, err
}

// Create 创建部门信息
func (s *departmentService) Create(ctx context.Context, input *model.DepartmentApiChangeReq) error {
	_, err := boot.DepertmentServer.Create(ctx, &v1.CreateDepartmentReq{
		Pid:    gconv.Int32(input.Pid),
		Name:   input.Name,
		Level:  gconv.Int32(input.Level),
		Remark: input.Remark,
	})

	return err
}

// Modify 更新部门信息
func (s *departmentService) Modify(ctx context.Context, input *model.DepartmentApiChangeReq) error {
	_, err := boot.DepertmentServer.Modify(ctx, &v1.ModifyDepartmentReq{
		Id:     gconv.Int32(input.ID),
		Pid:    gconv.Int32(input.Pid),
		Level:  gconv.Int32(input.Level),
		Name:   input.Name,
		Remark: input.Remark,
	})

	return err
}

// Delete 删除部门信息
func (s *departmentService) Delete(ctx context.Context, input *model.DepartmentApiDeleteReq) error {

	_, err := boot.DepertmentServer.Delete(ctx, &v1.DeleteDepartmentReq{
		Id: gconv.Int32(input.ID),
	})

	return err
}

func (s *departmentService) GetDepartmentName(departId string, departmentList []*v1.DepartmentInfo) string {
	departmentNames := make([]string, 0)
	if len(departmentList) == 0 {
		return ""
	}

	departIds := strings.Split(departId, ",")
	departIds = util.DeleteIntSlice(departIds)
	if len(departIds) > 0 {
		for _, dv := range strings.Split(departId, ",") {
			for _, v := range departmentList {
				if gconv.Int32(dv) == v.Id {
					departmentNames = append(departmentNames, v.GetName())
				}
			}
		}
	}

	if len(departmentNames) == 0 {
		return ""
	}
	return strings.Join(departmentNames, ",")
}
