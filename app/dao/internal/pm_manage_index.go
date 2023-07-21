// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmManageIndexDao is the data access object for table cqgf_pm_manage_index.
type PmManageIndexDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns PmManageIndexColumns // columns contains all the column names of Table for convenient usage.
}

// PmManageIndexColumns defines and stores column names for table cqgf_pm_manage_index.
type PmManageIndexColumns struct {
	Id            string //
	ScoreIndex    string // 管理指数
	ProductRoleId string // 项目角色ID
	Remark        string // 预留备注说明信息
	CreateTime    string // 新增数据时间
	UpdateTime    string // 最后一次更新数据时间
}

// pmManageIndexColumns holds the columns for table cqgf_pm_manage_index.
var pmManageIndexColumns = PmManageIndexColumns{
	Id:            "id",
	ScoreIndex:    "score_index",
	ProductRoleId: "product_role_id",
	Remark:        "remark",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewPmManageIndexDao creates and returns a new DAO object for table data access.
func NewPmManageIndexDao() *PmManageIndexDao {
	return &PmManageIndexDao{
		group:   "default",
		table:   "cqgf_pm_manage_index",
		columns: pmManageIndexColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmManageIndexDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmManageIndexDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmManageIndexDao) Columns() PmManageIndexColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmManageIndexDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmManageIndexDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmManageIndexDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
