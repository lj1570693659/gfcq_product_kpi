// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewSolveRule is the golang structure of table cqgf_crew_solve_rule for DAO operations like Where/Data.
type CrewSolveRule struct {
	g.Meta     `orm:"table:cqgf_crew_solve_rule, do:true"`
	Id         interface{} //
	Redio      interface{} // 浮动比例
	Demand     interface{} // 贡献标准（1：有突出贡献，2：有重要贡献，3：较低/无贡献）
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
