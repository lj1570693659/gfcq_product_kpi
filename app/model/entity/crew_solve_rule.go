// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewSolveRule is the golang structure for table crew_solve_rule.
type CrewSolveRule struct {
	Id         uint        `json:"id"         ` //
	Redio      float64     `json:"redio"      ` // 浮动比例
	Demand     uint        `json:"demand"     ` // 贡献标准（1：有突出贡献，2：有重要贡献，3：较低/无贡献）
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
