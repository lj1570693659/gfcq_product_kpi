// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductStageRadio is the golang structure of table cqgf_product_stage_radio for DAO operations like Where/Data.
type ProductStageRadio struct {
	g.Meta     `orm:"table:cqgf_product_stage_radio, do:true"`
	Id         interface{} //
	Sid        interface{} // 阶段ID
	ScoreMin   interface{} // 得分下限
	ScoreMax   interface{} // 得分上线
	ScoreRange interface{} // 得分区间包含关系（1：左闭右开，2：左开右闭）
	QuotaRadio interface{} // 激励额度占比
	Remark     interface{} // 预留备注信息
	CreateTime *gtime.Time // 新增数据时间
	UpdateTime *gtime.Time // 最后一次更新数据时间
}
