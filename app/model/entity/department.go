// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Department is the golang structure for table department.
type Department struct {
	Id         int         `json:"id"         ` //
	Name       string      `json:"name"       ` // 部门名称
	Pid        int         `json:"pid"        ` // 上级部门
	Level      uint        `json:"level"      ` // 部门层级
	Remark     string      `json:"remark"     ` // 预留备注信息
	CreateTime *gtime.Time `json:"createTime" ` // 数据新增时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
