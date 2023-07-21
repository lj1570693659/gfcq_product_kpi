// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductMemberDao is the data access object for table cqgf_product_member.
type ProductMemberDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ProductMemberColumns // columns contains all the column names of Table for convenient usage.
}

// ProductMemberColumns defines and stores column names for table cqgf_product_member.
type ProductMemberColumns struct {
	Id          string //
	ProId       string // 项目ID
	EmpId       string // 项目成员ID
	WorkNumber  string // 员工工号
	Attribute   string // 属性（1：全职，2：兼职）
	PrId        string // 项目角色ID
	PrName      string // 项目角色名称
	ManageIndex string // 管理指数
	JbId        string // 职级ID
	JbName      string // 职级名称
	DutyIndex   string // 责任指数
	Remark      string // 预留备注说明信息
	CreateTime  string // 新增数据时间
	UpdateTime  string // 最后一次更新数据时间
}

// productMemberColumns holds the columns for table cqgf_product_member.
var productMemberColumns = ProductMemberColumns{
	Id:          "id",
	ProId:       "pro_id",
	EmpId:       "emp_id",
	WorkNumber:  "work_number",
	Attribute:   "attribute",
	PrId:        "pr_id",
	PrName:      "pr_name",
	ManageIndex: "manage_index",
	JbId:        "jb_id",
	JbName:      "jb_name",
	DutyIndex:   "duty_index",
	Remark:      "remark",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewProductMemberDao creates and returns a new DAO object for table data access.
func NewProductMemberDao() *ProductMemberDao {
	return &ProductMemberDao{
		group:   "default",
		table:   "cqgf_product_member",
		columns: productMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductMemberDao) Columns() ProductMemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
