// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure of table gen_table_column for DAO operations like Where/Data.
type GenTableColumn struct {
	g.Meta        `orm:"table:gen_table_column, do:true"`
	ColumnId      interface{} // 编号
	TableId       interface{} // 归属表编号
	ColumnName    interface{} // 列名称
	ColumnComment interface{} // 列描述
	ColumnType    interface{} // 列类型
	ColumnDef     interface{} // 列默认值
	GoType        interface{} // go类型
	GoField       interface{} // go字段名
	JsonField     interface{} // json属性名
	IsPk          interface{} // 是否主键（1是）
	IsIncrement   interface{} // 是否自增（1是）
	IsRequired    interface{} // 是否必填（1是）
	IsInsert      interface{} // 是否为插入字段（1是）
	IsEdit        interface{} // 是否编辑字段（1是）
	IsList        interface{} // 是否列表字段（1是）
	IsQuery       interface{} // 是否查询字段（1是）
	QueryType     interface{} // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      interface{} // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      interface{} // 字典类型
	Sort          interface{} // 排序
	CreateBy      interface{} // 创建者
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      interface{} // 更新者
	UpdateTime    *gtime.Time // 更新时间
}
