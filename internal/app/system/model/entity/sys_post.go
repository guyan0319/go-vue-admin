// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId     int64       `json:"postId"     description:"岗位ID"`
	PostCode   string      `json:"postCode"   description:"岗位编码"`
	PostName   string      `json:"postName"   description:"岗位名称"`
	PostSort   int         `json:"postSort"   description:"显示顺序"`
	Status     string      `json:"status"     description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
	Remark     string      `json:"remark"     description:"备注"`
}
