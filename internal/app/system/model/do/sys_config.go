// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta      `orm:"table:sys_config, do:true"`
	ConfigId    interface{} // 参数主键
	ConfigName  interface{} // 参数名称
	ConfigKey   interface{} // 参数键名
	ConfigValue interface{} // 参数键值
	ConfigType  interface{} // 系统内置（Y是 N否）
	CreateBy    interface{} // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    interface{} // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      interface{} // 备注
}
