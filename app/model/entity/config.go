// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	Id         int         `json:"id"         ` //
	KeyName    string      `json:"keyName"    ` // 配置名称
	KeyValue   string      `json:"keyValue"   ` // 对应key_name的设置值
	Remark     string      `json:"remark"     ` // 预留补充说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 数据新增时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
