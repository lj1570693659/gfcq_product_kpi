// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewKpiRule is the golang structure for table crew_kpi_rule.
type CrewKpiRule struct {
	Id         uint        `json:"id"         ` //
	Redio      float64     `json:"redio"      ` // 比例
	ScoreMin   uint        `json:"scoreMin"   ` // 得分下限
	ScoreMax   uint        `json:"scoreMax"   ` // 得分上线
	ScoreRange uint        `json:"scoreRange" ` // 得分区间包含关系（1：左闭右开，2：左开右闭）
	LevelName  string      `json:"levelName"  ` // 等级名称
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
