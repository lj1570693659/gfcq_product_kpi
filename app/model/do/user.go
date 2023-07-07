// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table cqgf_user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:cqgf_user, do:true"`
	Id         interface{} //
	EmployeeId interface{} // 员工ID
	UserName   interface{} // 员工姓名
	WorkNumber interface{} // 员工工号
	Password   interface{} // 密码
	Sex        interface{} // 性别（0：未知 1：男  2：女）
	Phone      interface{} // 手机号码
	Email      interface{} // 邮箱号码
	Status     interface{} // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark     interface{} // 预留备注信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
