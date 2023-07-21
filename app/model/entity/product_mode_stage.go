// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductModeStage is the golang structure for table product_mode_stage.
type ProductModeStage struct {
	Id         uint        `json:"id"         ` //
	Tid        uint        `json:"tid"        ` // 项目开发模式ID
	Name       string      `json:"name"       ` // 项目阶段名称
	QuotaRadio float64     `json:"quotaRadio" ` // 阶段额度占比
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
