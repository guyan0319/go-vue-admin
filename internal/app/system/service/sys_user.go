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
	ISysUser interface {
		// Create creates user account.
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error)
		Udate(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error)
		// 假删除 支持批量删除
		Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
		// 更改状态
		ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error)
		// 更改密码
		ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error)
		// IsSignedIn checks and returns whether current user is already signed-in.
		IsSignedIn(ctx context.Context) bool
		// IsUserNameAvailable checks and returns given UserName is available for signing up.
		IsUserNameAvailable(ctx context.Context, UserName string) (bool, error)
		// IsNicknameAvailable checks and returns given nickname is available for signing up.
		IsNicknameAvailable(ctx context.Context, nickname string, userId int64) (bool, error)
		// GetProfile retrieves and returns current user info in session.
		GetProfile(ctx context.Context) *entity.SysUser
		GetUserById(ctx context.Context, userId int64) (user *model.SysUserRes, err error)
		GetOneUserById(ctx context.Context, id int64) (user *entity.SysUser, err error)
		GetUserListByDeptId(ctx context.Context, req *v1.GetUserListReq) (userList *v1.GetUserListRes, err error)
		GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error)
		PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
