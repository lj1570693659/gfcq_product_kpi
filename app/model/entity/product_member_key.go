// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMemberKey is the golang structure for table product_member_key.
type ProductMemberKey struct {
	Id         uint        `json:"id"         ` //
	StageKpiId uint        `json:"stageKpiId" ` // 项目绩效ID
	ProId      uint        `json:"proId"      ` // 项目ID
	ProEmpId   uint        `json:"proEmpId"   ` // 小组成员ID
	WorkNumber string      `json:"workNumber" ` // 成员工号
	Username   string      `json:"username"   ` // 成员姓名
	KeyName    string      `json:"keyName"    ` // 关键事件名称
	HappenTime *gtime.Time `json:"happenTime" ` // 发生时间
	Type       uint        `json:"type"       ` // 主体分类（1：加班贡献 2：解决问题贡献 3：其他事件贡献）
	Property   uint        `json:"property"   ` // 事件性质（1：正向激励 2：有待提高）
	Result     string      `json:"result"     ` // 当前关键事件的处理结果
	Remark     string      `json:"remark"     ` // 预留备注说明信息
	CreateTime *gtime.Time `json:"createTime" ` // 新增数据时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 最后一次更新数据时间
}
