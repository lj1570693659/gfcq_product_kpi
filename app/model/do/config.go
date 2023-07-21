// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table cqgf_config for DAO operations like Where/Data.
type Config struct {
	g.Meta     `orm:"table:cqgf_config, do:true"`
	Id         interface{} //
	KeyName    interface{} // 配置名称
	KeyValue   interface{} // 对应key_name的设置值
	Remark     interface{} // 预留补充说明信息
	CreateTime *gtime.Time // 数据新增时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
