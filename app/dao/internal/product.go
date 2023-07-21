// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductDao is the data access object for table cqgf_product.
type ProductDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ProductColumns // columns contains all the column names of Table for convenient usage.
}

// ProductColumns defines and stores column names for table cqgf_product.
type ProductColumns struct {
	Id              string //
	Tid             string // 项目类型（type关联表）
	Name            string // 项目名称
	SubName         string // 项目简称
	LcScore         string // 优先级评分
	LccId           string // 项目优先级ID
	LccName         string // 项目优先级
	Invest          string // 投资额度
	NetProfit       string // 首年净利润
	ModeId          string // 研发模式ID（product_mode）
	PmId            string // PM(员工信息表ID)
	Attribute       string // PM属性（1：全职，2：兼职）
	PmlId           string // 项目责任人ID（员工信息表ID）
	IncentiveBudget string // 项目激励预算
	FixBudget       string //
	FixType         string // 修正预算计算类型（1：增加，2减少）
	Status          string // 项目当前状态(1:未开始 2：未立项，3：进行中 4：暂停 5：已取消 6：延迟 7：异常 8：已完成未验收 9：客户已验收 10：结项)
	ProTypeStageId  string // 项目当前所处阶段
	Remark          string // 预留备注说明信息
	CreateTime      string // 新增数据时间
	UpdateTime      string // 最后一次更新数据时间
}

// productColumns holds the columns for table cqgf_product.
var productColumns = ProductColumns{
	Id:              "id",
	Tid:             "tid",
	Name:            "name",
	SubName:         "sub_name",
	LcScore:         "lc_score",
	LccId:           "lcc_id",
	LccName:         "lcc_name",
	Invest:          "invest",
	NetProfit:       "net_profit",
	ModeId:          "mode_id",
	PmId:            "pm_id",
	Attribute:       "attribute",
	PmlId:           "pml_id",
	IncentiveBudget: "incentive_budget",
	FixBudget:       "fix_budget",
	FixType:         "fix_type",
	Status:          "status",
	ProTypeStageId:  "pro_type_stage_id",
	Remark:          "remark",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewProductDao creates and returns a new DAO object for table data access.
func NewProductDao() *ProductDao {
	return &ProductDao{
		group:   "default",
		table:   "cqgf_product",
		columns: productColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductDao) Columns() ProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
