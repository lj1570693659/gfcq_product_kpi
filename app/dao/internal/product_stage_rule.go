// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductStageRuleDao is the data access object for table cqgf_product_stage_rule.
type ProductStageRuleDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ProductStageRuleColumns // columns contains all the column names of Table for convenient usage.
}

// ProductStageRuleColumns defines and stores column names for table cqgf_product_stage_rule.
type ProductStageRuleColumns struct {
	Id         string //
	Name       string // 阶段名称
	ProId      string // 项目ID
	ProStageId string // 项目-阶段ID
	QuotaRadio string // 阶段额度占比
	Remark     string // 预留备注信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productStageRuleColumns holds the columns for table cqgf_product_stage_rule.
var productStageRuleColumns = ProductStageRuleColumns{
	Id:         "id",
	Name:       "name",
	ProId:      "pro_id",
	ProStageId: "pro_stage_id",
	QuotaRadio: "quota_radio",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductStageRuleDao creates and returns a new DAO object for table data access.
func NewProductStageRuleDao() *ProductStageRuleDao {
	return &ProductStageRuleDao{
		group:   "default",
		table:   "cqgf_product_stage_rule",
		columns: productStageRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductStageRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductStageRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductStageRuleDao) Columns() ProductStageRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductStageRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductStageRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductStageRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
