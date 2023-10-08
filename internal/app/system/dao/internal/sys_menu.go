// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for table sys_menu.
type SysMenuColumns struct {
	MenuId     string // 菜单ID
	MenuName   string // 菜单名称
	ParentId   string // 父菜单ID
	OrderNum   string // 显示顺序
	Path       string // 路由地址
	ApiPath    string // 后台api路径
	Component  string // 组件路径
	Query      string // 路由参数
	IsFrame    string // 是否为外链（0是 1否）
	IsCache    string // 是否缓存（0缓存 1不缓存）
	MenuType   string // 菜单类型（M目录 C菜单 F按钮）
	Visible    string // 菜单状态（0显示 1隐藏）
	Status     string // 菜单状态（0正常 1停用）
	Perms      string // 权限标识
	Icon       string // 菜单图标
	CreateBy   string // 创建者
	CreateTime string // 创建时间
	UpdateBy   string // 更新者
	UpdateTime string // 更新时间
	Remark     string // 备注
}

// sysMenuColumns holds the columns for table sys_menu.
var sysMenuColumns = SysMenuColumns{
	MenuId:     "menu_id",
	MenuName:   "menu_name",
	ParentId:   "parent_id",
	OrderNum:   "order_num",
	Path:       "path",
	ApiPath:    "api_path",
	Component:  "component",
	Query:      "query",
	IsFrame:    "is_frame",
	IsCache:    "is_cache",
	MenuType:   "menu_type",
	Visible:    "visible",
	Status:     "status",
	Perms:      "perms",
	Icon:       "icon",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
