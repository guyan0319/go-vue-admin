// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	DictCode   int64       `json:"dictCode"   description:"字典编码"`
	DictSort   int         `json:"dictSort"   description:"字典排序"`
	DictLabel  string      `json:"dictLabel"  description:"字典标签"`
	DictValue  string      `json:"dictValue"  description:"字典键值"`
	DictType   string      `json:"dictType"   description:"字典类型"`
	CssClass   string      `json:"cssClass"   description:"样式属性（其他样式扩展）"`
	ListClass  string      `json:"listClass"  description:"表格回显样式"`
	IsDefault  string      `json:"isDefault"  description:"是否默认（Y是 N否）"`
	Status     string      `json:"status"     description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
	Remark     string      `json:"remark"     description:"备注"`
}
