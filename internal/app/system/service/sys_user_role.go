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
	ISysUserRole interface {
		AddUserRoles(ctx context.Context, tx gdb.TX, userId int64, roleIds []int64) (err error)
		// 角色分配用户
		AddRoleUsers(ctx context.Context, tx gdb.TX, roleId int64, userIds []int64) (err error)
		// 取消角色分配用户
		CancelRoleUsers(ctx context.Context, tx gdb.TX, roleId int64, userIds []int64) (err error)
		GetRoleIdByUid(ctx context.Context, uid int64) (roleId []int64, err error)
		GetUserIdByRoleId(ctx context.Context, roleId int64) (userId []int64, err error)
	}
)

var (
	localSysUserRole ISysUserRole
)

func SysUserRole() ISysUserRole {
	if localSysUserRole == nil {
		panic("implement not found for interface ISysUserRole, forgot register?")
	}
	return localSysUserRole
}

func RegisterSysUserRole(i ISysUserRole) {
	localSysUserRole = i
}
