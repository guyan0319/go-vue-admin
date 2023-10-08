// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	MenuId     interface{} // 菜单ID
	MenuName   interface{} // 菜单名称
	ParentId   interface{} // 父菜单ID
	OrderNum   interface{} // 显示顺序
	Path       interface{} // 路由地址
	ApiPath    interface{} // 后台api路径
	Component  interface{} // 组件路径
	Query      interface{} // 路由参数
	IsFrame    interface{} // 是否为外链（0是 1否）
	IsCache    interface{} // 是否缓存（0缓存 1不缓存）
	MenuType   interface{} // 菜单类型（M目录 C菜单 F按钮）
	Visible    interface{} // 菜单状态（0显示 1隐藏）
	Status     interface{} // 菜单状态（0正常 1停用）
	Perms      interface{} // 权限标识
	Icon       interface{} // 菜单图标
	CreateBy   interface{} // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   interface{} // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     interface{} // 备注
}
