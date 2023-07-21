// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewKpiRule is the golang structure of table cqgf_crew_kpi_rule for DAO operations like Where/Data.
type CrewKpiRule struct {
	g.Meta     `orm:"table:cqgf_crew_kpi_rule, do:true"`
	Id         interface{} //
	Redio      interface{} // 比例
	LevelName  interface{} // 等级名称
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
