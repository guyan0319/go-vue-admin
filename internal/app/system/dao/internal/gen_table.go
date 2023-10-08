// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GenTableDao is the data access object for table gen_table.
type GenTableDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns GenTableColumns // columns contains all the column names of Table for convenient usage.
}

// GenTableColumns defines and stores column names for table gen_table.
type GenTableColumns struct {
	TableId        string // 编号
	TableName      string // 表名称
	TableComment   string // 表描述
	SubTableName   string // 关联子表的表名
	SubTableFkName string // 子表关联的外键名
	ClassName      string // 实体类名称
	TplCategory    string // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string // 生成包路径
	ModuleName     string // 生成模块名
	BusinessName   string // 生成业务名
	FunctionName   string // 生成功能名
	FunctionAuthor string // 生成功能作者
	GenType        string // 生成代码方式（0zip压缩包 1自定义路径）
	Status         string // 状态（0正常 1删除）
	GenPath        string // 生成路径（不填默认项目路径）
	Options        string // 其它生成选项
	CreateBy       string // 创建者
	CreateTime     string // 创建时间
	UpdateBy       string // 更新者
	UpdateTime     string // 更新时间
	Remark         string // 备注
}

// genTableColumns holds the columns for table gen_table.
var genTableColumns = GenTableColumns{
	TableId:        "table_id",
	TableName:      "table_name",
	TableComment:   "table_comment",
	SubTableName:   "sub_table_name",
	SubTableFkName: "sub_table_fk_name",
	ClassName:      "class_name",
	TplCategory:    "tpl_category",
	PackageName:    "package_name",
	ModuleName:     "module_name",
	BusinessName:   "business_name",
	FunctionName:   "function_name",
	FunctionAuthor: "function_author",
	GenType:        "gen_type",
	Status:         "status",
	GenPath:        "gen_path",
	Options:        "options",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
	Remark:         "remark",
}

// NewGenTableDao creates and returns a new DAO object for table data access.
func NewGenTableDao() *GenTableDao {
	return &GenTableDao{
		group:   "default",
		table:   "gen_table",
		columns: genTableColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GenTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GenTableDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GenTableDao) Columns() GenTableColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GenTableDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GenTableDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GenTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
