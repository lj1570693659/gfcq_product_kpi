// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageRadio is the golang structure for table product_stage_radio.
type ProductStageRadio struct {
	Id         uint        `json:"id"         ` //
	Sid        uint        `json:"sid"        ` // 阶段ID
	ScoreMin   uint        `json:"scoreMin"   ` // 得分下限
	ScoreMax   uint        `json:"scoreMax"   ` // 得分上线
	ScoreRange uint        `json:"scoreRange" ` // 得分区间包含关系（1：左闭右开，2：左开右闭）
	QuotaRadio float64     `json:"quotaRadio" ` // 激励额度占比
	Remark     string      `json:"remark"     ` // 预留备注信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
