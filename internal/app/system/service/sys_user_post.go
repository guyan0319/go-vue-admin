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
	ISysUserPost interface {
		GetPostIdByUid(ctx context.Context, uid int64) (postId []int64, err error)
		AddUserPosts(ctx context.Context, tx gdb.TX, userId int64, PostIds []int64) (err error)
	}
)

var (
	localSysUserPost ISysUserPost
)

func SysUserPost() ISysUserPost {
	if localSysUserPost == nil {
		panic("implement not found for interface ISysUserPost, forgot register?")
	}
	return localSysUserPost
}

func RegisterSysUserPost(i ISysUserPost) {
	localSysUserPost = i
}
