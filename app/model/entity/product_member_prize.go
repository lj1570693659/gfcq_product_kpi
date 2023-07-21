// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMemberPrize is the golang structure for table product_member_prize.
type ProductMemberPrize struct {
	Id              uint        `json:"id"              ` //
	ProId           uint        `json:"proId"           ` // 项目ID
	IsPm            uint        `json:"isPm"            ` // 是否是PM（1：是 2：否）
	ProEmpId        uint        `json:"proEmpId"        ` // 小组成员ID
	ProStageId      uint        `json:"proStageId"      ` // 项目-阶段ID
	OvertimeRadio   float64     `json:"overtimeRadio"   ` // 工时占比
	OvertimeIndex   uint        `json:"overtimeIndex"   ` // 工时指数
	PrId            uint        `json:"prId"            ` // 项目角色ID
	PrName          string      `json:"prName"          ` // 项目角色名称
	ManageIndex     uint        `json:"manageIndex"     ` // 管理指数
	JbId            uint        `json:"jbId"            ` // 职级ID
	JbName          string      `json:"jbName"          ` // 职级名称
	DutyIndex       uint        `json:"dutyIndex"       ` // 责任指数
	BaseIndex       float64     `json:"baseIndex"       ` // 基准指数
	WeightAutoRadio float64     `json:"weightAutoRadio" ` // 权重基准（自动）
	WeightPmoRadio  float64     `json:"weightPmoRadio"  ` // 权重基准（PMO）
	SentBase        float64     `json:"sentBase"        ` // 发放基数
	RemaindQueto    float64     `json:"remaindQueto"    ` // 剩余额度
	FloatRaio       float64     `json:"floatRaio"       ` // 浮动贡献
	KpiLevelId      uint        `json:"kpiLevelId"      ` // 绩效等级ID(区分是PM还是成员)
	KpiLevel        string      `json:"kpiLevel"        ` // 绩效等级
	KpiRadio        float64     `json:"kpiRadio"        ` // 绩效比例
	SentQueto       float64     `json:"sentQueto"       ` // 实发额度
	Remark          string      `json:"remark"          ` // 预留备注说明信息
	CreateTime      *gtime.Time `json:"createTime"      ` // 新增数据时间
	UpdateTime      *gtime.Time `json:"updateTime"      ` // 最后一次更新数据时间
}
