package post

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type postController struct {
}

var Post = postController{}

func (s *postController) GetPostList(ctx context.Context, req *v1.GetPostListReq) (res *v1.GetPostListRes, err error) {
	res, err = service.SysPost().GetPostList(ctx, req)
	return
}
func (s *postController) Add(ctx context.Context, req *v1.PostPostAddReq) (res *v1.PostPostAddRes, err error) {
	err = service.SysPost().Add(ctx, req)
	return
}
func (s *postController) Update(ctx context.Context, req *v1.PutPostUpdateReq) (res *v1.PutPostUpdateRes, err error) {
	err = service.SysPost().Update(ctx, req)
	return
}
func (s *postController) Delete(ctx context.Context, req *v1.DeletePostReq) (res *v1.DeletePostRes, err error) {
	err = service.SysPost().Delete(ctx, req)
	return
}

// 获取岗位部门数据
func (c *postController) GetPostUpdate(ctx context.Context, req *v1.GetPostUpdateReq) (res *v1.GetPostUpdateRes, err error) {
	res = &v1.GetPostUpdateRes{}
	res.SysPost, err = service.SysPost().GetPostById(ctx, req.PostId)
	return
}
