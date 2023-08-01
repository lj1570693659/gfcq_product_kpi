// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMember is the golang structure of table cqgf_product_member for DAO operations like Where/Data.
type ProductMember struct {
	g.Meta       `orm:"table:cqgf_product_member, do:true"`
	Id           interface{} //
	ProId        interface{} // 项目ID
	EmpId        interface{} // 项目成员ID
	IsSpecial    interface{} // 1: 需要特殊处理 2：不需要特殊处理
	WorkNumber   interface{} // 员工工号
	Attribute    interface{} // 属性（1：全职，2：兼职）
	PrId         interface{} // 项目角色ID
	PrName       interface{} // 项目角色名称
	ManageIndex  interface{} // 管理指数
	JbId         interface{} // 职级ID
	JbName       interface{} // 职级名称
	DutyIndex    interface{} // 责任指数
	WorkAddress  interface{} // 工作地点
	SpecificDuty interface{} // 具体职责和职务
	Type         interface{} // 项目组内部分类使用
	PutInto      interface{} // 投入占比
	IsGuide      interface{} // 是否是主导方（1：是）
	IsSupport    interface{} // 是否是支持方（1：是）
	Remark       interface{} // 预留备注说明信息
	CreateTime   *gtime.Time // 新增数据时间
	UpdateTime   *gtime.Time // 最后一次更新数据时间
}
