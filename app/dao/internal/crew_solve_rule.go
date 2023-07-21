// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrewSolveRuleDao is the data access object for table cqgf_crew_solve_rule.
type CrewSolveRuleDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns CrewSolveRuleColumns // columns contains all the column names of Table for convenient usage.
}

// CrewSolveRuleColumns defines and stores column names for table cqgf_crew_solve_rule.
type CrewSolveRuleColumns struct {
	Id         string //
	Redio      string // 浮动比例
	Demand     string // 贡献标准（1：有突出贡献，2：有重要贡献，3：较低/无贡献）
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// crewSolveRuleColumns holds the columns for table cqgf_crew_solve_rule.
var crewSolveRuleColumns = CrewSolveRuleColumns{
	Id:         "id",
	Redio:      "redio",
	Demand:     "demand",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewCrewSolveRuleDao creates and returns a new DAO object for table data access.
func NewCrewSolveRuleDao() *CrewSolveRuleDao {
	return &CrewSolveRuleDao{
		group:   "default",
		table:   "cqgf_crew_solve_rule",
		columns: crewSolveRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrewSolveRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrewSolveRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrewSolveRuleDao) Columns() CrewSolveRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrewSolveRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrewSolveRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrewSolveRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
