// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogininfor is the golang structure of table sys_logininfor for DAO operations like Where/Data.
type SysLogininfor struct {
	g.Meta        `orm:"table:sys_logininfor, do:true"`
	InfoId        interface{} // 访问ID
	UserName      interface{} // 用户账号
	Ipaddr        interface{} // 登录IP地址
	LoginLocation interface{} // 登录地点
	Browser       interface{} // 浏览器类型
	Os            interface{} // 操作系统
	Status        interface{} // 登录状态（0成功 1失败）
	Msg           interface{} // 提示消息
	LoginTime     *gtime.Time // 访问时间
}
