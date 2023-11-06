package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/swagger"
	_ "github.com/lj1570693659/gfcq_product_kpi/packed"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	product "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	wechat "github.com/lj1570693659/gfcq_protoc/wechat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// DepertmentServer 部门服务
	DepertmentServer v1.DepartmentClient
	// JobServer 岗位服务
	JobServer v1.JobClient
	// EmployeeServer 员工服务
	EmployeeServer v1.EmployeeClient
	// JobLevelServer 职级服务
	JobLevelServer v1.JobLevelClient
	// EmployeeJobServer 员工&&岗位关联服务
	EmployeeJobServer v1.EmployeeJobClient
	// WechatCheckIn 员工打卡服务
	WechatCheckIn wechat.WechatCheckInClient

	// LevelAssessServer /* ------------------------------------ */
	// LevelAssessServer 项目等级评估配置信息
	LevelAssessServer product.LevelAssessClient
	// LevelConfirmServer 项目优先级确认配置信息
	LevelConfirmServer product.LevelConfirmClient
	// ModeServer 项目开发模式
	ModeServer product.ModeClient
	// ModeStageServer 项目开发模式-包含阶段关联信息
	ModeStageServer product.ModeStageClient
	// RolesServer 项目角色配置信息
	RolesServer product.RolesClient
	// TypeServer 项目类型信息
	TypeServer product.TypeClient

	// BudgetAssessServer 激励预算
	BudgetAssessServer inspirit.BudgetAssessClient
	// CrewDutyIndexServer 员工责任指数
	CrewDutyIndexServer inspirit.CrewDutyIndexClient
	// CrewHoursIndexServer 员工工时指数
	CrewHoursIndexServer inspirit.CrewHoursIndexClient
	// CrewManageIndexServer 员工管理指数
	CrewManageIndexServer inspirit.CrewManageIndexClient
	// CrewOvertimeRuleServer 项目组成员浮动贡献-加班贡献配置信息
	CrewOvertimeRuleServer inspirit.CrewOvertimeRuleClient
	// CrewSolveRuleServer 项目组成员浮动贡献-解决问题贡献配置信息
	CrewSolveRuleServer inspirit.CrewSolveRuleClient
	// CrewKpiRuleServer 团队成员绩效等级配置信息
	CrewKpiRuleServer inspirit.CrewKpiRuleClient
	// StageRadioServer 项目阶段应发激励占比配置信息
	StageRadioServer inspirit.StageRadioClient
)

// 用于应用初始化。
func init() {
	// 部门、员工基础信息服务
	organizeServerName := g.Config("config.toml").Get("grpc.organize.link")
	OrganizeServer, err := grpc.Dial(gconv.String(organizeServerName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	DepertmentServer = v1.NewDepartmentClient(OrganizeServer)
	EmployeeServer = v1.NewEmployeeClient(OrganizeServer)
	JobServer = v1.NewJobClient(OrganizeServer)
	JobLevelServer = v1.NewJobLevelClient(OrganizeServer)
	EmployeeJobServer = v1.NewEmployeeJobClient(OrganizeServer)
	WechatCheckIn = wechat.NewWechatCheckInClient(OrganizeServer)

	// 公共配置服务
	configServerName := g.Config("config.toml").Get("grpc.config.link")
	ConfigServer, err := grpc.Dial(gconv.String(configServerName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	LevelAssessServer = product.NewLevelAssessClient(ConfigServer)
	LevelConfirmServer = product.NewLevelConfirmClient(ConfigServer)
	ModeServer = product.NewModeClient(ConfigServer)
	ModeStageServer = product.NewModeStageClient(ConfigServer)
	RolesServer = product.NewRolesClient(ConfigServer)
	TypeServer = product.NewTypeClient(ConfigServer)

	BudgetAssessServer = inspirit.NewBudgetAssessClient(ConfigServer)
	CrewDutyIndexServer = inspirit.NewCrewDutyIndexClient(ConfigServer)
	CrewHoursIndexServer = inspirit.NewCrewHoursIndexClient(ConfigServer)
	CrewManageIndexServer = inspirit.NewCrewManageIndexClient(ConfigServer)
	CrewOvertimeRuleServer = inspirit.NewCrewOvertimeRuleClient(ConfigServer)
	CrewSolveRuleServer = inspirit.NewCrewSolveRuleClient(ConfigServer)
	CrewKpiRuleServer = inspirit.NewCrewKpiRuleClient(ConfigServer)
	StageRadioServer = inspirit.NewStageRadioClient(ConfigServer)

	s := g.Server()
	s.SetFileServerEnabled(true)
	s.AddSearchPath("./public/excel")
	s.Plugin(&swagger.Swagger{})
}
