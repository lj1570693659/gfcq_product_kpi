// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageRule is the golang structure of table cqgf_product_stage_rule for DAO operations like Where/Data.
type ProductStageRule struct {
	g.Meta     `orm:"table:cqgf_product_stage_rule, do:true"`
	Id         interface{} //
	Name       interface{} // 阶段名称
	ProId      interface{} // 项目ID
	ProStageId interface{} // 项目-阶段ID
	QuotaRadio interface{} // 阶段额度占比
	Remark     interface{} // 预留备注信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
