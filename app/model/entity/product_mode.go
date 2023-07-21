// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMode is the golang structure for table product_mode.
type ProductMode struct {
	Id         uint        `json:"id"         ` //
	Name       string      `json:"name"       ` // 开发模型名称
	Factor     float64     `json:"factor"     ` // 开发系数
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
