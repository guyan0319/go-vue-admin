// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure for table gen_table.
type GenTable struct {
	TableId        int64       `json:"tableId"        description:"编号"`
	TableName      string      `json:"tableName"      description:"表名称"`
	TableComment   string      `json:"tableComment"   description:"表描述"`
	SubTableName   string      `json:"subTableName"   description:"关联子表的表名"`
	SubTableFkName string      `json:"subTableFkName" description:"子表关联的外键名"`
	ClassName      string      `json:"className"      description:"实体类名称"`
	TplCategory    string      `json:"tplCategory"    description:"使用的模板（crud单表操作 tree树表操作）"`
	PackageName    string      `json:"packageName"    description:"生成包路径"`
	ModuleName     string      `json:"moduleName"     description:"生成模块名"`
	BusinessName   string      `json:"businessName"   description:"生成业务名"`
	FunctionName   string      `json:"functionName"   description:"生成功能名"`
	FunctionAuthor string      `json:"functionAuthor" description:"生成功能作者"`
	GenType        string      `json:"genType"        description:"生成代码方式（0zip压缩包 1自定义路径）"`
	Status         string      `json:"status"         description:"状态（0正常 1删除）"`
	GenPath        string      `json:"genPath"        description:"生成路径（不填默认项目路径）"`
	Options        string      `json:"options"        description:"其它生成选项"`
	CreateBy       string      `json:"createBy"       description:"创建者"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建时间"`
	UpdateBy       string      `json:"updateBy"       description:"更新者"`
	UpdateTime     *gtime.Time `json:"updateTime"     description:"更新时间"`
	Remark         string      `json:"remark"         description:"备注"`
}
