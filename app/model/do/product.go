// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table cqgf_product for DAO operations like Where/Data.
type Product struct {
	g.Meta          `orm:"table:cqgf_product, do:true"`
	Id              interface{} //
	Tid             interface{} // 项目类型（type关联表）
	Name            interface{} // 项目名称
	ProNumber       interface{} // 项目编号
	SubName         interface{} // 项目简称
	LcScore         interface{} // 优先级评分
	LccId           interface{} // 项目优先级ID
	LccName         interface{} // 项目优先级
	Invest          interface{} // 投资额度
	NetProfit       interface{} // 首年净利润
	ModeId          interface{} // 研发模式ID（product_mode）
	PmId            interface{} // PM(员工信息表ID)
	Attribute       interface{} // PM属性（1：全职，2：兼职）
	PmlId           interface{} // 项目责任人ID（员工信息表ID）
	IncentiveBudget interface{} // 项目激励预算
	FixBudget       interface{} //
	FixType         interface{} // 修正预算计算类型（1：增加，2减少）
	Status          interface{} // 项目当前状态(1:未开始 2：未立项，3：进行中 4：暂停 5：已取消 6：延迟 7：异常 8：已完成未验收 9：客户已验收 10：结项)
	ProTypeStageId  interface{} // 项目当前所处阶段
	Remark          interface{} // 预留备注说明信息
	CreateTime      *gtime.Time // 新增数据时间
	UpdateTime      *gtime.Time // 最后一次更新数据时间
}
