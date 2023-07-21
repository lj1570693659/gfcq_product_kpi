// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductBudgetAccess is the golang structure for table product_budget_access.
type ProductBudgetAccess struct {
	Id          int         `json:"id"          ` //
	ScoreMin    uint        `json:"scoreMin"    ` // 分值下限
	ScoreMax    uint        `json:"scoreMax"    ` // 分值上限
	ScoreRange  uint        `json:"scoreRange"  ` // 分数区间包含关系（1：左闭右开，2：左开右闭）
	BudgetMin   float64     `json:"budgetMin"   ` // 预算额度下限
	BudgetMax   float64     `json:"budgetMax"   ` // 预算额度上线
	BudgetRange uint        `json:"budgetRange" ` // 额度区间包含关系（1：左闭右开，2：左开右闭）
	Remark      string      `json:"remark"      ` // 预留备注说明信息
	CreateTime  *gtime.Time `json:"createTime"  ` // 新增数据时间
	UpdateTime  *gtime.Time `json:"updateTime"  ` // 最后一次更新数据时间
}
