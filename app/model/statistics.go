package model

type Inspire struct {
	ProductCount    int     `json:"productCount"`    // 项目总数量
	IncentiveBudget float64 `json:"incentiveBudget"` // 激励预算
	StageQuota      float64 `json:"stageQuota"`      // 实发汇总
	StageBudget     float64 `json:"stageBudget"`     // 应发汇总
}

type StageStatic struct {
	StageName   string  `json:"stageName"`   // 阶段名称
	StageQuota  float64 `json:"stageQuota"`  // 实发汇总
	StageBudget float64 `json:"stageBudget"` // 应发汇总
}

// ProductStageScore 项目阶段得分折线图
type ProductStageScore struct {
	ProductSubName string `json:"productSubName"` // 阶段名称(项目简称/阀点)
	ProductName    string `json:"productName"`    // 阶段名称(项目简称/阀点)
	StageName      string `json:"stageName"`      // 阶段名称(项目简称/阀点)
	StageScore     uint   `json:"stageScore"`     // 阶段得分
}

// ProductStageLint 导入项目组成员数据结构
type ProductStageLint struct {
	StageName   []string  `json:"stageName"`
	StageQuota  []float64 `json:"stageQuota"`
	StageBudget []float64 `json:"stageBudget"`
}

// ProductStageTop 项目阶段得分折线图
type ProductStageTop struct {
	ProductSubName  string  `json:"productSubName"`  // 阶段名称(项目简称/阀点)
	ProductName     string  `json:"productName"`     // 阶段名称(项目简称/阀点)
	StageName       string  `json:"stageName"`       // 阶段名称(项目简称/阀点)
	TeamNumber      int     `json:"teamNumber"`      // 团队人数
	ShouldSentRadio float64 `json:"shouldSentRadio"` // 发放比例
	StageScore      uint    `json:"stageScore"`      // 阶段得分
	StageQuota      float64 `json:"stageQuota"`
	StageBudget     float64 `json:"stageBudget"`
}

// ProductMemberStaticWhere 项目阶段得分折线图
type ProductMemberStaticWhere struct {
	ProId      uint   `json:"proId"`      // 项目ID
	WorkNumber string `json:"workNumber"` // 项目名称
	EmpId      []uint `json:"empId"`      // 项目成员ID
	PrId       []uint `json:"prId"`       // 项目角色ID
	JbId       []uint `json:"jbId"`       // 职级ID
	Attribute  []uint `json:"attribute"`  // PM属性（1：全职，2：兼职）
}

// ProductMemberStatic 项目阶段得分折线图
type ProductMemberStatic struct {
	Title     []string         `json:"title"`     // 绩效等级
	Legend    []string         `json:"legend"`    // 项目名称
	YAxisData map[string][]int `json:"yAxisData"` // 项目名称
}
