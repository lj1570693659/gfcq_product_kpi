// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductMemberKey is the golang structure of table cqgf_product_member_key for DAO operations like Where/Data.
type ProductMemberKey struct {
	g.Meta     `orm:"table:cqgf_product_member_key, do:true"`
	Id         interface{} //
	StageKpiId interface{} // 项目绩效ID
	ProId      interface{} // 项目ID
	ProEmpId   interface{} // 小组成员ID
	ProStageId interface{} // 项目-阶段ID
	WorkNumber interface{} // 成员工号
	Username   interface{} // 成员姓名
	KeyName    interface{} // 关键事件名称
	HappenTime *gtime.Time // 发生时间
	Type       interface{} // 主体分类（1：加班贡献 2：解决问题贡献 3：其他事件贡献）
	Property   interface{} // 事件性质（1：正向激励 2：有待提高）
	Result     interface{} // 当前关键事件的处理结果
	Remark     interface{} // 预留备注说明信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
