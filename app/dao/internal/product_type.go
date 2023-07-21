// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductTypeDao is the data access object for table cqgf_product_type.
type ProductTypeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProductTypeColumns // columns contains all the column names of Table for convenient usage.
}

// ProductTypeColumns defines and stores column names for table cqgf_product_type.
type ProductTypeColumns struct {
	Id         string //
	Name       string // 开发模型名称
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productTypeColumns holds the columns for table cqgf_product_type.
var productTypeColumns = ProductTypeColumns{
	Id:         "id",
	Name:       "name",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductTypeDao creates and returns a new DAO object for table data access.
func NewProductTypeDao() *ProductTypeDao {
	return &ProductTypeDao{
		group:   "default",
		table:   "cqgf_product_type",
		columns: productTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductTypeDao) Columns() ProductTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
