// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductLevelAssess is the golang structure of table cqgf_product_level_assess for DAO operations like Where/Data.
type ProductLevelAssess struct {
	g.Meta             `orm:"table:cqgf_product_level_assess, do:true"`
	Id                 interface{} //
	EvaluateDimensions interface{} // 评价维度
	EvaluateCriteria   interface{} // 评价标准
	ScoreCriteria      interface{} // 评分标准
	EvaluateId         interface{} // 上级评价维度
	Weight             interface{} // 权重
	Remark             interface{} // 预留备注说明信息
	CreateTime         *gtime.Time // 新增数据时间
	UpdateTime         *gtime.Time // 最新更新数据
}
