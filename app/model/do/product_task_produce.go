// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductTaskProduce is the golang structure of table cqgf_product_task_produce for DAO operations like Where/Data.
type ProductTaskProduce struct {
	g.Meta              `orm:"table:cqgf_product_task_produce, do:true"`
	Id                  interface{} //
	ProductId           interface{} // 项目唯一标志
	TaskName            interface{} // 任务名称
	TaskCate            interface{} // 任务分类
	TaskDesc            interface{} // 任务描述
	TaskStatus          interface{} // 任务状态（1：未开启，2：运行中，3：已完成 4：暂停）
	Level               interface{} // 优先级（值越大，优先级越高）
	IsSendGroupChat     interface{} // 1: 发送至群 2：不发送至群
	GroupChat           interface{} // 发送至群信息
	RemindGroupChatUser interface{} // 发送到群时，提醒相关关注人
	DutyEmployeeId      interface{} // 责任人员工信息
	DutyWorkNumber      interface{} // 责任人员工工号
	JoinWorkNumber      interface{} // 关联责任人
	PalnStartTime       *gtime.Time // 计划开始时间
	PalnEndTime         *gtime.Time // 计划结束时间
	RealStartTime       *gtime.Time // 实际开始时间
	RealEndTime         *gtime.Time // 实际结束时间
	UpgradeFirst        interface{} // 第一次升级@工号
	UpgradeFirstTime    *gtime.Time // 第一次升级时间
	UpgradeTwo          interface{} // 第二次升级@工号
	UpgradeTwoTime      *gtime.Time // 第三次升级@工号
	UpgradeThree        interface{} // 第三次升级@工号
	UpgradeThreeTime    *gtime.Time // 第三次升级时间
	AssignEmployeeId    interface{} // 指派人员工信息
	AssignWorkNumber    interface{} // 指派人员工工号
	Remark              interface{} // 预留备注说明信息
	CreateTime          *gtime.Time // 新增数据时间
	UpdateTime          *gtime.Time // 最后一次更新数据时间
}
