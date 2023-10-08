// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure of table gen_table for DAO operations like Where/Data.
type GenTable struct {
	g.Meta         `orm:"table:gen_table, do:true"`
	TableId        interface{} // 编号
	TableName      interface{} // 表名称
	TableComment   interface{} // 表描述
	SubTableName   interface{} // 关联子表的表名
	SubTableFkName interface{} // 子表关联的外键名
	ClassName      interface{} // 实体类名称
	TplCategory    interface{} // 使用的模板（crud单表操作 tree树表操作）
	PackageName    interface{} // 生成包路径
	ModuleName     interface{} // 生成模块名
	BusinessName   interface{} // 生成业务名
	FunctionName   interface{} // 生成功能名
	FunctionAuthor interface{} // 生成功能作者
	GenType        interface{} // 生成代码方式（0zip压缩包 1自定义路径）
	Status         interface{} // 状态（0正常 1删除）
	GenPath        interface{} // 生成路径（不填默认项目路径）
	Options        interface{} // 其它生成选项
	CreateBy       interface{} // 创建者
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       interface{} // 更新者
	UpdateTime     *gtime.Time // 更新时间
	Remark         interface{} // 备注
}
