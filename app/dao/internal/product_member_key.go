// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductMemberKeyDao is the data access object for table cqgf_product_member_key.
type ProductMemberKeyDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ProductMemberKeyColumns // columns contains all the column names of Table for convenient usage.
}

// ProductMemberKeyColumns defines and stores column names for table cqgf_product_member_key.
type ProductMemberKeyColumns struct {
	Id         string //
	StageKpiId string // 项目绩效ID
	ProId      string // 项目ID
	ProEmpId   string // 小组成员ID
	ProStageId string // 项目-阶段ID
	WorkNumber string // 成员工号
	Username   string // 成员姓名
	KeyName    string // 关键事件名称
	HappenTime string // 发生时间
	Type       string // 主体分类（1：加班贡献 2：解决问题贡献 3：其他事件贡献）
	Property   string // 事件性质（1：正向激励 2：有待提高）
	Result     string // 当前关键事件的处理结果
	Remark     string // 预留备注说明信息
	CreateTime string // 新增数据时间
	UpdateTime string // 最后一次更新数据时间
}

// productMemberKeyColumns holds the columns for table cqgf_product_member_key.
var productMemberKeyColumns = ProductMemberKeyColumns{
	Id:         "id",
	StageKpiId: "stage_kpi_id",
	ProId:      "pro_id",
	ProEmpId:   "pro_emp_id",
	ProStageId: "pro_stage_id",
	WorkNumber: "work_number",
	Username:   "username",
	KeyName:    "key_name",
	HappenTime: "happen_time",
	Type:       "type",
	Property:   "property",
	Result:     "result",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewProductMemberKeyDao creates and returns a new DAO object for table data access.
func NewProductMemberKeyDao() *ProductMemberKeyDao {
	return &ProductMemberKeyDao{
		group:   "default",
		table:   "cqgf_product_member_key",
		columns: productMemberKeyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductMemberKeyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductMemberKeyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductMemberKeyDao) Columns() ProductMemberKeyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductMemberKeyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductMemberKeyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductMemberKeyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
