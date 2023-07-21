package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type BudgetRadio entity.ProductStageRadio

// BudgetRadioApiChangeReq 项目激励预算信息变更
type BudgetRadioApiChangeReq struct {
	ID         int     `json:"id"`                                                                    // 主键
	Sid        int     `json:"sid"`                                                                   // 阶段ID
	ScoreMin   uint    `v:"required|min:0#阶段评分下限不能为空|评分下限大于等于0" json:"scoreMin"`                      // 得分下限
	ScoreMax   uint    `v:"required|max:100#阶段评分上限不能为空|评分上限小于等于100" json:"scoreMax"`                  // 得分上线
	ScoreRange uint    `v:"required|in:1,2,3#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左闭右闭" json:"scoreRange"` // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	QuotaRadio float64 `v:"required|min:0#分配比例不能为空|分配比例大于等于0" json:"quotaRadio"`                      // 激励额度占比
	Remark     string  `json:"remark"`                                                                // 预留备注说明信息
}

// BudgetRadioApiDeleteReq 删除项目激励预算信息
type BudgetRadioApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// BudgetRadioApiGetListReq 项目激励预算信息列表
type BudgetRadioApiGetListReq struct {
	BudgetRadio BudgetRadio `json:"BudgetRadio"`
}
