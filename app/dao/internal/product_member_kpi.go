// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductMemberKpiDao is the data access object for table cqgf_product_member_kpi.
type ProductMemberKpiDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ProductMemberKpiColumns // columns contains all the column names of Table for convenient usage.
}

// ProductMemberKpiColumns defines and stores column names for table cqgf_product_member_kpi.
type ProductMemberKpiColumns struct {
	Id            string //
	IsPm          string // 是否是PM（1：是 2：否）
	ProId         string // 项目ID
	ProEmpId      string // 小组成员ID
	ProStageId    string // 项目-阶段ID
	OvertimeRadio string // 工时占比
	PrId          string // 项目角色ID
	PrName        string // 项目角色名称
	JbId          string // 职级ID
	JbName        string // 职级名称
	FloatRaio     string // 浮动贡献
	KpiLevelId    string // 绩效等级ID(区分是PM还是成员)
	KpiLevel      string // 绩效等级
	KpiRadio      string // 绩效比例
	Remark        string // 预留备注说明信息
	CreateTime    string // 新增数据时间
	UpdateTime    string // 最后一次更新数据时间
}

// productMemberKpiColumns holds the columns for table cqgf_product_member_kpi.
var productMemberKpiColumns = ProductMemberKpiColumns{
	Id:            "id",
	IsPm:          "is_pm",
	ProId:         "pro_id",
	ProEmpId:      "pro_emp_id",
	ProStageId:    "pro_stage_id",
	OvertimeRadio: "overtime_radio",
	PrId:          "pr_id",
	PrName:        "pr_name",
	JbId:          "jb_id",
	JbName:        "jb_name",
	FloatRaio:     "float_raio",
	KpiLevelId:    "kpi_level_id",
	KpiLevel:      "kpi_level",
	KpiRadio:      "kpi_radio",
	Remark:        "remark",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewProductMemberKpiDao creates and returns a new DAO object for table data access.
func NewProductMemberKpiDao() *ProductMemberKpiDao {
	return &ProductMemberKpiDao{
		group:   "default",
		table:   "cqgf_product_member_kpi",
		columns: productMemberKpiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductMemberKpiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductMemberKpiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductMemberKpiDao) Columns() ProductMemberKpiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductMemberKpiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductMemberKpiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductMemberKpiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
