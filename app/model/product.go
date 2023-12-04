package model

import (
	"github.com/lj1570693659/gfcq_product_kpi/app/model/entity"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
)

type Product entity.Product
type ProductMember entity.ProductMember
type ProductRoles entity.ProductRoles
type ProductType entity.ProductType
type ProductModeStage entity.ProductModeStage

type ProductWhere struct {
	Id             uint   `json:"id"              `   // 主键
	Ids            []uint `json:"ids"              `  // 主键
	Name           string `json:"name"            `   // 项目名称
	SubName        string `json:"subName"         `   // 项目简称
	ProNumber      string `json:"proNumber"         ` // 项目编号
	LcScore        uint   `json:"lcScore"         `   // 优先级评分
	LccId          []uint `json:"lccId"           `   // 项目优先级ID
	Tid            []uint `json:"tid"            `    // 项目类型ID
	PmId           []uint `json:"pmId"            `   // PM(员工信息表ID)
	Attribute      uint   `json:"attribute"       `   // PM属性（1：全职，2：兼职）
	PmlId          []uint `json:"pmlId"           `   // 项目责任人ID（员工信息表ID）
	Status         []uint `json:"status"          `   // 项目当前状态(1:正常 2：结束收尾，3：异常收尾)
	ProTypeStageId uint   `json:"proTypeStageId"  `   // 项目当前所处阶段
}

// ProductApiGetListReq 项目清单
type ProductApiGetListReq struct {
	ProductWhere
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
}

type GetStage struct {
	ProductModeStage ProductModeStage `json:"productModeStage"`
	ProductType      ProductType      `json:"productType"`
}

// GetProduct 项目清单中关联信息
type GetProduct struct {
	ProductInfo  Product     `json:"productInfo"`  // 项目信息
	ProductMode  Mode        `json:"productMode"`  // 研发模式
	ProductPm    Employee    `json:"productPm"`    // 项目经理
	ProductPml   Employee    `json:"productPml"`   // 项目负责人
	ProductType  ProductType `json:"productType"`  // 项目类型
	ProductStage *ModeStage  `json:"productStage"` // 项目所处阶段
}

// ProductApiGetOneReq 项目详情
type ProductApiGetOneReq struct {
	Product
}

// ProductApiChangeReq 更新项目信息
type ProductApiChangeReq struct { // 员工姓名
	Id              uint    `json:"id"     `                                                                  // 员工姓名
	Tid             uint    `v:"required|min:0#项目所属类型不能为空|请选择项目所属类型" json:"tid"             `                 // 项目类型（type关联表）
	Name            string  `v:"required|length:2,64#项目名称不能为空|项目名称长度应当在:2到:64之间" json:"name"     `            // 员工姓名
	ProNumber       string  `v:"required|length:2,16#项目编号不能为空|项目编号长度应当在:2到:16之间" json:"proNumber"       `     // 项目编号
	SubName         string  `v:"required|length:2,32#项目简称不能为空|项目简称长度应当在:2到:32之间" json:"subName"     `         // 员工姓名
	LcScore         uint    `v:"required|integer#请输入优先级评分|请输入整数格式的优先级评分" json:"lcScore"          `            // 优先级评分
	LccId           uint    `json:"lccId"`                                                                    // 项目优先级ID，根据得分自动计算，不用传数据
	LccName         string  `json:"lccName"`                                                                  // 项目优先级,根据得分自动计算，不用传数据
	FixBudget       float64 `v:"float#请输入预算修正" json:"fixBudget"          `                                    // 预算修正
	FixType         uint    `v:"integer#请选择修正预算计算类型" json:"fixType"          `                                // 修正预算计算类型（1：增加，2减少）
	IncentiveBudget float64 `json:"incentiveBudget" `                                                         // 项目激励预算
	Invest          float64 `v:"required|float#请输入投资额度|请输入正确格式的投资额度" json:"invest"          `                 // 投资额度
	NetProfit       float64 `v:"required|float#请输入首年净利润|请输入正确格式的首年净利润" json:"netProfit"       `               // 首年净利润
	ModeId          uint    `v:"required#请选择项目研发模式" json:"modeId"          `                                  // 研发模式ID（product_mode）
	PmId            uint    `v:"integer#请选择项目经理" json:"pmId"            `                                     // PM(员工信息表ID)
	Attribute       uint    `v:"required-unless:pmId,0#请选择项目经理投入属性" json:"attribute"       `                  // PM属性（1：全职，2：兼职）
	PmlId           uint    `json:"pmlId"           `                                                         // 项目责任人ID（员工信息表ID）
	Status          uint    `v:"required|in:1,2,3,4,5,6,7,8,9,10#请选择项目状态|请选择正确的项目状态" json:"status"          ` // 项目当前状态(1:未开始 2：未立项，3：进行中 4：暂停 5：已取消 6：延迟 7：异常 8：已完成未验收 9：客户已验收 10：结项)
	ProTypeStageId  uint    `v:"required|integer#请选择项目当前所处阶段|请选择正确的项目阶段" json:"proTypeStageId"  `             // 项目当前所处阶段
	Remark          string  `json:"remark"       `                                                            // 预留备注信息
}

type ProductMemberWhere struct {
	ProId      uint   `json:"proId"`            // 项目ID
	ProStageId uint   `json:"proStageId"      ` // 项目-阶段ID
	WorkNumber string `json:"workNumber"`       // 项目名称
	EmpId      []uint `json:"empId"`            // 项目成员ID
	PrId       []uint `json:"prId"`             // 项目角色ID
	JbId       []uint `json:"jbId"`             // 职级ID
	Attribute  []uint `json:"attribute"`        // PM属性（1：全职，2：兼职）
}

// ProductMemberGetListReq 项目清单
type ProductMemberGetListReq struct {
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
	ProductMemberWhere
}

// ProductMemberGetListRes 项目清单
type ProductMemberGetListRes struct {
	ProductMemberInfo ProductMember      `json:"productMemberInfo"`
	ProductInfo       Product            `json:"productInfo"`
	LeaderInfo        map[string]string  `json:"leaderInfo"`
	EmployeeInfo      *v1.EmployeeInfo   `json:"employeeInfo"`
	JobLevelInfo      *v1.JobLevelInfo   `json:"jobLevelInfo"`
	DepartmentInfo    *v1.DepartmentInfo `json:"departmentInfo"`
}

// ProductMemberApiGetOneReq 项目详情
type ProductMemberApiGetOneReq struct {
	ProductMember
}

// ProductMemberApiDeleteReq 项目详情
type ProductMemberApiDeleteReq struct {
	Id    uint `v:"required|min:0#项目唯一标注不能为空|唯一标注数据错误" json:"id"             ` // 主键信息不为空
	ProId uint `v:"required#请选择项目" json:"proId"`                               // 项目ID
}

// ProductMemberApiChangeReq 更新项目组成员信息
type ProductMemberApiChangeReq struct {
	Id            uint    `json:"id"`                                                             // 员工姓名
	ProId         uint    `v:"required#请选择项目" json:"proId"`                                       // 项目ID
	EmpId         uint    `json:"empId"`                                                          // 项目成员ID
	IsSpecial     uint    `json:"isSpecial"    `                                                  // 1: 需要特殊处理 2：不需要特殊处理
	JbId          uint    `json:"jbId"`                                                           // 职级ID
	JbName        string  `json:"jbName"`                                                         // 职级名称
	DutyIndex     uint32  `json:"dutyIndex"   `                                                   // 责任指数
	WorkNumber    string  `v:"required|length:2,64#员工工号不能为空|员工工号长度应当在:2到:64之间" json:"workNumber"` // 员工工号
	Attribute     uint    `json:"attribute"`                                                      // 属性（1：全职，2：兼职）
	AttributeName string  `v:"required|in:兼职,全职#投入属性值不能为空|请选择正确的投入属性值" json:"attributeName"`      // 属性（1：全职，2：兼职）
	PrId          uint    `json:"prId"`                                                           // 项目角色ID
	PrName        string  `v:"required#请输入项目角色" json:"prName"`                                    // 项目角色名称
	ManageIndex   uint    `json:"manageIndex" `                                                   // 管理指数
	WorkAddress   string  `json:"workAddress"  `                                                  // 工作地点
	SpecificDuty  string  `json:"specificDuty" `                                                  // 具体职责和职务
	Type          string  `json:"type"         `                                                  // 项目组内部分类使用
	PutInto       float64 `json:"putInto"      `                                                  // 投入占比
	IsGuide       uint    `json:"isGuide"      `                                                  // 是否是主导方（1：是）
	IsSupport     uint    `json:"isSupport"    `                                                  // 是否是支持方（1：是）
	Remark        string  `json:"remark"`                                                         // 预留备注信息
}

// ProductMemberApiImportReq 更新项目组成员信息
type ProductMemberApiImportReq struct {
	ProId       uint                     `v:"required#请选择项目" json:"proId"`          // 项目ID
	TableData   []map[string]interface{} `v:"required#文件内容不能为空" json:"tableData"`   // 项目成员ID
	TableHeader []string                 `v:"required#文件标题不能为空" json:"tableHeader"` // 职级ID
}

// ProductMemberApiWebImportReq 更新项目组成员信息
type ProductMemberApiWebImportReq struct {
	ProId      uint                `v:"required#请选择项目" json:"proId"`         // 项目ID
	UseridList map[string][]string `v:"required#文件内容不能为空" json:"useridList"` // 项目成员ID
}

// GetDataOrder 项目绩效清单
type GetDataOrder struct {
	KeyName   string `json:"keyName"`   // 排序字段名称
	OrderDesc bool   `json:"orderDesc"` // 排序方式
	OrderAsc  bool   `json:"orderAsc"`  // 排序方式
}
