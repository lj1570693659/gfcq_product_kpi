// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewManageIndex is the golang structure of table cqgf_crew_manage_index for DAO operations like Where/Data.
type CrewManageIndex struct {
	g.Meta        `orm:"table:cqgf_crew_manage_index, do:true"`
	Id            interface{} //
	ScoreIndex    interface{} // 管理指数
	ProductRoleId interface{} // 项目角色ID
	Remark        interface{} // 预留备注说明信息
	CreateTime    *gtime.Time // 新增数据时间
	UpdateTime    *gtime.Time // 最后一次更新数据时间
}
