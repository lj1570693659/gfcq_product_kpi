// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewOvertimeRule is the golang structure of table cqgf_crew_overtime_rule for DAO operations like Where/Data.
type CrewOvertimeRule struct {
	g.Meta     `orm:"table:cqgf_crew_overtime_rule, do:true"`
	Id         interface{} //
	Redio      interface{} // 浮动比例
	ScoreMin   interface{} // 得分下限
	ScoreMax   interface{} // 得分上线
	ScoreRange interface{} // 得分区间包含关系（1：左闭右开，2：左开右闭）
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
