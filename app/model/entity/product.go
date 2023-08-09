// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	Id              uint        `json:"id"              ` //
	Tid             uint        `json:"tid"             ` // 项目类型（type关联表）
	Name            string      `json:"name"            ` // 项目名称
	ProNumber       string      `json:"proNumber"       ` // 项目编号
	SubName         string      `json:"subName"         ` // 项目简称
	LcScore         uint        `json:"lcScore"         ` // 优先级评分
	LccId           uint        `json:"lccId"           ` // 项目优先级ID
	LccName         string      `json:"lccName"         ` // 项目优先级
	Invest          float64     `json:"invest"          ` // 投资额度
	NetProfit       float64     `json:"netProfit"       ` // 首年净利润
	ModeId          uint        `json:"modeId"          ` // 研发模式ID（product_mode）
	PmId            uint        `json:"pmId"            ` // PM(员工信息表ID)
	Attribute       uint        `json:"attribute"       ` // PM属性（1：全职，2：兼职）
	PmlId           uint        `json:"pmlId"           ` // 项目责任人ID（员工信息表ID）
	IncentiveBudget float64     `json:"incentiveBudget" ` // 项目激励预算
	FixBudget       float64     `json:"fixBudget"       ` //
	FixType         uint        `json:"fixType"         ` // 修正预算计算类型（1：增加，2减少）
	Status          uint        `json:"status"          ` // 项目当前状态(1:未开始 2：未立项，3：进行中 4：暂停 5：已取消 6：延迟 7：异常 8：已完成未验收 9：客户已验收 10：结项)
	ProTypeStageId  uint        `json:"proTypeStageId"  ` // 项目当前所处阶段
	Remark          string      `json:"remark"          ` // 预留备注说明信息
	CreateTime      *gtime.Time `json:"createTime"      ` // 新增数据时间
	UpdateTime      *gtime.Time `json:"updateTime"      ` // 最后一次更新数据时间
}
