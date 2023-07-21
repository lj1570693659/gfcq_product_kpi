// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrewOvertimeRuleDao is the data access object for table cqgf_crew_overtime_rule.
type CrewOvertimeRuleDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns CrewOvertimeRuleColumns // columns contains all the column names of Table for convenient usage.
}

// CrewOvertimeRuleColumns defines and stores column names for table cqgf_crew_overtime_rule.
type CrewOvertimeRuleColumns struct {
	Id         string //
	Redio      string // 浮动比例
	ScoreMin   string // 得分下限
	ScoreMax   string // 得分上线
	ScoreRange string // 得分区间包含关系（1：左闭右开，2：左开右闭）
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// crewOvertimeRuleColumns holds the columns for table cqgf_crew_overtime_rule.
var crewOvertimeRuleColumns = CrewOvertimeRuleColumns{
	Id:         "id",
	Redio:      "redio",
	ScoreMin:   "score_min",
	ScoreMax:   "score_max",
	ScoreRange: "score_range",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewCrewOvertimeRuleDao creates and returns a new DAO object for table data access.
func NewCrewOvertimeRuleDao() *CrewOvertimeRuleDao {
	return &CrewOvertimeRuleDao{
		group:   "default",
		table:   "cqgf_crew_overtime_rule",
		columns: crewOvertimeRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrewOvertimeRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrewOvertimeRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrewOvertimeRuleDao) Columns() CrewOvertimeRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrewOvertimeRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrewOvertimeRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrewOvertimeRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
