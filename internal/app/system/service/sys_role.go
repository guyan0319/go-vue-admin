// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type (
	ISysRole interface {
		GetNomalRole(ctx context.Context) (roles []*entity.SysRole, err error)
		GetRoleByUid(ctx context.Context, uid int64) (roles []*entity.SysRole, err error)
		GetRolesByUid(ctx context.Context, uid int64) (rolesList *model.SysRolesRes, err error)
		GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (roleList *v1.GetRoleListRes, err error)
		IsRoleNameAvailable(ctx context.Context, roleName string, roleId int64) (bool, error)
		Add(ctx context.Context, req *v1.PostRoleReq) (res *v1.PostRoleRes, err error)
		GetRoleUpdate(ctx context.Context, req *v1.GetRoleUpdateReq) (res *v1.GetRoleUpdateRes, err error)
		GetRoleUpdateTreeSelect(ctx context.Context, req *v1.GetRoleUpdateTreeSelectReq) (res *v1.GetRoleUpdateTreeSelectRes, err error)
		Update(ctx context.Context, req *v1.PutRoleUpdateReq) (res *v1.PutRoleUpdateRes, err error)
		UpdateDataScope(ctx context.Context, req *v1.PutRoleDataScopeReq) (res *v1.PutRoleDataScopeRes, err error)
		// 更角色改状态
		ChangeStatus(ctx context.Context, req *v1.ChangeStatusRoleReq) (res *v1.ChangeStatusRoleRes, err error)
		// 更角色改状态
		Delete(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error)
		// 获取分配用户
		GetRoleAuthUser(ctx context.Context, req *v1.GetRoleAuthUserReq) (res *v1.GetRoleAuthUserRes, err error)
		// 获取添加分配用户
		GetRoleAddAuthUser(ctx context.Context, req *v1.GetRoleAddAuthUserReq) (res *v1.GetRoleAddAuthUserRes, err error)
		// 添加分配用户
		PutRoleAddAuthUser(ctx context.Context, req *v1.PutRoleAddAuthUserReq) (res *v1.PutRoleAddAuthUserRes, err error)
		// 取消分配用户
		PutRoleCancelAuthUser(ctx context.Context, req *v1.PutRoleCancelAuthUserReq) (res *v1.PutRoleCancelAuthUserRes, err error)
		// 批量取消分配用户
		PutRoleCancelAllAuthUser(ctx context.Context, req *v1.PutRoleCancelAllAuthUserReq) (res *v1.PutRoleCancelAllAuthUserRes, err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
