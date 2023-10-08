// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysJobLogDao is the data access object for table sys_job_log.
type SysJobLogDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysJobLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysJobLogColumns defines and stores column names for table sys_job_log.
type SysJobLogColumns struct {
	JobLogId      string // 任务日志ID
	JobName       string // 任务名称
	JobGroup      string // 任务组名
	InvokeTarget  string // 调用目标字符串
	JobMessage    string // 日志信息
	Status        string // 执行状态（0正常 1失败）
	ExceptionInfo string // 异常信息
	CreateTime    string // 创建时间
}

// sysJobLogColumns holds the columns for table sys_job_log.
var sysJobLogColumns = SysJobLogColumns{
	JobLogId:      "job_log_id",
	JobName:       "job_name",
	JobGroup:      "job_group",
	InvokeTarget:  "invoke_target",
	JobMessage:    "job_message",
	Status:        "status",
	ExceptionInfo: "exception_info",
	CreateTime:    "create_time",
}

// NewSysJobLogDao creates and returns a new DAO object for table data access.
func NewSysJobLogDao() *SysJobLogDao {
	return &SysJobLogDao{
		group:   "default",
		table:   "sys_job_log",
		columns: sysJobLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysJobLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysJobLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysJobLogDao) Columns() SysJobLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysJobLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysJobLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysJobLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
