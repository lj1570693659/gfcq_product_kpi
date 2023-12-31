package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type CrewManageIndex entity.CrewManageIndex
type CrewHoursIndex entity.CrewHoursIndex
type CrewDutyIndex entity.CrewDutyIndex
type CrewSolveRule entity.CrewSolveRule
type CrewOvertimeRule entity.CrewOvertimeRule
type CrewKpiRule entity.CrewKpiRule

// CrewManageIndexApiChangeReq 管理指数预算信息变更
type CrewManageIndexApiChangeReq struct {
	ID            int    `json:"id"`                                                // 主键
	ScoreIndex    int    `v:"required|min:0#管理指数不能为空|管理指数大于等于0" json:"scoreIndex"`  // 管理指数
	ProductRoleId uint   `v:"required|min:0#项目角色不能为空|请选择项目角色" json:"productRoleId"` // 项目角色ID
	Remark        string `json:"remark"`                                            // 预留备注说明信息
}

// CrewManageIndexApiAll 管理指数预算信息变更
type CrewManageIndexApiAll struct {
	CrewManageInfo  CrewManageIndex `json:"crewManageInfo"`
	ProductRoleName string          `json:"productRoleName"`
}

// CrewManageIndexApiDeleteReq 删除管理指数预算信息
type CrewManageIndexApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// CrewHoursIndexApiChangeReq 工时指数预算信息变更
type CrewHoursIndexApiChangeReq struct {
	ID         int     `json:"id"`                                                                           // 主键
	ScoreIndex int     `v:"required|min:0#管理指数不能为空|管理指数大于等于0" json:"scoreIndex"`                             // 管理指数
	ScoreMin   float64 `v:"required|min:0#工时占比下限不能为空|工时占比下限大于等于0" json:"scoreMin"`                           // 得分下限
	ScoreMax   float64 `v:"required|max:100#工时占比上限不能为空|工时占比上限小于等于100" json:"scoreMax"`                       // 得分上线
	ScoreRange uint    `v:"required|in:1,2,3,4#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左闭右闭、左右开口" json:"scoreRange"` // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	Remark     string  `json:"remark"`                                                                       // 预留备注说明信息
}

// CrewHoursIndexApiDeleteReq 删除工时指数预算信息
type CrewHoursIndexApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// CrewDutyIndexApiChangeReq 责任指数预算信息变更
type CrewDutyIndexApiChangeReq struct {
	ID         int    `json:"id"`                                                          // 主键
	ScoreIndex int    `v:"required|min:0#责任指数不能为空|责任指数大于等于0" json:"scoreIndex"`            // 职责指数
	JobLevelId uint   `v:"required|min:0#职级不能为空|请选择对应职级" json:"jobLevelId"`                // 职级ID
	Arith      string `v:"required|in:eq,gt,lt,egt,elt,neq#运算方式不能为空|请选择运算方式" json:"arith"` // 运算方式
	Remark     string `json:"remark"`                                                      // 预留备注说明信息
}

// CrewDutyIndexApiAll 责任指数预算信息变更
type CrewDutyIndexApiAll struct {
	CrewDutyInfo CrewDutyIndex `json:"crewDutyInfo"`
	JobLevelName string        `json:"jobLevelName"`
}

// CrewDutyIndexApiDeleteReq 删除责任指数预算信息
type CrewDutyIndexApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// CrewSolveRuleApiChangeReq 解决问题贡献信息变更
type CrewSolveRuleApiChangeReq struct {
	ID     int     `json:"id"`                                            // 主键
	Redio  float64 `v:"required#浮动比例不能为空" json:"redio"`                   // 浮动比例
	Demand uint    `v:"required|in:1,2,3#贡献标准为空|请选择对应贡献标准" json:"demand"` // 贡献标准
	Remark string  `json:"remark"`                                        // 预留备注说明信息
}

// CrewSolveRuleApiDeleteReq 删除解决问题贡献信息
type CrewSolveRuleApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// CrewOvertimeRuleApiChangeReq 加班贡献信息变更
type CrewOvertimeRuleApiChangeReq struct {
	ID         int     `json:"id"`                                                                            // 主键
	ScoreMin   float64 `v:"required|float|min:0#加班贡献比例下限不能为空|请输入数据格式的加班贡献比例|加班贡献比例下限大于等于0" json:"scoreMin"`   // 得分下限
	ScoreMax   float64 `v:"required|float|max:1.0#加班贡献比例上限不能为空|请输入数据格式的加班贡献比例|加班贡献比例上限小于等于1" json:"scoreMax"` // 得分上线
	Redio      float64 `v:"required|float#浮动比例不能为空|请输入数据格式的浮动比例" json:"redio"`                                // 浮动比例
	ScoreRange uint    `v:"required|in:1,2,3,4#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左闭右闭、左右开口" json:"scoreRange"`  // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	Remark     string  `json:"remark"`                                                                        // 预留备注说明信息
}

// CrewOvertimeRuleApiDeleteReq 删除加班贡献预算信息
type CrewOvertimeRuleApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// CrewKpiRuleApiChangeReq 绩效等级信息变更
type CrewKpiRuleApiChangeReq struct {
	ID         int     `json:"id"`                                                                            // 主键
	LevelName  string  `v:"required#绩效等级名称不能为空" json:"levelName"`                                             // 绩效等级名称
	ScoreMin   int32   `v:"required|integer|min:0#绩效等级评分下限不能为空|请输入数据格式的绩效等级评分|绩效等级评分下限大于等于0" json:"scoreMin"` // 得分下限
	ScoreMax   int32   `v:"required|integer|min:1#绩效等级评分上限不能为空|请输入数据格式的绩效等级评分|绩效等级评分上限小于等于1" json:"scoreMax"` // 得分上线
	ScoreRange uint    `v:"required|in:1,2,3,4#评分区间包含关系不能为空|评分区间包含关系有左闭右开、左开右闭、左右闭口、左右开口" json:"scoreRange"`  // 得分区间包含关系（1：左闭右开，2：左开右闭,3:左闭右闭）
	Redio      float64 `v:"required|float#浮动比例不能为空|请输入数据格式的浮动比例" json:"redio"`                                // 浮动比例
	Remark     string  `json:"remark"`                                                                        // 预留备注说明信息
}

// CrewKpiRuleApiDeleteReq 删除加班贡献预算信息
type CrewKpiRuleApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}
