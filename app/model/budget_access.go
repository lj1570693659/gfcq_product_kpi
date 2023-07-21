package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type BudgetAccess entity.ProductBudgetAccess

// BudgetAccessApiChangeReq 项目激励预算信息变更
type BudgetAccessApiChangeReq struct {
	ID          int     `json:"id"`                                                                     // 主键
	ScoreMin    uint    `v:"required|min:0#激励预算评分下限不能为空|评分下限大于等于0" json:"scoreMin"`                     // 得分下限
	ScoreMax    uint    `v:"required|max:100#激励预算评分上限不能为空|评分上限小于等于100" json:"scoreMax"`                 // 得分上线
	ScoreRange  uint    `v:"required|in:1,2,3#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左闭右闭" json:"scoreRange"`  // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	BudgetMin   float64 `v:"required|min:0#预算下限不能为空|预算下限大于等于0" json:"budgetMin"`                        // 预算下限
	BudgetMax   float64 `v:"required|max:100000000#预算上限不能为空|预算上限小于等于100000000" json:"budgetMax"`        // 预算上线
	BudgetRange uint    `v:"required|in:1,2,3#预算区间包含关系不能为空|预算区间包含关系有左闭右开、左开右闭、左闭右闭" json:"budgetRange"` // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	Remark      string  `json:"remark"`                                                                 // 预留备注说明信息
}

// BudgetAccessApiDeleteReq 删除项目激励预算信息
type BudgetAccessApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// BudgetAccessApiGetListReq 项目激励预算信息列表
type BudgetAccessApiGetListReq struct {
	Page         int32        `json:"page"` // 页码
	Size         int32        `json:"size"` // 每页显示数据大小
	BudgetAccess BudgetAccess `json:"BudgetAccess"`
}
