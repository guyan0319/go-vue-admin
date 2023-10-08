// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	DictId     int64       `json:"dictId"     description:"字典主键"`
	DictName   string      `json:"dictName"   description:"字典名称"`
	DictType   string      `json:"dictType"   description:"字典类型"`
	Status     string      `json:"status"     description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
	Remark     string      `json:"remark"     description:"备注"`
}
