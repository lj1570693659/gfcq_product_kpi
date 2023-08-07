package consts

const (
	// FullTime 全职
	FullTime = 1
	// PartTime 兼职
	PartTime = 2

	// SexUnknown 未知-用户未选择
	SexUnknown = 0
	// SexMan 男性
	SexMan = 1
	// SexWoMan 女性
	SexWoMan = 2

	// ScoreRangeMin 左闭右开
	ScoreRangeMin = 1
	// ScoreRangeMax 左开右闭
	ScoreRangeMax = 2
	// ScoreRangeMinAndMax 左闭右闭
	ScoreRangeMinAndMax = 3

	// ProductBudgetByScoreType [项目预算在预算区间中取值方式（1：取最小 2：取最大）] config配置信息
	ProductBudgetByScoreType = "product_budget_by_score_type"
	// ProductBudgetByMin 取最小
	ProductBudgetByMin = 1
	// ProductBudgetByMax 取最大
	ProductBudgetByMax = 2

	// BudgetNpRadio 截取净利润比例
	BudgetNpRadio = "budget_np_radio"

	// BudgetFixAdd 增加修正预算
	BudgetFixAdd = 1
	// BudgetFixLess 减少修正预算
	BudgetFixLess = 2

	// InspiritConfigDimension 绩效计算使用配置信息维度（crew:团队成员和PM共用成员配置信息，pm:团队成员和PM共用pm配置信息, single: 团队成员和PM用各自的配置信息）
	InspiritConfigDimension = "inspirit_config_dimension"

	// BusinessSupportRadio 业务支持部门激励比例
	BusinessSupportRadio = "business_support_radio"
	// TeamBuildingRadio 团建比例
	TeamBuildingRadio = "team_building_radio"
	// TeamRadio 项目团队比例
	TeamRadio = "team_radio"

	// HoursIndexRadio 基准指数中工时指数占比
	HoursIndexRadio = "hours_index_radio"
	// ManageIndexRadio 基准指数中管理指数占比
	ManageIndexRadio = "manage_index_radio"
	// DutyIndexRadio 基准指数中责任指数占比
	DutyIndexRadio = "duty_index_radio"
	// StageTopNumber 首页阀点绩效TOP排名
	StageTopNumber = "stage_top_number"
	// IsPm 项目经理对应数据
	IsPm = 1
	// IsNotPm 项目组成员对应数据
	IsNotPm = 2

	// OverTimeDevote 加班贡献
	OverTimeDevote = 1
	// SolveProblemDevote 解决问题贡献
	SolveProblemDevote = 2
	// ElseDevote 其他事件贡献
	ElseDevote = 3

	// ForwardDirection 1：正向激励
	ForwardDirection = 1
	// ReverseDirection 2：有待提高
	ReverseDirection = 2

	// MethodGET GET
	MethodGET = 1
	// MethodPOST POST
	MethodPOST = 2
	// MethodPUT PUT
	MethodPUT = 3
	// MethodDELETE DELETE
	MethodDELETE = 4
	// TokenName token值
	TokenName = "gfcq_token"
)
