package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type Task entity.ProductTask
type TaskNpi entity.ProductTaskNpi
type TaskProduce entity.ProductTaskProduce
type TaskDevelop entity.ProductTaskDevelop
type TaskQuality entity.ProductTaskQuality

//type Roles entity.ProductRoles

// ProductTaskField 任务查询显示字段
type ProductTaskField struct {
	Task
	ID        int    `json:"id"`        // 主键
	Name      string `json:"name"`      // 项目名称
	SubName   string `json:"subName"`   // 项目简称
	ProNumber string `json:"proNumber"` // 项目简称
}

// ProductTaskWhere 任务查询条件
type ProductTaskWhere struct {
	Task
	ID      int    `json:"id"`      // 主键
	ProID   int    `json:"proId"`   // 项目主键
	Name    string `json:"name"`    // 研发模式名称
	SubName string `json:"subName"` // 研发模式名称
	Remark  string `json:"remark"`  // 预留备注信息
	Page    int32  `json:"page"`    // 员工姓名
	Size    int32  `json:"size"`    // 员工姓名
}

// ProductTaskRes 任务查询条件
type ProductTaskRes struct {
	Task        []*ProductTaskField        `json:"task"`
	TaskNpi     []*ProductTaskNpiField     `json:"task_npi"`
	TaskProduce []*ProductTaskProduceField `json:"task_produce"`
	TaskDevelop []*ProductTaskDevelopField `json:"task_develop"`
	TaskQuality []*ProductTaskQualityField `json:"task_quality"`
}

// ProductTaskNpiField 任务查询显示字段
type ProductTaskNpiField struct {
	TaskNpi
	ID        int    `json:"id"`        // 主键
	Name      string `json:"name"`      // 项目名称
	SubName   string `json:"subName"`   // 项目简称
	ProNumber string `json:"proNumber"` // 项目简称
}

// ProductTaskProduceField 任务查询显示字段
type ProductTaskProduceField struct {
	TaskProduce
	ID        int    `json:"id"`        // 主键
	Name      string `json:"name"`      // 项目名称
	SubName   string `json:"subName"`   // 项目简称
	ProNumber string `json:"proNumber"` // 项目简称
}

// ProductTaskDevelopField 任务查询显示字段
type ProductTaskDevelopField struct {
	TaskDevelop
	ID        int    `json:"id"`        // 主键
	Name      string `json:"name"`      // 项目名称
	SubName   string `json:"subName"`   // 项目简称
	ProNumber string `json:"proNumber"` // 项目简称
}

// ProductTaskQualityField 任务查询显示字段
type ProductTaskQualityField struct {
	TaskQuality
	ID        int    `json:"id"`        // 主键
	Name      string `json:"name"`      // 项目名称
	SubName   string `json:"subName"`   // 项目简称
	ProNumber string `json:"proNumber"` // 项目简称
}
