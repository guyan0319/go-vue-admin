// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure for table gen_table_column.
type GenTableColumn struct {
	ColumnId      int64       `json:"columnId"      description:"编号"`
	TableId       string      `json:"tableId"       description:"归属表编号"`
	ColumnName    string      `json:"columnName"    description:"列名称"`
	ColumnComment string      `json:"columnComment" description:"列描述"`
	ColumnType    string      `json:"columnType"    description:"列类型"`
	ColumnDef     string      `json:"columnDef"     description:"列默认值"`
	GoType        string      `json:"goType"        description:"go类型"`
	GoField       string      `json:"goField"       description:"go字段名"`
	JsonField     string      `json:"jsonField"     description:"json属性名"`
	IsPk          string      `json:"isPk"          description:"是否主键（1是）"`
	IsIncrement   string      `json:"isIncrement"   description:"是否自增（1是）"`
	IsRequired    string      `json:"isRequired"    description:"是否必填（1是）"`
	IsInsert      string      `json:"isInsert"      description:"是否为插入字段（1是）"`
	IsEdit        string      `json:"isEdit"        description:"是否编辑字段（1是）"`
	IsList        string      `json:"isList"        description:"是否列表字段（1是）"`
	IsQuery       string      `json:"isQuery"       description:"是否查询字段（1是）"`
	QueryType     string      `json:"queryType"     description:"查询方式（等于、不等于、大于、小于、范围）"`
	HtmlType      string      `json:"htmlType"      description:"显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）"`
	DictType      string      `json:"dictType"      description:"字典类型"`
	Sort          int         `json:"sort"          description:"排序"`
	CreateBy      string      `json:"createBy"      description:"创建者"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
	UpdateBy      string      `json:"updateBy"      description:"更新者"`
	UpdateTime    *gtime.Time `json:"updateTime"    description:"更新时间"`
}
