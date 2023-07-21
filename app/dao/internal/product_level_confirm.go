// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductLevelConfirmDao is the data access object for table cqgf_product_level_confirm.
type ProductLevelConfirmDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns ProductLevelConfirmColumns // columns contains all the column names of Table for convenient usage.
}

// ProductLevelConfirmColumns defines and stores column names for table cqgf_product_level_confirm.
type ProductLevelConfirmColumns struct {
	Id            string //
	Name          string // 项目优先级
	ScoreMin      string // 得分下限
	ScoreMax      string // 得分上线
	ScoreRange    string // 得分区间包含关系（1：左闭右开，2：左开右闭）
	IsNeedPm      string // 是否委派PM(1:是 2：否)
	PmDemand      string // pm要求
	ProductDemand string // 项目工作相关要求
	MonitDemand   string // 监控要求
	IsNeedPml     string // 是否需要项目负责人(1:是 2：否)
	Remark        string // 预留备注说明信息
	CreateTime    string // 新增数据时间
	UpdateTime    string // 最新更新数据
}

// productLevelConfirmColumns holds the columns for table cqgf_product_level_confirm.
var productLevelConfirmColumns = ProductLevelConfirmColumns{
	Id:            "id",
	Name:          "name",
	ScoreMin:      "score_min",
	ScoreMax:      "score_max",
	ScoreRange:    "score_range",
	IsNeedPm:      "is_need_pm",
	PmDemand:      "pm_demand",
	ProductDemand: "product_demand",
	MonitDemand:   "monit_demand",
	IsNeedPml:     "is_need_pml",
	Remark:        "remark",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewProductLevelConfirmDao creates and returns a new DAO object for table data access.
func NewProductLevelConfirmDao() *ProductLevelConfirmDao {
	return &ProductLevelConfirmDao{
		group:   "default",
		table:   "cqgf_product_level_confirm",
		columns: productLevelConfirmColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductLevelConfirmDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductLevelConfirmDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductLevelConfirmDao) Columns() ProductLevelConfirmColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductLevelConfirmDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductLevelConfirmDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductLevelConfirmDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
