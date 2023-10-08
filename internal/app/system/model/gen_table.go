package model

import (
	"go-vue-admin/internal/app/system/model/entity"
)

type GenTableList struct {
	*entity.GenTable
	Columns *entity.GenTableColumn `json:"columns"   description:"表字段"`
}

type GenTableInfo struct {
	*entity.GenTable
	Columns []*entity.GenTableColumn `json:"columns"   description:"表字段"`
}
type GenTableUpdate struct {
	TableId        int64                    `p:"tableId"       v:"required"  description:"编号"`
	TableName      string                   `p:"tableName"     v:"required" description:"表名称"`
	TableComment   string                   `p:"tableComment"  v:"required"  description:"表描述"`
	SubTableName   string                   `p:"subTableName"   description:"关联子表的表名"`
	SubTableFkName string                   `p:"subTableFkName" description:"子表关联的外键名"`
	ClassName      string                   `p:"className"     v:"required" description:"实体类名称"`
	TplCategory    string                   `p:"tplCategory"   v:"required" description:"使用的模板（crud单表操作 tree树表操作）"`
	PackageName    string                   `p:"packageName"  v:"required"  description:"生成包路径"`
	ModuleName     string                   `p:"moduleName"     description:"生成模块名"`
	BusinessName   string                   `p:"businessName"   description:"生成业务名"`
	FunctionName   string                   `p:"functionName"   description:"生成功能名"`
	FunctionAuthor string                   `p:"functionAuthor" description:"生成功能作者"`
	GenType        string                   `p:"genType"        description:"生成代码方式（0zip压缩包 1自定义路径）"`
	Status         string                   `p:"status"         description:"状态（0正常 1删除）"`
	GenPath        string                   `p:"genPath"        description:"生成路径（不填默认项目路径）"`
	Options        string                   `p:"options"        description:"其它生成选项"`
	Remark         string                   `p:"remark"         description:"备注"`
	Columns        []*entity.GenTableColumn `json:"columns"   description:"表字段"`
}

type TplContent struct {
	TypeName    string
	TypeContent string
}

// controller 参数
type TplCtrlContent struct {
	FuncName       string
	ServiceReq     string
	ServiceRes     string
	ServiceName    string
	ApiPackageName string
	NameController string
}

// logic 参数
type TplLogicContent struct {
	FuncName       string `json:"funcName"      description:"函数名"`
	ParamReq       string `json:"paramReq"      description:"输入参数"`
	ParamRes       string `json:"paramRes"      description:"输出参数"`
	Notes          string `json:"paramReq"      description:"注释"`
	ApiPackageName string `json:"apiPackageName"      description:"接口包名"`
	ClassName      string `json:"className"      description:"结构体名"`
	Content        string `json:"content"      description:"内容"`
}

// api js 参数
type TplApiJsContent struct {
	FuncName string `json:"funcName"      description:"函数名"`
	Notes    string `json:"paramReq"      description:"注释"`
	Content  string `json:"content"      description:"内容"`
}

// vue 检索字段 参数
type TplVueSearchContent struct {
	ColumnComment string `json:"columnComment"      description:"label"`
	JsonField     string `json:"jsonField"      description:"pop"`
	Content       string `json:"content"      description:"内容"`
}
type TplVueContent struct {
	Search         []TplVueSearchContent
	DictType       string `json:"dictType"      description:"调入字典类型"`
	ImportApiFunc  string `json:"importApiFunc"      description:"引入api方法"`
	SearchParam    string `json:"searchParam"      description:"查询参数"`
	ColumnRules    string `json:"ColumnRules"      description:"字段验证规则"`
	ClaseName      string `json:"claseName"      description:"方法后缀"`
	TableComment   string `json:"tableComment"      description:"表注释"`
	ColumnReset    string `json:"columnReset"      description:"字段初始化"`
	PKComment      string `json:"pKComment"      description:"主键字段注释"`
	PKJsonField    string `json:"pKJsonField"      description:"主键字段jsonfield"`
	EditColumn     string `json:"editColumn"      description:"编辑字段"`
	ListColumn     string `json:"listColumn"      description:"列表字段"`
	ResetQuery     string `json:"resetQuery"      description:"重置查询"`
	ListParamQuery string `json:"listParamQuery"      description:"查询传参"`
	PermiPath      string `json:"permiPath"      description:"权限路径"`
}
