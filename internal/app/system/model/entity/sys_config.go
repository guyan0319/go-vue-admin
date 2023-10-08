// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId    int         `json:"configId"    description:"参数主键"`
	ConfigName  string      `json:"configName"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   description:"参数键名"`
	ConfigValue string      `json:"configValue" description:"参数键值"`
	ConfigType  string      `json:"configType"  description:"系统内置（Y是 N否）"`
	CreateBy    string      `json:"createBy"    description:"创建者"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
	Remark      string      `json:"remark"      description:"备注"`
}
