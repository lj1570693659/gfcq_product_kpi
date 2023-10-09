package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type JobLevel entity.JobLevel
type Job entity.Job

// JobLevelApiChangeReq 职级信息变更
type JobLevelApiChangeReq struct {
	ID     int    `json:"id"`                                                       // 主键
	Name   string `v:"required|length:2,16#职级名称不能为空|职级名称长度应当在:2到:16之间" json:"name"` // 职级名称
	Remark string `json:"remark"`                                                   // 预留备注信息
}

// JobLevelApiDeleteReq 删除职级信息
type JobLevelApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// JobLevelApiGetListReq 职级信息列表
type JobLevelApiGetListReq struct {
	Page     int32    `json:"page"` // 页码
	Size     int32    `json:"size"` // 每页显示数据大小
	JobLevel JobLevel `json:"jobLevel"`
}

// JobApiRes 岗位信息列表
type JobApiRes struct {
	Job        Job        `json:"job"`        // 岗位信息
	Department Department `json:"department"` // 部门信息
}

// JobApiGetListReq 岗位信息列表
type JobApiGetListReq struct {
	Job
	Page int32 `json:"page"` // 页码
	Size int32 `json:"size"` // 每页显示数据大小
}

// JobApiChangeReq 职级信息变更
type JobApiChangeReq struct {
	ID       int32  `json:"id"`                                                       // 主键
	Name     string `v:"required|length:2,16#岗位名称不能为空|岗位名称长度应当在:2到:16之间" json:"name"` // 职级名称
	DepartId int32  `v:"required|min:0#所属部门不能为空|部门信息错误" json:"departId"`              // 所属部门
	Remark   string `json:"remark"`                                                   // 预留备注信息
}

// JobApiDeleteReq 删除职级信息
type JobApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}
