// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure for table user_log.
type UserLog struct {
	Id                  int         `json:"id"                  ` //
	MethodName          uint        `json:"methodName"          ` // 对应HTTP请求中Method参数值(1: GET 2: POST 3:PUT 4:DELETE)
	RequestUri          string      `json:"requestUri"          ` // http请求接口地址
	WorkNumber          string      `json:"workNumber"          ` // 唯一账号（员工工号）
	RequestModule       string      `json:"requestModule"       ` // 请求模块（一级模块）
	RequestSecondModule string      `json:"requestSecondModule" ` // 请求模块（二级模块）
	ChangeTypeName      string      `json:"changeTypeName"      ` // 操作对象全称
	RequestBody         string      `json:"requestBody"         ` // 具体请求参数
	Remark              string      `json:"remark"              ` // 预留补充说明信息
	CreateTime          *gtime.Time `json:"createTime"          ` // 数据新增时间
	UpdateTime          *gtime.Time `json:"updateTime"          ` // 最后一次更新数据时间
}
