// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type SysRoleColumns struct {
	RoleId            string // 角色ID
	RoleName          string // 角色名称
	RoleKey           string // 角色权限字符串
	RoleSort          string // 显示顺序
	DataScope         string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly string // 菜单树选择项是否关联显示
	DeptCheckStrictly string // 部门树选择项是否关联显示
	Status            string // 角色状态（0正常 1停用）
	DelFlag           string // 删除标志（0代表存在 2代表删除）
	CreateBy          string // 创建者
	CreateTime        string // 创建时间
	UpdateBy          string // 更新者
	UpdateTime        string // 更新时间
	Remark            string // 备注
}

// sysRoleColumns holds the columns for table sys_role.
var sysRoleColumns = SysRoleColumns{
	RoleId:            "role_id",
	RoleName:          "role_name",
	RoleKey:           "role_key",
	RoleSort:          "role_sort",
	DataScope:         "data_scope",
	MenuCheckStrictly: "menu_check_strictly",
	DeptCheckStrictly: "dept_check_strictly",
	Status:            "status",
	DelFlag:           "del_flag",
	CreateBy:          "create_by",
	CreateTime:        "create_time",
	UpdateBy:          "update_by",
	UpdateTime:        "update_time",
	Remark:            "remark",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		group:   "default",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
