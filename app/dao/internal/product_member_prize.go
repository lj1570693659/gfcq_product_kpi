// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductMemberPrizeDao is the data access object for table cqgf_product_member_prize.
type ProductMemberPrizeDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ProductMemberPrizeColumns // columns contains all the column names of Table for convenient usage.
}

// ProductMemberPrizeColumns defines and stores column names for table cqgf_product_member_prize.
type ProductMemberPrizeColumns struct {
	Id              string //
	ProId           string // 项目ID
	IsPm            string // 是否是PM（1：是 2：否）
	ProEmpId        string // 小组成员ID
	ProStageId      string // 项目-阶段ID
	OvertimeRadio   string // 工时占比
	OvertimeIndex   string // 工时指数
	PrId            string // 项目角色ID
	PrName          string // 项目角色名称
	ManageIndex     string // 管理指数
	JbId            string // 职级ID
	JbName          string // 职级名称
	DutyIndex       string // 责任指数
	BaseIndex       string // 基准指数
	WeightAutoRadio string // 权重基准（自动）
	WeightPmoRadio  string // 权重基准（PMO）
	SentBase        string // 发放基数
	RemaindQueto    string // 剩余额度
	FloatRaio       string // 浮动贡献
	KpiLevelId      string // 绩效等级ID(区分是PM还是成员)
	KpiLevel        string // 绩效等级
	KpiRadio        string // 绩效比例
	SentQueto       string // 实发额度
	Remark          string // 预留备注说明信息
	CreateTime      string // 新增数据时间
	UpdateTime      string // 最后一次更新数据时间
}

// productMemberPrizeColumns holds the columns for table cqgf_product_member_prize.
var productMemberPrizeColumns = ProductMemberPrizeColumns{
	Id:              "id",
	ProId:           "pro_id",
	IsPm:            "is_pm",
	ProEmpId:        "pro_emp_id",
	ProStageId:      "pro_stage_id",
	OvertimeRadio:   "overtime_radio",
	OvertimeIndex:   "overtime_index",
	PrId:            "pr_id",
	PrName:          "pr_name",
	ManageIndex:     "manage_index",
	JbId:            "jb_id",
	JbName:          "jb_name",
	DutyIndex:       "duty_index",
	BaseIndex:       "base_index",
	WeightAutoRadio: "weight_auto_radio",
	WeightPmoRadio:  "weight_pmo_radio",
	SentBase:        "sent_base",
	RemaindQueto:    "remaind_queto",
	FloatRaio:       "float_raio",
	KpiLevelId:      "kpi_level_id",
	KpiLevel:        "kpi_level",
	KpiRadio:        "kpi_radio",
	SentQueto:       "sent_queto",
	Remark:          "remark",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewProductMemberPrizeDao creates and returns a new DAO object for table data access.
func NewProductMemberPrizeDao() *ProductMemberPrizeDao {
	return &ProductMemberPrizeDao{
		group:   "default",
		table:   "cqgf_product_member_prize",
		columns: productMemberPrizeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductMemberPrizeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductMemberPrizeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductMemberPrizeDao) Columns() ProductMemberPrizeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductMemberPrizeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductMemberPrizeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductMemberPrizeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
