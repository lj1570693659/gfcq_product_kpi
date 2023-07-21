// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewDutyIndex is the golang structure of table cqgf_crew_duty_index for DAO operations like Where/Data.
type CrewDutyIndex struct {
	g.Meta     `orm:"table:cqgf_crew_duty_index, do:true"`
	Id         interface{} //
	ScoreIndex interface{} // 职责指数
	JobLevelId interface{} // 职级ID
	Arith      interface{} // 运算方式
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
