// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLogDao is the data access object for table cqgf_user_log.
type UserLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns UserLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserLogColumns defines and stores column names for table cqgf_user_log.
type UserLogColumns struct {
	Id                  string //
	MethodName          string // 对应HTTP请求中Method参数值(1: GET 2: POST 3:PUT 4:DELETE)
	RequestUri          string // http请求接口地址
	WorkNumber          string // 唯一账号（员工工号）
	RequestModule       string // 请求模块（一级模块）
	RequestSecondModule string // 请求模块（二级模块）
	ChangeTypeName      string // 操作对象全称
	RequestBody         string // 具体请求参数
	Remark              string // 预留补充说明信息
	CreateTime          string // 数据新增时间
	UpdateTime          string // 最后一次更新数据时间
}

// userLogColumns holds the columns for table cqgf_user_log.
var userLogColumns = UserLogColumns{
	Id:                  "id",
	MethodName:          "method_name",
	RequestUri:          "request_uri",
	WorkNumber:          "work_number",
	RequestModule:       "request_module",
	RequestSecondModule: "request_second_module",
	ChangeTypeName:      "change_type_name",
	RequestBody:         "request_body",
	Remark:              "remark",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewUserLogDao creates and returns a new DAO object for table data access.
func NewUserLogDao() *UserLogDao {
	return &UserLogDao{
		group:   "default",
		table:   "cqgf_user_log",
		columns: userLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLogDao) Columns() UserLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
