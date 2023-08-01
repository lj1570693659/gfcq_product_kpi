// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductRolesDao is the data access object for table cqgf_product_roles.
type ProductRolesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ProductRolesColumns // columns contains all the column names of Table for convenient usage.
}

// ProductRolesColumns defines and stores column names for table cqgf_product_roles.
type ProductRolesColumns struct {
	Id         string //
	Name       string //
	Pid        string // 上级角色
	Explain    string // 角色与职责说明
	IsSpecial  string // 1: 需要特殊处理 2：不需要特殊处理
	Remark     string // 预留备注信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productRolesColumns holds the columns for table cqgf_product_roles.
var productRolesColumns = ProductRolesColumns{
	Id:         "id",
	Name:       "name",
	Pid:        "pid",
	Explain:    "explain",
	IsSpecial:  "is_special",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductRolesDao creates and returns a new DAO object for table data access.
func NewProductRolesDao() *ProductRolesDao {
	return &ProductRolesDao{
		group:   "default",
		table:   "cqgf_product_roles",
		columns: productRolesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductRolesDao) Columns() ProductRolesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductRolesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
