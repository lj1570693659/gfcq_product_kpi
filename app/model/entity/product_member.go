// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMember is the golang structure for table product_member.
type ProductMember struct {
	Id          uint        `json:"id"          ` //
	ProId       uint        `json:"proId"       ` // 项目ID
	EmpId       uint        `json:"empId"       ` // 项目成员ID
	WorkNumber  string      `json:"workNumber"  ` // 员工工号
	Attribute   uint        `json:"attribute"   ` // 属性（1：全职，2：兼职）
	PrId        uint        `json:"prId"        ` // 项目角色ID
	PrName      string      `json:"prName"      ` // 项目角色名称
	ManageIndex uint        `json:"manageIndex" ` // 管理指数
	JbId        uint        `json:"jbId"        ` // 职级ID
	JbName      string      `json:"jbName"      ` // 职级名称
	DutyIndex   int         `json:"dutyIndex"   ` // 责任指数
	Remark      string      `json:"remark"      ` // 预留备注说明信息
	CreateTime  *gtime.Time `json:"createTime"  ` // 新增数据时间
	UpdateTime  *gtime.Time `json:"updateTime"  ` // 最后一次更新数据时间
}
