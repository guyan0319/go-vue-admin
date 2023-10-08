// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotice is the golang structure of table sys_notice for DAO operations like Where/Data.
type SysNotice struct {
	g.Meta        `orm:"table:sys_notice, do:true"`
	NoticeId      interface{} // 公告ID
	NoticeTitle   interface{} // 公告标题
	NoticeType    interface{} // 公告类型（1通知 2公告）
	NoticeContent []byte      // 公告内容
	Status        interface{} // 公告状态（0正常 1关闭）
	CreateBy      interface{} // 创建者
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      interface{} // 更新者
	UpdateTime    *gtime.Time // 更新时间
	Remark        interface{} // 备注
}
