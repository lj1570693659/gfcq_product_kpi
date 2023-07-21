// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMemberKpi is the golang structure of table cqgf_product_member_kpi for DAO operations like Where/Data.
type ProductMemberKpi struct {
	g.Meta        `orm:"table:cqgf_product_member_kpi, do:true"`
	Id            interface{} //
	ProId         interface{} // 项目ID
	ProEmpId      interface{} // 小组成员ID
	ProStageId    interface{} // 项目-阶段ID
	OvertimeRadio interface{} // 工时占比
	PrId          interface{} // 项目角色ID
	PrName        interface{} // 项目角色名称
	JbId          interface{} // 职级ID
	JbName        interface{} // 职级名称
	FloatRaio     interface{} // 浮动贡献
	KpiLevelId    interface{} // 绩效等级ID(区分是PM还是成员)
	KpiLevel      interface{} // 绩效等级
	KpiRadio      interface{} // 绩效比例
	Remark        interface{} // 预留备注说明信息
	CreateTime    *gtime.Time // 新增数据时间
	UpdateTime    *gtime.Time // 最后一次更新数据时间
}
