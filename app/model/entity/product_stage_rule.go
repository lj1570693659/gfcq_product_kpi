// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageRule is the golang structure for table product_stage_rule.
type ProductStageRule struct {
	Id         uint        `json:"id"         ` //
	Name       string      `json:"name"       ` // 阶段名称
	ProId      uint        `json:"proId"      ` // 项目ID
	ProStageId uint        `json:"proStageId" ` // 项目-阶段ID
	QuotaRadio float64     `json:"quotaRadio" ` // 阶段额度占比
	Remark     string      `json:"remark"     ` // 预留备注信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
