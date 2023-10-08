// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GenTableColumnDao is the data access object for table gen_table_column.
type GenTableColumnDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns GenTableColumnColumns // columns contains all the column names of Table for convenient usage.
}

// GenTableColumnColumns defines and stores column names for table gen_table_column.
type GenTableColumnColumns struct {
	ColumnId      string // 编号
	TableId       string // 归属表编号
	ColumnName    string // 列名称
	ColumnComment string // 列描述
	ColumnType    string // 列类型
	ColumnDef     string // 列默认值
	GoType        string // go类型
	GoField       string // go字段名
	JsonField     string // json属性名
	IsPk          string // 是否主键（1是）
	IsIncrement   string // 是否自增（1是）
	IsRequired    string // 是否必填（1是）
	IsInsert      string // 是否为插入字段（1是）
	IsEdit        string // 是否编辑字段（1是）
	IsList        string // 是否列表字段（1是）
	IsQuery       string // 是否查询字段（1是）
	QueryType     string // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string // 字典类型
	Sort          string // 排序
	CreateBy      string // 创建者
	CreateTime    string // 创建时间
	UpdateBy      string // 更新者
	UpdateTime    string // 更新时间
}

// genTableColumnColumns holds the columns for table gen_table_column.
var genTableColumnColumns = GenTableColumnColumns{
	ColumnId:      "column_id",
	TableId:       "table_id",
	ColumnName:    "column_name",
	ColumnComment: "column_comment",
	ColumnType:    "column_type",
	ColumnDef:     "column_def",
	GoType:        "go_type",
	GoField:       "go_field",
	JsonField:     "json_field",
	IsPk:          "is_pk",
	IsIncrement:   "is_increment",
	IsRequired:    "is_required",
	IsInsert:      "is_insert",
	IsEdit:        "is_edit",
	IsList:        "is_list",
	IsQuery:       "is_query",
	QueryType:     "query_type",
	HtmlType:      "html_type",
	DictType:      "dict_type",
	Sort:          "sort",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
}

// NewGenTableColumnDao creates and returns a new DAO object for table data access.
func NewGenTableColumnDao() *GenTableColumnDao {
	return &GenTableColumnDao{
		group:   "default",
		table:   "gen_table_column",
		columns: genTableColumnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GenTableColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GenTableColumnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GenTableColumnDao) Columns() GenTableColumnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GenTableColumnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GenTableColumnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GenTableColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
