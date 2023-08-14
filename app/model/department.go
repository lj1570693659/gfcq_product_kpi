package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type Department entity.Department

// DepartmentApiChangeReq 部门信息变更
type DepartmentApiChangeReq struct {
	ID     int    `json:"id"`                                                              // 主键
	Name   string `v:"required|length:2,16#部门名称不能为空|部门名称长度应当在:2到:16之间" json:"name"`        // 部门名称
	Level  uint   `v:"required|in:1,2,3,4,5,6#部门层级不能为空|部门层级数据错误，请联系研发" json:"level"      ` // 部门层级
	Pid    int    `json:"pid"`                                                             // 上级部门
	Remark string `json:"remark"`                                                          // 预留备注信息
}

// DepartmentApiDeleteReq 删除部门信息
type DepartmentApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// DepartmentApiGetListReq 部门信息列表
type DepartmentApiGetListReq struct {
	Department
}

// DepartmentInfoWithEmploy 部门信息列表
type DepartmentInfoWithEmploy struct {
	EmployeeCount int32      `json:"employeeCount"` // 员工数量
	Department    Department `json:"department"`    // 上级部门信息
}

// DepartmentApiGetList 部门信息列表(带上下级关系)
type DepartmentApiGetList struct {
	ID            int                    `json:"id"`            // 主键
	Name          string                 `json:"name"`          // 部门名称
	Pid           int                    `json:"parentId"`      // 上级部门
	Remark        string                 `json:"remark"`        // 预留备注信息
	EmployeeCount int32                  `json:"employeeCount"` // 员工数量
	Children      []DepartmentApiGetList `json:"children"`      // 子级部门信息
}

// DepartmentApiGetOneRes 部门信息列表
type DepartmentApiGetOneRes struct {
	Department   Department `json:"department"`   // 部门信息
	EmployeeList []Employee `json:"employeeList"` // 部门员工信息
}
