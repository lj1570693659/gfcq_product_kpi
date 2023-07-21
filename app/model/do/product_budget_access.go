// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductBudgetAccess is the golang structure of table cqgf_product_budget_access for DAO operations like Where/Data.
type ProductBudgetAccess struct {
	g.Meta      `orm:"table:cqgf_product_budget_access, do:true"`
	Id          interface{} //
	ScoreMin    interface{} // 分值下限
	ScoreMax    interface{} // 分值上限
	ScoreRange  interface{} // 分数区间包含关系（1：左闭右开，2：左开右闭）
	BudgetMin   interface{} // 预算额度下限
	BudgetMax   interface{} // 预算额度上线
	BudgetRange interface{} // 额度区间包含关系（1：左闭右开，2：左开右闭）
	Remark      interface{} // 预留备注说明信息
	CreateTime  *gtime.Time // 新增数据时间
	UpdateTime  *gtime.Time // 最后一次更新数据时间
}
