// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMemberPrize is the golang structure of table cqgf_product_member_prize for DAO operations like Where/Data.
type ProductMemberPrize struct {
	g.Meta          `orm:"table:cqgf_product_member_prize, do:true"`
	Id              interface{} //
	ProId           interface{} // 项目ID
	IsPm            interface{} // 是否是PM（1：是 2：否）
	ProEmpId        interface{} // 小组成员ID
	ProStageId      interface{} // 项目-阶段ID
	OvertimeRadio   interface{} // 工时占比
	OvertimeIndex   interface{} // 工时指数
	PrId            interface{} // 项目角色ID
	PrName          interface{} // 项目角色名称
	ManageIndex     interface{} // 管理指数
	JbId            interface{} // 职级ID
	JbName          interface{} // 职级名称
	DutyIndex       interface{} // 责任指数
	BaseIndex       interface{} // 基准指数
	WeightAutoRadio interface{} // 权重基准（自动）
	WeightPmoRadio  interface{} // 权重基准（PMO）
	SentBase        interface{} // 发放基数
	RemaindQueto    interface{} // 剩余额度
	FloatRaio       interface{} // 浮动贡献
	KpiLevelId      interface{} // 绩效等级ID(区分是PM还是成员)
	KpiLevel        interface{} // 绩效等级
	KpiRadio        interface{} // 绩效比例
	SentQueto       interface{} // 实发额度
	Remark          interface{} // 预留备注说明信息
	CreateTime      *gtime.Time // 新增数据时间
	UpdateTime      *gtime.Time // 最后一次更新数据时间
}
