// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductRoles is the golang structure of table cqgf_product_roles for DAO operations like Where/Data.
type ProductRoles struct {
	g.Meta     `orm:"table:cqgf_product_roles, do:true"`
	Id         interface{} //
	Name       interface{} //
	Pid        interface{} // 上级角色
	Explain    interface{} // 角色与职责说明
	Remark     interface{} // 预留备注信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
