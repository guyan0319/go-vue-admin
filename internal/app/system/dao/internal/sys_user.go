// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for table sys_user.
type SysUserColumns struct {
	UserId      string // 用户ID
	DeptId      string // 部门ID
	UserName    string // 用户账号
	NickName    string // 用户昵称
	UserType    string // 用户类型（00系统用户）
	Email       string // 用户邮箱
	Phonenumber string // 手机号码
	Sex         string // 用户性别（0男 1女 2未知）
	Avatar      string // 头像地址
	Password    string // 密码
	Status      string // 帐号状态（0正常 1停用）
	DelFlag     string // 删除标志（0代表存在 2代表删除）
	LoginIp     string // 最后登录IP
	LoginDate   string // 最后登录时间
	CreateBy    string // 创建者
	CreateTime  string // 创建时间
	UpdateBy    string // 更新者
	UpdateTime  string // 更新时间
	Remark      string // 备注
}

// sysUserColumns holds the columns for table sys_user.
var sysUserColumns = SysUserColumns{
	UserId:      "user_id",
	DeptId:      "dept_id",
	UserName:    "user_name",
	NickName:    "nick_name",
	UserType:    "user_type",
	Email:       "email",
	Phonenumber: "phonenumber",
	Sex:         "sex",
	Avatar:      "avatar",
	Password:    "password",
	Status:      "status",
	DelFlag:     "del_flag",
	LoginIp:     "login_ip",
	LoginDate:   "login_date",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
