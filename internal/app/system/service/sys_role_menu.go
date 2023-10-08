// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysRoleMenu interface {
		AddRoleMenus(ctx context.Context, tx gdb.TX, roleId int64, MenuIds []int64) (err error)
		GetMenuIdsByRoleId(ctx context.Context, roleId int64) (menuIds []int64, err error)
	}
)

var (
	localSysRoleMenu ISysRoleMenu
)

func SysRoleMenu() ISysRoleMenu {
	if localSysRoleMenu == nil {
		panic("implement not found for interface ISysRoleMenu, forgot register?")
	}
	return localSysRoleMenu
}

func RegisterSysRoleMenu(i ISysRoleMenu) {
	localSysRoleMenu = i
}
