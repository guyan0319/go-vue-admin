package model

import "go-vue-admin/internal/app/system/model/entity"

type RolePerm struct {
	AllPerm  []string
	MapPerms map[int64][]string
}
type RoleMenu struct {
	NenuMap map[int64][]string
}

type PermsData struct {
	RoleId int64  `json:"roleId" description:"角色ID"`
	Perms  string `json:"perms"      description:"权限标识"`
}
type UserMenuRes struct {
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Hidden     bool           `json:"hidden"`
	Redirect   string         `json:"redirect"`
	Component  string         `json:"component"`
	AlwaysShow bool           `json:"alwaysShow"`
	Meta       *MenuMeta      `json:"meta"`
	Children   []*UserMenuRes `json:"children"`
}
type MenuMeta struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Link    string `json:"link"`
}

type SysMenuTreeRes struct {
	Id       int64             `json:"id"     description:"菜单id"`
	Label    string            `json:"label"   description:"菜单名称"`
	Children []*SysMenuTreeRes `json:"children"   description:"子菜单"`
}

type SysMenuList struct {
	*entity.SysMenu
	ParentName string         `json:"parentName"   description:"父菜单名称"`
	Children   []*SysMenuList `json:"children"   description:"子菜单"`
}
