// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductModeDao is the data access object for table cqgf_product_mode.
type ProductModeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProductModeColumns // columns contains all the column names of Table for convenient usage.
}

// ProductModeColumns defines and stores column names for table cqgf_product_mode.
type ProductModeColumns struct {
	Id         string //
	Name       string // 开发模型名称
	Factor     string // 开发系数
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productModeColumns holds the columns for table cqgf_product_mode.
var productModeColumns = ProductModeColumns{
	Id:         "id",
	Name:       "name",
	Factor:     "factor",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductModeDao creates and returns a new DAO object for table data access.
func NewProductModeDao() *ProductModeDao {
	return &ProductModeDao{
		group:   "default",
		table:   "cqgf_product_mode",
		columns: productModeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductModeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductModeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductModeDao) Columns() ProductModeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductModeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductModeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductModeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
