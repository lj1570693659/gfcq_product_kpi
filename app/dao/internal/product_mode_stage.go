// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductModeStageDao is the data access object for table cqgf_product_mode_stage.
type ProductModeStageDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ProductModeStageColumns // columns contains all the column names of Table for convenient usage.
}

// ProductModeStageColumns defines and stores column names for table cqgf_product_mode_stage.
type ProductModeStageColumns struct {
	Id         string //
	Tid        string // 项目开发模式ID
	Name       string // 项目阶段名称
	QuotaRadio string // 阶段额度占比
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productModeStageColumns holds the columns for table cqgf_product_mode_stage.
var productModeStageColumns = ProductModeStageColumns{
	Id:         "id",
	Tid:        "tid",
	Name:       "name",
	QuotaRadio: "quota_radio",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductModeStageDao creates and returns a new DAO object for table data access.
func NewProductModeStageDao() *ProductModeStageDao {
	return &ProductModeStageDao{
		group:   "default",
		table:   "cqgf_product_mode_stage",
		columns: productModeStageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductModeStageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductModeStageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductModeStageDao) Columns() ProductModeStageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductModeStageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductModeStageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductModeStageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
