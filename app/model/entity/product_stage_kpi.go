// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageKpi is the golang structure for table product_stage_kpi.
type ProductStageKpi struct {
	Id               uint        `json:"id"               ` //
	ProId            uint        `json:"proId"            ` // 项目主表ID
	StageId          uint        `json:"stageId"          ` // 项目所处阶段（cqgf_product_stage_rule.id）
	StageRadio       float64     `json:"stageRadio"       ` // 阶段比例
	StageBudget      float64     `json:"stageBudget"      ` // 阶段预算
	StageScore       uint        `json:"stageScore"       ` // 阶段得分
	ShouldSentRadio  float64     `json:"shouldSentRadio"  ` // 应发比例
	StageQuota       float64     `json:"stageQuota"       ` // 阶段额度
	CrewQuota        float64     `json:"crewQuota"        ` // 团队额度
	TeamBuildQuota   float64     `json:"teamBuildQuota"   ` // 团建额度
	SupportQuota     float64     `json:"supportQuota"     ` // 业务支持额度
	PmRadio          float64     `json:"pmRadio"          ` // PM分配比例
	PmBase           float64     `json:"pmBase"           ` // PM发放基础
	PmFloatRadio     float64     `json:"pmFloatRadio"     ` // PM浮动比例
	PmKpiLevelId     uint        `json:"pmKpiLevelId"     ` // PM绩效等级
	PmKpiLevelScore  uint        `json:"pmKpiLevelScore"  ` // PM绩效得分
	PmKpiLevelName   string      `json:"pmKpiLevelName"   ` // PM绩效等级名称
	PmKpiLevelRadio  float64     `json:"pmKpiLevelRadio"  ` // PM绩效等级比例
	PmIncentiveQuota float64     `json:"pmIncentiveQuota" ` // PM实际应发额度
	Remark           string      `json:"remark"           ` // 预留备注信息
	CreateTime       *gtime.Time `json:"createTime"       ` // 新增数据时间
	UpdateTime       *gtime.Time `json:"updateTime"       ` // 最后一次更新数据时间
}
