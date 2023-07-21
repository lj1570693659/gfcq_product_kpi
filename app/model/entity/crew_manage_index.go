// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrewManageIndex is the golang structure for table crew_manage_index.
type CrewManageIndex struct {
	Id            uint        `json:"id"            ` //
	ScoreIndex    uint        `json:"scoreIndex"    ` // 管理指数
	ProductRoleId uint        `json:"productRoleId" ` // 项目角色ID
	Remark        string      `json:"remark"        ` // 预留备注说明信息
	CreateTime    *gtime.Time `json:"createTime"    ` // 新增数据时间
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 最后一次更新数据时间
}
