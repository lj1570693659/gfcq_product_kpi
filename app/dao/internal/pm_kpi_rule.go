// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmKpiRuleDao is the data access object for table cqgf_pm_kpi_rule.
type PmKpiRuleDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PmKpiRuleColumns // columns contains all the column names of Table for convenient usage.
}

// PmKpiRuleColumns defines and stores column names for table cqgf_pm_kpi_rule.
type PmKpiRuleColumns struct {
	Id         string //
	Redio      string // 比例
	LevelName  string // 等级名称
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// pmKpiRuleColumns holds the columns for table cqgf_pm_kpi_rule.
var pmKpiRuleColumns = PmKpiRuleColumns{
	Id:         "id",
	Redio:      "redio",
	LevelName:  "level_name",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewPmKpiRuleDao creates and returns a new DAO object for table data access.
func NewPmKpiRuleDao() *PmKpiRuleDao {
	return &PmKpiRuleDao{
		group:   "default",
		table:   "cqgf_pm_kpi_rule",
		columns: pmKpiRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmKpiRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmKpiRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmKpiRuleDao) Columns() PmKpiRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmKpiRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmKpiRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmKpiRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
