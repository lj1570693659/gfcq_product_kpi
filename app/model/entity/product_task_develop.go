// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductTaskDevelop is the golang structure for table product_task_develop.
type ProductTaskDevelop struct {
	Id                  uint        `json:"id"                  ` //
	ProductId           int         `json:"productId"           ` // 项目唯一标志
	TaskName            string      `json:"taskName"            ` // 任务名称
	TaskCate            string      `json:"taskCate"            ` // 任务分类
	TaskDesc            string      `json:"taskDesc"            ` // 任务描述
	TaskStatus          int         `json:"taskStatus"          ` // 任务状态（1：未开启，2：运行中，3：已完成 4：暂停）
	Level               int         `json:"level"               ` // 优先级（值越大，优先级越高）
	IsSendGroupChat     int         `json:"isSendGroupChat"     ` // 1: 发送至群 2：不发送至群
	GroupChat           string      `json:"groupChat"           ` // 发送至群信息
	RemindGroupChatUser string      `json:"remindGroupChatUser" ` // 发送到群时，提醒相关关注人
	DutyEmployeeId      uint        `json:"dutyEmployeeId"      ` // 责任人员工信息
	DutyWorkNumber      string      `json:"dutyWorkNumber"      ` // 责任人员工工号
	JoinWorkNumber      string      `json:"joinWorkNumber"      ` // 关联责任人
	PalnStartTime       *gtime.Time `json:"palnStartTime"       ` // 计划开始时间
	PalnEndTime         *gtime.Time `json:"palnEndTime"         ` // 计划结束时间
	RealStartTime       *gtime.Time `json:"realStartTime"       ` // 实际开始时间
	RealEndTime         *gtime.Time `json:"realEndTime"         ` // 实际结束时间
	UpgradeFirst        string      `json:"upgradeFirst"        ` // 第一次升级@工号
	UpgradeFirstTime    *gtime.Time `json:"upgradeFirstTime"    ` // 第一次升级时间
	UpgradeTwo          string      `json:"upgradeTwo"          ` // 第二次升级@工号
	UpgradeTwoTime      *gtime.Time `json:"upgradeTwoTime"      ` // 第三次升级@工号
	UpgradeThree        string      `json:"upgradeThree"        ` // 第三次升级@工号
	UpgradeThreeTime    *gtime.Time `json:"upgradeThreeTime"    ` // 第三次升级时间
	AssignEmployeeId    int         `json:"assignEmployeeId"    ` // 指派人员工信息
	AssignWorkNumber    string      `json:"assignWorkNumber"    ` // 指派人员工工号
	Remark              string      `json:"remark"              ` // 预留备注说明信息
	CreateTime          *gtime.Time `json:"createTime"          ` // 新增数据时间
	UpdateTime          *gtime.Time `json:"updateTime"          ` // 最后一次更新数据时间
}
