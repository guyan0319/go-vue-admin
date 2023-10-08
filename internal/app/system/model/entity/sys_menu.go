// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	MenuId     int64       `json:"menuId"     description:"菜单ID"`
	MenuName   string      `json:"menuName"   description:"菜单名称"`
	ParentId   int64       `json:"parentId"   description:"父菜单ID"`
	OrderNum   int         `json:"orderNum"   description:"显示顺序"`
	Path       string      `json:"path"       description:"路由地址"`
	ApiPath    string      `json:"apiPath"    description:"后台api路径"`
	Component  string      `json:"component"  description:"组件路径"`
	Query      string      `json:"query"      description:"路由参数"`
	IsFrame    int         `json:"isFrame"    description:"是否为外链（0是 1否）"`
	IsCache    int         `json:"isCache"    description:"是否缓存（0缓存 1不缓存）"`
	MenuType   string      `json:"menuType"   description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible    string      `json:"visible"    description:"菜单状态（0显示 1隐藏）"`
	Status     string      `json:"status"     description:"菜单状态（0正常 1停用）"`
	Perms      string      `json:"perms"      description:"权限标识"`
	Icon       string      `json:"icon"       description:"菜单图标"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
	Remark     string      `json:"remark"     description:"备注"`
}
