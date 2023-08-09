// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrewKpiRuleDao is the data access object for table cqgf_crew_kpi_rule.
type CrewKpiRuleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CrewKpiRuleColumns // columns contains all the column names of Table for convenient usage.
}

// CrewKpiRuleColumns defines and stores column names for table cqgf_crew_kpi_rule.
type CrewKpiRuleColumns struct {
	Id         string //
	Redio      string // 比例
	ScoreMin   string // 得分下限
	ScoreMax   string // 得分上线
	ScoreRange string // 得分区间包含关系（1：左闭右开，2：左开右闭）
	LevelName  string // 等级名称
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// crewKpiRuleColumns holds the columns for table cqgf_crew_kpi_rule.
var crewKpiRuleColumns = CrewKpiRuleColumns{
	Id:         "id",
	Redio:      "redio",
	ScoreMin:   "score_min",
	ScoreMax:   "score_max",
	ScoreRange: "score_range",
	LevelName:  "level_name",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewCrewKpiRuleDao creates and returns a new DAO object for table data access.
func NewCrewKpiRuleDao() *CrewKpiRuleDao {
	return &CrewKpiRuleDao{
		group:   "default",
		table:   "cqgf_crew_kpi_rule",
		columns: crewKpiRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrewKpiRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrewKpiRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrewKpiRuleDao) Columns() CrewKpiRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrewKpiRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrewKpiRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrewKpiRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
