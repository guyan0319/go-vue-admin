// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDataDao is the data access object for table sys_dict_data.
type SysDictDataDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysDictDataColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictDataColumns defines and stores column names for table sys_dict_data.
type SysDictDataColumns struct {
	DictCode   string // 字典编码
	DictSort   string // 字典排序
	DictLabel  string // 字典标签
	DictValue  string // 字典键值
	DictType   string // 字典类型
	CssClass   string // 样式属性（其他样式扩展）
	ListClass  string // 表格回显样式
	IsDefault  string // 是否默认（Y是 N否）
	Status     string // 状态（0正常 1停用）
	CreateBy   string // 创建者
	CreateTime string // 创建时间
	UpdateBy   string // 更新者
	UpdateTime string // 更新时间
	Remark     string // 备注
}

// sysDictDataColumns holds the columns for table sys_dict_data.
var sysDictDataColumns = SysDictDataColumns{
	DictCode:   "dict_code",
	DictSort:   "dict_sort",
	DictLabel:  "dict_label",
	DictValue:  "dict_value",
	DictType:   "dict_type",
	CssClass:   "css_class",
	ListClass:  "list_class",
	IsDefault:  "is_default",
	Status:     "status",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewSysDictDataDao creates and returns a new DAO object for table data access.
func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		group:   "default",
		table:   "sys_dict_data",
		columns: sysDictDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDictDataDao) Columns() SysDictDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDictDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
