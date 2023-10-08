package model

import "go-vue-admin/internal/app/system/model/entity"

type SysRoleRes struct {
	*entity.SysRole
	Flag        bool     `json:"flag" `
	MenuIds     string   `json:"menuIds" `
	DeptIds     string   `json:"deptIds" `
	Admin       bool     `json:"admin"            description:"是否是admin"`
	Permissions []string `json:"permissions"            description:"权限"`
}
type SysRolesRes struct {
	RoleIds []int64           `json:"roleIds" `
	Roles   []string          `json:"roles" `
	SysRole []*entity.SysRole `json:"SysRole" `
}

type RoleList struct {
	*entity.SysRole
	Flag    bool   `json:"flag" `
	MenuIds string `json:"menuIds" `
	DeptIds string `json:"deptIds" `
	Admin   bool   `json:"admin"            description:"是否是admin"`
}
