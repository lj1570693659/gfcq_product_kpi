package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type LevelAssess entity.ProductLevelAssess

// LevelAssessApiChangeReq 项目等级评估信息变更
type LevelAssessApiChangeReq struct {
	ID                 int     `json:"id"`                                                                        // 主键
	EvaluateDimensions string  `v:"required|length:2,16#评价维度不能为空|评价维度名称长度应当在:2到:16之间" json:"evaluateDimensions" ` // 评价维度
	EvaluateId         uint    `json:"evaluateId"`                                                                // 上级评价维度
	EvaluateCriteria   string  `json:"evaluateCriteria"`                                                          // 评价标准
	ScoreCriteria      string  `v:"required-without:evaluateId" json:"scoreCriteria"`                             // 评分标准
	Weight             float64 `v:"required-without:evaluateId" json:"weight"`                                    // 权重
	Remark             string  `json:"remark"`                                                                    // 预留备注信息
}

// LevelAssessApiDeleteReq 删除项目等级评估信息
type LevelAssessApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// LevelAssessApiGetListReq 部门信息列表
type LevelAssessApiGetListReq struct {
	LevelAssess
}

// LevelAssessApiGetList 项目等级评估信息列表(带上下级关系)
type LevelAssessApiGetList struct {
	ID                 int                     `json:"id"`                  // 主键
	EvaluateDimensions string                  `json:"evaluateDimensions" ` // 评价维度
	EvaluateCriteria   string                  `json:"evaluateCriteria"   ` // 评价标准
	ScoreCriteria      string                  `json:"scoreCriteria"      ` // 评分标准
	EvaluateId         uint                    `json:"evaluateId"         ` // 上级评价维度
	Weight             float64                 `json:"weight"             ` // 权重
	Remark             string                  `json:"remark"             ` // 预留备注说明信息
	EmployeeCount      int32                   `json:"employeeCount"`       // 员工数量
	Children           []LevelAssessApiGetList `json:"children"`            // 子级评估信息
}
