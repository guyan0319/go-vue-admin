// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model/entity"
)

type (
	ISysPost interface {
		GetAllPostByStatus(ctx context.Context, status string) (posts []*entity.SysPost, err error)
		// 部门列表
		GetPostList(ctx context.Context, req *v1.GetPostListReq) (postList *v1.GetPostListRes, err error)
		// 添加岗位数据
		Add(ctx context.Context, req *v1.PostPostAddReq) (err error)
		// 修改数据
		Update(ctx context.Context, req *v1.PutPostUpdateReq) (err error)
		// 删除数据
		Delete(ctx context.Context, req *v1.DeletePostReq) (err error)
		GetPostById(ctx context.Context, id int64) (post *entity.SysPost, err error)
	}
)

var (
	localSysPost ISysPost
)

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}
