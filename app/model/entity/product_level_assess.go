// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductLevelAssess is the golang structure for table product_level_assess.
type ProductLevelAssess struct {
	Id                 int         `json:"id"                 ` //
	EvaluateDimensions string      `json:"evaluateDimensions" ` // 评价维度
	EvaluateCriteria   string      `json:"evaluateCriteria"   ` // 评价标准
	ScoreCriteria      string      `json:"scoreCriteria"      ` // 评分标准
	EvaluateId         uint        `json:"evaluateId"         ` // 上级评价维度
	Weight             float64     `json:"weight"             ` // 权重
	Remark             string      `json:"remark"             ` // 预留备注说明信息
	CreateTime         *gtime.Time `json:"createTime"         ` // 新增数据时间
	UpdateTime         *gtime.Time `json:"updateTime"         ` // 最新更新数据
}
