// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductTaskDao is the data access object for table cqgf_product_task.
type ProductTaskDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProductTaskColumns // columns contains all the column names of Table for convenient usage.
}

// ProductTaskColumns defines and stores column names for table cqgf_product_task.
type ProductTaskColumns struct {
	Id                  string //
	ProductId           string // 项目唯一标志
	TaskName            string // 任务名称
	TaskCate            string // 任务分类
	TaskDesc            string // 任务描述
	TaskStatus          string // 任务状态（1：未开启，2：运行中，3：已完成 4：暂停）
	Level               string // 优先级（值越大，优先级越高）
	IsSendGroupChat     string // 1: 发送至群 2：不发送至群
	GroupChat           string // 发送至群信息
	RemindGroupChatUser string // 发送到群时，提醒相关关注人
	DutyEmployeeId      string // 责任人员工信息
	DutyWorkNumber      string // 责任人员工工号
	JoinWorkNumber      string // 关联责任人
	PalnStartTime       string // 计划开始时间
	PalnEndTime         string // 计划结束时间
	RealStartTime       string // 实际开始时间
	RealEndTime         string // 实际结束时间
	UpgradeFirst        string // 第一次升级@工号
	UpgradeFirstTime    string // 第一次升级时间
	UpgradeTwo          string // 第二次升级@工号
	UpgradeTwoTime      string // 第三次升级@工号
	UpgradeThree        string // 第三次升级@工号
	UpgradeThreeTime    string // 第三次升级时间
	AssignEmployeeId    string // 指派人员工信息
	AssignWorkNumber    string // 指派人员工工号
	Remark              string // 预留备注说明信息
	CreateTime          string // 新增数据时间
	UpdateTime          string // 最后一次更新数据时间
}

// productTaskColumns holds the columns for table cqgf_product_task.
var productTaskColumns = ProductTaskColumns{
	Id:                  "id",
	ProductId:           "product_id",
	TaskName:            "task_name",
	TaskCate:            "task_cate",
	TaskDesc:            "task_desc",
	TaskStatus:          "task_status",
	Level:               "level",
	IsSendGroupChat:     "is_send_group_chat",
	GroupChat:           "group_chat",
	RemindGroupChatUser: "remind_group_chat_user",
	DutyEmployeeId:      "duty_employee_id",
	DutyWorkNumber:      "duty_work_number",
	JoinWorkNumber:      "join_work_number",
	PalnStartTime:       "paln_start_time",
	PalnEndTime:         "paln_end_time",
	RealStartTime:       "real_start_time",
	RealEndTime:         "real_end_time",
	UpgradeFirst:        "upgrade_first",
	UpgradeFirstTime:    "upgrade_first_time",
	UpgradeTwo:          "upgrade_two",
	UpgradeTwoTime:      "upgrade_two_time",
	UpgradeThree:        "upgrade_three",
	UpgradeThreeTime:    "upgrade_three_time",
	AssignEmployeeId:    "assign_employee_id",
	AssignWorkNumber:    "assign_work_number",
	Remark:              "remark",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewProductTaskDao creates and returns a new DAO object for table data access.
func NewProductTaskDao() *ProductTaskDao {
	return &ProductTaskDao{
		group:   "default",
		table:   "cqgf_product_task",
		columns: productTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductTaskDao) Columns() ProductTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
