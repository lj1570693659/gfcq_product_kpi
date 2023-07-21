// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrewDutyIndexDao is the data access object for table cqgf_crew_duty_index.
type CrewDutyIndexDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns CrewDutyIndexColumns // columns contains all the column names of Table for convenient usage.
}

// CrewDutyIndexColumns defines and stores column names for table cqgf_crew_duty_index.
type CrewDutyIndexColumns struct {
	Id         string //
	ScoreIndex string // 职责指数
	JobLevelId string // 职级ID
	Arith      string // 运算方式
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// crewDutyIndexColumns holds the columns for table cqgf_crew_duty_index.
var crewDutyIndexColumns = CrewDutyIndexColumns{
	Id:         "id",
	ScoreIndex: "score_index",
	JobLevelId: "job_level_id",
	Arith:      "arith",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewCrewDutyIndexDao creates and returns a new DAO object for table data access.
func NewCrewDutyIndexDao() *CrewDutyIndexDao {
	return &CrewDutyIndexDao{
		group:   "default",
		table:   "cqgf_crew_duty_index",
		columns: crewDutyIndexColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrewDutyIndexDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrewDutyIndexDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrewDutyIndexDao) Columns() CrewDutyIndexColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrewDutyIndexDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrewDutyIndexDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrewDutyIndexDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
