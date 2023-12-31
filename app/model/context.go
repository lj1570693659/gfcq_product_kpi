// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

const (
	// ContextKey 上下文变量存储键名
	ContextKey = "ContextKey"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	UserInfo          *UserInfo           `json:"userInfo"`          // 系统账号信息
	RoleLevel         int                 `json:"roleLevel"`         // 系统账号信息
	EmployeeInfo      *Employee           `json:"employeeInfo"`      // 员工主体信息
	JobInfo           []entity.Job        `json:"jobInfo"`           // 员工岗位信息
	DepartmentInfo    []entity.Department `json:"departmentInfo"`    // 员工所在部门信息
	ProductRole       int                 `json:"productRole"`       // 员工参与项目主键信息
	ProductIds        []uint              `json:"productIds"`        // 员工参与项目主键信息
	ProductMemberList []*ProductMember    `json:"productMemberList"` // 员工参与项目信息
}

type UserInfo struct {
	Id           uint   // 用户ID
	EmployeeId   int    `json:"employeeId" `   // 员工ID
	UserName     string `json:"userName"     ` // 员工姓名
	Sex          uint   `json:"sex"          ` // 性别（0：未知 1：男  2：女）
	WorkNumber   string `json:"workNumber"   ` // 员工工号
	Password     string `json:"password"     ` // 登录密码
	Departs      []int  `json:"departId"     ` // 所属部门
	JobLevel     uint   `json:"jobLevel"     ` // 职级
	JobId        []uint `json:"jobId"        ` // 岗位信息
	InstructorId int    `json:"instructorId" ` // 指导老师
	Status       int    `json:"status"       ` // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
}
