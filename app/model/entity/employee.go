// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Employee is the golang structure for table employee.
type Employee struct {
	Id           int         `json:"id"           ` //
	UserName     string      `json:"userName"     ` // 员工姓名
	WorkNumber   string      `json:"workNumber"   ` // 员工工号
	Sex          uint        `json:"sex"          ` // 性别（0：未知 1：男  2：女）
	Phone        string      `json:"phone"        ` // 手机号码
	Email        string      `json:"email"        ` // 邮箱号码
	DepartId     string      `json:"departId"     ` // 所属部门
	JobLevel     uint        `json:"jobLevel"     ` // 职级
	JobId        string      `json:"jobId"        ` // 岗位信息
	InstructorId int         `json:"instructorId" ` // 指导老师
	Status       int         `json:"status"       ` // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark       string      `json:"remark"       ` // 预留备注信息
	CreateTime   *gtime.Time `json:"createTime"   ` // 新增数据时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 最后一次更新数据时间
}
