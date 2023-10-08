// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogininfor is the golang structure for table sys_logininfor.
type SysLogininfor struct {
	InfoId        int64       `json:"infoId"        description:"访问ID"`
	UserName      string      `json:"userName"      description:"用户账号"`
	Ipaddr        string      `json:"ipaddr"        description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" description:"登录地点"`
	Browser       string      `json:"browser"       description:"浏览器类型"`
	Os            string      `json:"os"            description:"操作系统"`
	Status        string      `json:"status"        description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"           description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     description:"访问时间"`
}
