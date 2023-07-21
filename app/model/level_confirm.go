package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type LevelConfirm entity.ProductLevelConfirm

// LevelConfirmApiChangeReq 项目优先级信息变更
type LevelConfirmApiChangeReq struct {
	ID            int     `json:"id"`                                                                  // 主键
	Name          string  `v:"required#项目优先级名称不能为空" json:"name"`                                       // 项目优先级
	ScoreMin      float64 `v:"required|min:0#优先级评分下限不能为空|优先级评分下限大于等于0" json:"scoreMin"`                // 得分下限
	ScoreMax      float64 `v:"required|max:100#优先级评分上限不能为空|优先级评分上限小于等于100" json:"scoreMax"`            // 得分上线
	ScoreRange    uint    `v:"required|in:1,2#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左闭右闭" json:"scoreRange"` // 得分区间包含关系（1：左闭右开，2：左开右闭）
	IsNeedPm      uint    `json:"isNeedPm"`                                                            // 是否委派PM(1:是 2：否)
	PmDemand      string  `json:"pmDemand"`                                                            // pm要求
	ProductDemand string  `json:"productDemand"`                                                       // 项目工作相关要求
	MonitDemand   string  `json:"monitDemand"`                                                         // 监控要求
	IsNeedPml     uint    `json:"isNeedPml"`                                                           // 是否需要项目负责人(1:是 2：否)
	Remark        string  `json:"remark"`                                                              // 预留备注说明信息
}

// LevelConfirmApiDeleteReq 删除项目优先级信息
type LevelConfirmApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// LevelConfirmApiGetListReq 项目优先级信息列表
type LevelConfirmApiGetListReq struct {
	Page         int32        `json:"page"` // 页码
	Size         int32        `json:"size"` // 每页显示数据大小
	LevelConfirm LevelConfirm `json:"LevelConfirm"`
}

// LevelConfirmApiGetOneRes 项目优先级信息列表 TODO
//type LevelConfirmApiGetOneRes struct {
//	LevelConfirm     LevelConfirm   `json:"LevelConfirm"`     // 项目优先级信息
//	EmployeeList []Employee `json:"employeeList"` // 项目优先级员工信息
//}
