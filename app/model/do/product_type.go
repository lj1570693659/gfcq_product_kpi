// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductType is the golang structure of table cqgf_product_type for DAO operations like Where/Data.
type ProductType struct {
	g.Meta     `orm:"table:cqgf_product_type, do:true"`
	Id         interface{} //
	Name       interface{} // 开发模型名称
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
