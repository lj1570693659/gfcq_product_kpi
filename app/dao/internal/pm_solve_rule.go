// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmSolveRuleDao is the data access object for table cqgf_pm_solve_rule.
type PmSolveRuleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PmSolveRuleColumns // columns contains all the column names of Table for convenient usage.
}

// PmSolveRuleColumns defines and stores column names for table cqgf_pm_solve_rule.
type PmSolveRuleColumns struct {
	Id         string //
	Redio      string // 浮动比例
	Demand     string // 贡献标准（1：有突出贡献，2：有重要贡献，3：较低/无贡献）
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// pmSolveRuleColumns holds the columns for table cqgf_pm_solve_rule.
var pmSolveRuleColumns = PmSolveRuleColumns{
	Id:         "id",
	Redio:      "redio",
	Demand:     "demand",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewPmSolveRuleDao creates and returns a new DAO object for table data access.
func NewPmSolveRuleDao() *PmSolveRuleDao {
	return &PmSolveRuleDao{
		group:   "default",
		table:   "cqgf_pm_solve_rule",
		columns: pmSolveRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmSolveRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmSolveRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmSolveRuleDao) Columns() PmSolveRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmSolveRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmSolveRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmSolveRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
