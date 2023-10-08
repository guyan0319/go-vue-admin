// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta     `orm:"table:sys_dict_type, do:true"`
	DictId     interface{} // 字典主键
	DictName   interface{} // 字典名称
	DictType   interface{} // 字典类型
	Status     interface{} // 状态（0正常 1停用）
	CreateBy   interface{} // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   interface{} // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     interface{} // 备注
}
