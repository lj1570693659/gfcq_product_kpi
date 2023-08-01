package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
)

type Employee entity.Employee

// EmployeeApiCreateReq 创建员工信息
type EmployeeApiCreateReq struct { // 员工姓名
	UserName     string `v:"required|length:2,16#姓名不能为空|姓名长度应当在:2到:16之间" json:"userName"     `        // 员工姓名
	WorkNumber   string `json:"workNumber"     `                                                      // 员工姓名
	Sex          uint   `v:"required|in:0,1,2#请选择性别|请选择正确的性别" json:"sex"          `                   // 性别（0：未知 1：男  2：女）
	Phone        string `v:"phone-loose#请输入正确格式的手机号码" json:"phone"        `                           // 手机号码
	Email        string `v:"email#请输入正确格式的邮箱" json:"email"        `                                   // 邮箱号码
	JobLevel     uint   `v:"bail|required|integer#请选择职级|请选择正确的职级信息" json:"jobLevel"     `             // 职级
	JobId        []int  `v:"bail|required|json#请选择所属岗位信息|请选择正确的岗位信息|请选择正确的岗位信息" json:"jobId"        ` // 岗位信息
	InstructorId int    `json:"instructorId" `                                                        // 指导老师
	Status       int    `v:"required|in:1,2,3,4#请选择在职状态|请选择正确的在职状态" json:"status"       `             // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark       string `json:"remark"       `                                                        // 预留备注信息
}

// EmployeeApiModifyReq 更新员工信息
type EmployeeApiModifyReq struct {
	ID           int    `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"     `                       // 员工姓名
	UserName     string `v:"required|length:2,16#姓名不能为空|姓名长度应当在:2到:16之间" json:"userName"     `        // 员工姓名
	WorkNumber   string `json:"workNumber"     `                                                      // 员工姓名
	Sex          uint   `v:"required|in:0,1,2#请选择性别|请选择正确的性别" json:"sex"          `                   // 性别（0：未知 1：男  2：女）
	Phone        string `v:"phone-loose#请输入正确格式的手机号码" json:"phone"        `                           // 手机号码
	Email        string `v:"email#请输入正确格式的邮箱" json:"email"        `                                   // 邮箱号码
	JobLevel     uint   `v:"bail|required|integer#请选择职级|请选择正确的职级信息" json:"jobLevel"     `             // 职级
	JobId        []int  `v:"bail|required|json#请选择所属岗位信息|请选择正确的岗位信息|请选择正确的岗位信息" json:"jobId"        ` // 岗位信息
	InstructorId int    `json:"instructorId" `                                                        // 指导老师
	Status       int    `v:"required|in:1,2,3,4#请选择在职状态|请选择正确的在职状态" json:"status"       `             // 在职状态（1：在职 2：试用期 3：实习期 4：已离职）
	Remark       string `json:"remark"       `                                                        // 预留备注信息
}

// EmployeeApiDeleteReq 删除员工信息
type EmployeeApiDeleteReq struct {
	ID string `v:"required|integer#删除数据源不能为空|删除数据源错误" json:"id"     ` // 员工姓名
}

// EmployeeApiGetListReq 员工信息列表
type EmployeeApiGetListReq struct {
	Page     int32    `json:"page"` // 员工姓名
	Size     int32    `json:"size"` // 员工姓名
	Employee Employee `json:"employee"`
}

// EmployeeApiGetOneReq 员工信息列表
type EmployeeApiGetOneReq struct {
	Employee
}

// EmployeeApiGetOneRes 员工信息列表
type EmployeeApiGetOneRes struct {
	EmployeeInfo   Employee            `json:"employeeInfo"`   // 员工主体信息
	SexName        string              `json:"sexName"`        // 性别-中文显示
	StatusName     string              `json:"statusName"`     // 性别-中文显示
	JobInfo        []entity.Job        `json:"jobInfo"`        // 员工岗位信息
	JobName        string              `json:"jobName"`        // 员工岗位信息
	LevelInfo      JobLevel            `json:"levelInfo"`      // 员工职级信息
	InstructorInfo Employee            `json:"instructorInfo"` // 员工职级信息
	DepartmentInfo []entity.Department `json:"departmentInfo"` // 员工所在部门信息
	DepartmentName string              `json:"departmentName"` // 员工岗位信息
	ProductInfo    []entity.Product    `json:"productInfo"`    // 员工涉及项目信息
}

// GetListEmployeeRes 列表接口输出数据结构
type GetListEmployeeRes struct {
	Page      int32                  `json:"Page"`
	Size      int32                  `json:"Size"`
	TotalSize int32                  `json:"TotalSize"`
	Data      []EmployeeApiGetOneRes `json:"Data"`
}
