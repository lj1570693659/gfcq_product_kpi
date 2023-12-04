package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type Mode entity.ProductMode
type Roles entity.ProductRoles

// ModeApiChangeReq 研发模式信息变更
type ModeApiChangeReq struct {
	ID     int     `json:"id"`                                                           // 主键
	Name   string  `v:"required|length:2,16#研发模式名称不能为空|研发模式名称长度应当在:2到:16之间" json:"name"` // 研发模式名称
	Factor float64 `v:"required|between:0,1#开发系数不能为空|开发系数位于区间0~1.0之间" json:"factor"`     // 开发系数
	Remark string  `json:"remark"`                                                       // 预留备注信息
}

// ModeApiDeleteReq 删除研发模式信息
type ModeApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// ProductRolesApiChangeReq 项目角色信息变更
type ProductRolesApiChangeReq struct {
	ID      int    `json:"id"`                                                        // 主键
	Name    string `v:"required|length:2,16#角色名称不能为空|角色名称长度应当在:2到:16之间" json:"name" ` // 评价维度
	Pid     uint   `v:"required|min:0#上级评价维度不能为空|请选择评价维度" json:"pid"`                 // 上级评价维度
	Explain string `json:"explain"`                                                   // 角色与职责说明
	Remark  string `json:"remark"`                                                    // 预留备注信息
}

// ProductRolesApiGetListReq 项目角色信息列表
type ProductRolesApiGetListReq struct {
	Roles
}

// ProductRolesApiGetList 项目角色信息列表(带上下级关系)
type ProductRolesApiGetList struct {
	ID        int                      `json:"id"`            // 主键
	Name      string                   `json:"name"       `   //
	Pid       uint                     `json:"pid"        `   // 上级角色
	Explain   string                   `json:"explain"    `   // 角色与职责说明
	Remark    string                   `json:"remark"    `    // 角色与职责说明
	IsSpecial int32                    `json:"isSpecial"    ` // 角色与职责说明
	Children  []ProductRolesApiGetList `json:"children"`      // 子级评估信息
}

// ProductRolesApiDeleteReq 删除项目等级评估信息
type ProductRolesApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}

// TypeApiChangeReq 研发模式信息变更
type TypeApiChangeReq struct {
	ID     int    `json:"id"`                                                           // 主键
	Name   string `v:"required|length:2,16#研发模式名称不能为空|研发模式名称长度应当在:2到:16之间" json:"name"` // 研发模式名称
	Remark string `json:"remark"`                                                       // 预留备注信息
}

// TypeApiDeleteReq 删除研发模式信息
type TypeApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"` // 主键
}
