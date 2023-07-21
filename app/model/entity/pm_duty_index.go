// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmDutyIndex is the golang structure for table pm_duty_index.
type PmDutyIndex struct {
	Id         uint        `json:"id"         ` //
	ScoreIndex uint        `json:"scoreIndex" ` // 职责指数
	JobLevelId uint        `json:"jobLevelId" ` // 职级ID
	Arith      string      `json:"arith"      ` // 运算方式
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
