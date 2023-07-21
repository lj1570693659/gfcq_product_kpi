// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmHoursIndexDao is the data access object for table cqgf_pm_hours_index.
type PmHoursIndexDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns PmHoursIndexColumns // columns contains all the column names of Table for convenient usage.
}

// PmHoursIndexColumns defines and stores column names for table cqgf_pm_hours_index.
type PmHoursIndexColumns struct {
	Id         string //
	ScoreIndex string // 工时指数
	ScoreMin   string // 得分下限
	ScoreMax   string // 得分上线
	ScoreRange string // 得分区间包含关系（1：左闭右开，2：左开右闭）
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// pmHoursIndexColumns holds the columns for table cqgf_pm_hours_index.
var pmHoursIndexColumns = PmHoursIndexColumns{
	Id:         "id",
	ScoreIndex: "score_index",
	ScoreMin:   "score_min",
	ScoreMax:   "score_max",
	ScoreRange: "score_range",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewPmHoursIndexDao creates and returns a new DAO object for table data access.
func NewPmHoursIndexDao() *PmHoursIndexDao {
	return &PmHoursIndexDao{
		group:   "default",
		table:   "cqgf_pm_hours_index",
		columns: pmHoursIndexColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmHoursIndexDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmHoursIndexDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmHoursIndexDao) Columns() PmHoursIndexColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmHoursIndexDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmHoursIndexDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmHoursIndexDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
