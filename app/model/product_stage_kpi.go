package model

import "github.com/lj1570693659/gfcq_product_kpi/app/model/entity"

// ProductStageKpi 项目阶段绩效
type ProductStageKpi entity.ProductStageKpi
type ProductMemberKpi entity.ProductMemberKpi
type ProductMemberPrize entity.ProductMemberPrize
type ProductMemberKey entity.ProductMemberKey

// ProductStageRule 单个具体项目个性化激励额度占比配置信息
type ProductStageRule entity.ProductStageRule
type ModeStage entity.ProductModeStage

type ProductStageKpiWhere struct {
	ID           uint   `json:"id"               ` //
	ProId        []uint `json:"proId"`             // 项目主表ID
	StageId      []uint `json:"stageId"`           // 项目所处阶段（cqgf_product_stage_rule.id）
	StageScore   uint   `json:"stageScore"`        // 阶段得分
	PmKpiLevelId []uint `json:"pmKpiLevelId"`      // PM绩效等级
	Remark       string `json:"remark"`            // 预留备注信息
}

// ProductStageKpiApiGetListReq 项目绩效清单
type ProductStageKpiApiGetListReq struct {
	ProductStageKpiWhere
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
}

type StageInfo struct {
	Single    *ProductStageRule
	ModeStage *ModeStage
}

// ProductStageKpiList 项目绩效清单
type ProductStageKpiList struct {
	ProductStageKpi *ProductStageKpi  `json:"productStageKpi"` // 绩效详情 - 基本信息
	ProductInfo     Product           `json:"productInfo"`     // 项目信息
	StageInfo       *ProductStageRule `json:"stageInfo"`       // 项目阶段信息
}

type ProductInfo struct {
	Product     *Product     `json:"product"`     // 项目信息
	ProductType *ProductType `json:"productType"` // 项目信息
	ProductMode *Mode        `json:"productMode"` // 项目信息
}

// ProductStageKpiInfo 项目绩效详情
type ProductStageKpiInfo struct {
	ProductInfo        *ProductInfo        `json:"productInfo"`        // 项目信息
	PmInfo             *Employee           `json:"pmInfo"`             // 项目经理信息
	PmlInfo            *Employee           `json:"pmlInfo"`            // 项目负责人信息
	StageInfo          *StageInfo          `json:"stageInfo"`          // 项目阶段信息
	ProductStageKpi    *ProductStageKpi    `json:"productStageKpi"`    // 绩效详情 - 基本信息
	ProductMemberKpi   *ProductMemberKpi   `json:"productMemberKpi"`   // 绩效详情 - 团队成员绩效
	ProductMemberPrize *ProductMemberPrize `json:"productMemberPrize"` // 绩效详情 - 团队成员奖金
	ProductMemberKey   *ProductMemberKey   `json:"productMemberKey"`   // 绩效详情 - 关键事件清单
}

// ProductStageKpiApiChangeReq 更新项目阶段绩效
// 必传参数：阶段、阶段得分、PM分配比例、浮动贡献、绩效等级
type ProductStageKpiApiChangeReq struct {
	ID           uint    `json:"id"               `                                              //
	ProId        uint    `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"`     // 项目主表ID
	StageId      uint    `v:"required|integer#请选择绩效对应项目阀点|请选择正确的项目阀点" json:"stageId"`            // 项目所处阶段（cqgf_product_stage_rule.id）
	StageScore   uint    `v:"required|integer|min:0#请输入阶段得分阀点|得分为正整数|得分不能小于0" json:"stageScore"` // 阶段得分
	PmRadio      float64 `v:"required|min:0#请输入项目经理分配比例|项目经理分配比例不能小于0" json:"pmRadio"`           // PM分配比例
	PmFloatRadio float64 `v:"required|min:0#请输入项目经理浮动比例|项目经理浮动比例不能小于0" json:"pmFloatRadio"`      // PM浮动比例
	PmKpiLevelId uint    `v:"required|min:0#请选择项目经理绩效等级|项目经理绩效等级不能小于0" json:"pmKpiLevelId"`      // PM绩效等级
	Remark       string  `json:"remark"`                                                         // 预留备注信息
}

// ----------------------------------团队成员绩效相关数据结构--------------------------------------

// ProductMemberKpiApiGetListReq 项目绩效清单
type ProductMemberKpiApiGetListReq struct {
	ProductMemberKpi
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
}

// ProductMemberExport 团队成员信息导出
type ProductMemberExport struct {
	ProId uint `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"` // 项目主表ID
}

type ProductMemberKpiInfo struct {
	ProductMemberKpi *ProductMemberKpi `json:"productMemberKpi"`
	ProductMember    *ProductMember    `json:"productMember"`
	DepartmentName   string            `json:"departmentName"`
	UserName         string            `json:"userName"`
}

type ProductMemberInfo struct {
	Employee        Employee        `json:"employee"`
	ProductMember   ProductMember   `json:"productMember"`
	ProductRoles    ProductRoles    `json:"productRoles"`
	CrewManageIndex CrewManageIndex `json:"crewManageIndex"`
}

// ProductMemberKpiImportReq 团队成员绩效导入
type ProductMemberKpiImportReq struct {
	ID          uint                     `json:"id"               `                                          // 主键
	ProId       uint                     `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"` // 项目主表ID
	StageId     uint                     `v:"required|integer#请选择绩效对应项目阀点|请选择正确的项目阀点" json:"stageId"`        // 项目所处阶段（cqgf_product_stage_rule.id）
	TableData   []map[string]interface{} `v:"required#文件内容不能为空" json:"tableData"`                            // 项目成员ID
	TableHeader []string                 `v:"required#文件标题不能为空" json:"tableHeader"`                          // 职级ID
}

// ProductMemberKpiChangeReq 团队成员绩效
type ProductMemberKpiChangeReq struct {
	ProductMemberKey ProductMemberKeyChangeReq `json:"productMemberKey"`
	ID               uint                      `json:"id"               `                                                          // 主键
	IsPm             uint                      `json:"isPm"          `                                                             // 是否是PM（1：是 2：否）
	ProId            uint                      `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"`                 // 项目主表ID
	ProStageId       uint                      `v:"required|integer#请选择绩效对应项目阀点|请选择正确的项目阀点" json:"proStageId"`                     // 项目所处阶段（cqgf_product_stage_rule.id）
	WorkNumber       string                    `v:"required|length:2,64#员工工号不能为空|员工工号长度应当在:2到:64之间" json:"workNumber"`             // 员工工号
	ProEmpId         uint                      `json:"proEmpId"      `                                                             // 小组成员ID
	OvertimeRadio    float64                   `v:"required|float|max:1.0#工时占比不能为空|请输入正确的工时占比|工时占比不能大于1.0" json:"overtimeRadio" `  // 工时占比
	PrId             uint                      `json:"prId"          `                                                             // 项目角色ID
	PrName           string                    `json:"prName"        `                                                             // 项目角色名称
	JbId             uint                      `json:"jbId"          `                                                             // 职级ID
	JbName           string                    `json:"jbName"        `                                                             // 职级名称
	FloatRaio        float64                   `v:"required|float|max:1.0#浮动贡献不能为空|请输入正确的浮动贡献值|浮动贡献不能大于1.0" json:"floatRaio"     ` // 浮动贡献
	KpiLevelId       int32                     `json:"kpiLevelId"    `                                                             // 绩效等级ID(区分是PM还是成员)
	KpiLevel         string                    `v:"required|in:S,A,B,C,D,E,F#绩效等级不能为空|请输入正确的绩效等级" json:"kpiLevel"      `           // 绩效等级
	KpiRadio         float64                   `json:"kpiRadio"      `                                                             // 绩效比例
	Remark           string                    `json:"remark"`                                                                     // 预留备注信息
}

// ProductMemberKeyChangeReq 团队成员绩效导入
type ProductMemberKeyChangeReq struct {
	ID         uint   `json:"id"               ` // 主键
	WorkNumber string `json:"workNumber"`
	KeyName    string `json:"keyName"`
	Type       string `json:"type"`
	Property   string `json:"property"`
	Result     string `json:"result"`
	HappenTime string `json:"happenTime"`
}

// ProductMemberKeyListsReq 团队成员绩效导入
type ProductMemberKeyListsReq struct {
	ProductMemberKey
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
}

// ProductMemberKeyList 团队成员绩效导入
type ProductMemberKeyList struct {
	ProductMemberKey ProductMemberKey `json:"productMemberKey"`
	ProductMemberKpi ProductMemberKpi `json:"productMemberKpi"`
}

// ----------------------------------团队成员奖金相关数据结构--------------------------------------

type ProductMemberPrizeChangeReq struct {
	ID           uint    `json:"id"               `                                              //
	ProId        uint    `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"`     // 项目主表ID
	StageId      uint    `v:"required|integer#请选择绩效对应项目阀点|请选择正确的项目阀点" json:"stageId"`            // 项目所处阶段（cqgf_product_stage_rule.id）
	StageScore   uint    `v:"required|integer|min:0#请输入阶段得分阀点|得分为正整数|得分不能小于0" json:"stageScore"` // 阶段得分
	PmRadio      float64 `v:"required|min:0#请输入项目经理分配比例|项目经理分配比例不能小于0" json:"pmRadio"`           // PM分配比例
	PmFloatRadio float64 `v:"required|min:0#请输入项目经理浮动比例|项目经理浮动比例不能小于0" json:"pmFloatRadio"`      // PM浮动比例
	PmKpiLevelId uint    `v:"required|min:0#请选择项目经理绩效等级|项目经理绩效等级不能小于0" json:"pmKpiLevelId"`      // PM绩效等级
	Remark       string  `json:"remark"`                                                         // 预留备注信息
}

type ProductMemberPrizeComputeReq struct {
	ID      uint `json:"id"               `                                          //
	ProId   uint `v:"required|integer|min:1#请选择绩效对应项目|当前项目不存在|当前项目不存在" json:"proId"` // 项目主表ID
	StageId uint `v:"required|integer#请选择项目阀点|请选择正确的项目阀点" json:"stageId"`            // 项目所处阶段（cqgf_product_stage_rule.id）
}

// ProductMemberPrizeApiGetListReq 项目绩效激励清单
type ProductMemberPrizeApiGetListReq struct {
	ProductMemberPrize
	Page int32 `json:"page"` // 员工姓名
	Size int32 `json:"size"` // 员工姓名
}

// ProductMemberPrizeApiGetListRes 项目绩效激励清单
type ProductMemberPrizeApiGetListRes struct {
	ProductMemberPrize ProductMemberPrize `json:"productMemberPrize"` // 奖金分配
	ProductMemberKpi   ProductMemberKpi   `json:"productMemberKpi"`
	ProductMember      *ProductMember     `json:"productMember"`
	DepartmentName     string             `json:"departmentName"`
	UserName           string             `json:"userName"`
}

//type ProductMemberKpiInfo struct {
//	ProductMemberKpi *ProductMemberKpi `json:"productMemberKpi"`
//	ProductMember    *ProductMember    `json:"productMember"`
//	DepartmentName   string            `json:"departmentName"`
//	UserName         string            `json:"userName"`
//}
