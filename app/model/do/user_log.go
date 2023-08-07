// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure of table cqgf_user_log for DAO operations like Where/Data.
type UserLog struct {
	g.Meta              `orm:"table:cqgf_user_log, do:true"`
	Id                  interface{} //
	MethodName          interface{} // 对应HTTP请求中Method参数值(1: GET 2: POST 3:PUT 4:DELETE)
	RequestUri          interface{} // http请求接口地址
	WorkNumber          interface{} // 唯一账号（员工工号）
	RequestModule       interface{} // 请求模块（一级模块）
	RequestSecondModule interface{} // 请求模块（二级模块）
	ChangeTypeName      interface{} // 操作对象全称
	RequestBody         interface{} // 具体请求参数
	Remark              interface{} // 预留补充说明信息
	CreateTime          *gtime.Time // 数据新增时间
	UpdateTime          *gtime.Time // 最后一次更新数据时间
}
