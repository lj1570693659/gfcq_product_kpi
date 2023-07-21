// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageKpi is the golang structure of table cqgf_product_stage_kpi for DAO operations like Where/Data.
type ProductStageKpi struct {
	g.Meta           `orm:"table:cqgf_product_stage_kpi, do:true"`
	Id               interface{} //
	ProId            interface{} // 项目主表ID
	StageId          interface{} // 项目所处阶段（cqgf_product_stage_rule.id）
	StageRadio       interface{} // 阶段比例
	StageBudget      interface{} // 阶段预算
	StageScore       interface{} // 阶段得分
	ShouldSentRadio  interface{} // 应发比例
	StageQuota       interface{} // 阶段额度
	CrewQuota        interface{} // 团队额度
	TeamBuildQuota   interface{} // 团建额度
	SupportQuota     interface{} // 业务支持额度
	PmRadio          interface{} // PM分配比例
	PmBase           interface{} // PM发放基础
	PmFloatRadio     interface{} // PM浮动比例
	PmKpiLevelId     interface{} // PM绩效等级
	PmKpiLevelName   interface{} // PM绩效等级名称
	PmKpiLevelRadio  interface{} // PM绩效等级比例
	PmIncentiveQuota interface{} // PM实际应发额度
	Remark           interface{} // 预留备注信息
	CreateTime       *gtime.Time // 新增数据时间
	UpdateTime       *gtime.Time // 最后一次更新数据时间
}
