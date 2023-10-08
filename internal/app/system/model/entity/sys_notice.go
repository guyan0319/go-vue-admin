// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotice is the golang structure for table sys_notice.
type SysNotice struct {
	NoticeId      int         `json:"noticeId"      description:"公告ID"`
	NoticeTitle   string      `json:"noticeTitle"   description:"公告标题"`
	NoticeType    string      `json:"noticeType"    description:"公告类型（1通知 2公告）"`
	NoticeContent []byte      `json:"noticeContent" description:"公告内容"`
	Status        string      `json:"status"        description:"公告状态（0正常 1关闭）"`
	CreateBy      string      `json:"createBy"      description:"创建者"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
	UpdateBy      string      `json:"updateBy"      description:"更新者"`
	UpdateTime    *gtime.Time `json:"updateTime"    description:"更新时间"`
	Remark        string      `json:"remark"        description:"备注"`
}
