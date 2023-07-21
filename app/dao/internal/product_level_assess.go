// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductLevelAssessDao is the data access object for table cqgf_product_level_assess.
type ProductLevelAssessDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ProductLevelAssessColumns // columns contains all the column names of Table for convenient usage.
}

// ProductLevelAssessColumns defines and stores column names for table cqgf_product_level_assess.
type ProductLevelAssessColumns struct {
	Id                 string //
	EvaluateDimensions string // 评价维度
	EvaluateCriteria   string // 评价标准
	ScoreCriteria      string // 评分标准
	EvaluateId         string // 上级评价维度
	Weight             string // 权重
	Remark             string // 预留备注说明信息
	CreateTime         string // 新增数据时间
	UpdateTime         string // 最新更新数据
}

// productLevelAssessColumns holds the columns for table cqgf_product_level_assess.
var productLevelAssessColumns = ProductLevelAssessColumns{
	Id:                 "id",
	EvaluateDimensions: "evaluate_dimensions",
	EvaluateCriteria:   "evaluate_criteria",
	ScoreCriteria:      "score_criteria",
	EvaluateId:         "evaluate_id",
	Weight:             "weight",
	Remark:             "remark",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewProductLevelAssessDao creates and returns a new DAO object for table data access.
func NewProductLevelAssessDao() *ProductLevelAssessDao {
	return &ProductLevelAssessDao{
		group:   "default",
		table:   "cqgf_product_level_assess",
		columns: productLevelAssessColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductLevelAssessDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductLevelAssessDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductLevelAssessDao) Columns() ProductLevelAssessColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductLevelAssessDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductLevelAssessDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductLevelAssessDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
