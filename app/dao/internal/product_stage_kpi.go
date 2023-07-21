// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductStageKpiDao is the data access object for table cqgf_product_stage_kpi.
type ProductStageKpiDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ProductStageKpiColumns // columns contains all the column names of Table for convenient usage.
}

// ProductStageKpiColumns defines and stores column names for table cqgf_product_stage_kpi.
type ProductStageKpiColumns struct {
	Id               string //
	ProId            string // 项目主表ID
	StageId          string // 项目所处阶段（cqgf_product_stage_rule.id）
	StageRadio       string // 阶段比例
	StageBudget      string // 阶段预算
	StageScore       string // 阶段得分
	ShouldSentRadio  string // 应发比例
	StageQuota       string // 阶段额度
	CrewQuota        string // 团队额度
	TeamBuildQuota   string // 团建额度
	SupportQuota     string // 业务支持额度
	PmRadio          string // PM分配比例
	PmBase           string // PM发放基础
	PmFloatRadio     string // PM浮动比例
	PmKpiLevelId     string // PM绩效等级
	PmKpiLevelName   string // PM绩效等级名称
	PmKpiLevelRadio  string // PM绩效等级比例
	PmIncentiveQuota string // PM实际应发额度
	Remark           string // 预留备注信息
	CreateTime       string // 新增数据时间
	UpdateTime       string // 最后一次更新数据时间
}

// productStageKpiColumns holds the columns for table cqgf_product_stage_kpi.
var productStageKpiColumns = ProductStageKpiColumns{
	Id:               "id",
	ProId:            "pro_id",
	StageId:          "stage_id",
	StageRadio:       "stage_radio",
	StageBudget:      "stage_budget",
	StageScore:       "stage_score",
	ShouldSentRadio:  "should_sent_radio",
	StageQuota:       "stage_quota",
	CrewQuota:        "crew_quota",
	TeamBuildQuota:   "team_build_quota",
	SupportQuota:     "support_quota",
	PmRadio:          "pm_radio",
	PmBase:           "pm_base",
	PmFloatRadio:     "pm_float_radio",
	PmKpiLevelId:     "pm_kpi_level_id",
	PmKpiLevelName:   "pm_kpi_level_name",
	PmKpiLevelRadio:  "pm_kpi_level_radio",
	PmIncentiveQuota: "pm_incentive_quota",
	Remark:           "remark",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewProductStageKpiDao creates and returns a new DAO object for table data access.
func NewProductStageKpiDao() *ProductStageKpiDao {
	return &ProductStageKpiDao{
		group:   "default",
		table:   "cqgf_product_stage_kpi",
		columns: productStageKpiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductStageKpiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductStageKpiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductStageKpiDao) Columns() ProductStageKpiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductStageKpiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductStageKpiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductStageKpiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
