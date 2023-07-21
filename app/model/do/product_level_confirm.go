// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductLevelConfirm is the golang structure of table cqgf_product_level_confirm for DAO operations like Where/Data.
type ProductLevelConfirm struct {
	g.Meta        `orm:"table:cqgf_product_level_confirm, do:true"`
	Id            interface{} //
	Name          interface{} // 项目优先级
	ScoreMin      interface{} // 得分下限
	ScoreMax      interface{} // 得分上线
	ScoreRange    interface{} // 得分区间包含关系（1：左闭右开，2：左开右闭）
	IsNeedPm      interface{} // 是否委派PM(1:是 2：否)
	PmDemand      interface{} // pm要求
	ProductDemand interface{} // 项目工作相关要求
	MonitDemand   interface{} // 监控要求
	IsNeedPml     interface{} // 是否需要项目负责人(1:是 2：否)
	Remark        interface{} // 预留备注说明信息
	CreateTime    *gtime.Time // 新增数据时间
	UpdateTime    *gtime.Time // 最新更新数据
}
