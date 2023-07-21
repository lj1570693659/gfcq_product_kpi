// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductModeStage is the golang structure of table cqgf_product_mode_stage for DAO operations like Where/Data.
type ProductModeStage struct {
	g.Meta     `orm:"table:cqgf_product_mode_stage, do:true"`
	Id         interface{} //
	Tid        interface{} // 项目开发模式ID
	Name       interface{} // 项目阶段名称
	QuotaRadio interface{} // 阶段额度占比
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
