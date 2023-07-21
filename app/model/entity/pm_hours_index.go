// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmHoursIndex is the golang structure for table pm_hours_index.
type PmHoursIndex struct {
	Id         uint        `json:"id"         ` //
	ScoreIndex uint        `json:"scoreIndex" ` // 工时指数
	ScoreMin   float64     `json:"scoreMin"   ` // 得分下限
	ScoreMax   float64     `json:"scoreMax"   ` // 得分上线
	ScoreRange uint        `json:"scoreRange" ` // 得分区间包含关系（1：左闭右开，2：左开右闭）
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
