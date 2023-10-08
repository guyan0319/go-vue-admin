package sys_user_post

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
)

func init() {
	service.RegisterSysUserPost(New())
}

func New() *sSysUserPost {
	return &sSysUserPost{}
}

type sSysUserPost struct {
}

func (s *sSysUserPost) GetPostIdByUid(ctx context.Context, uid int64) (postId []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var user []*entity.SysUserPost
		//用户用户信息
		err = dao.SysUserPost.Ctx(ctx).Fields("post_id").Where(dao.SysUserPost.Columns().UserId, uid).Scan(&user)
		utility.WriteErrLog(ctx, err, "获取用户数据失败")
		for _, v := range user {
			postId = append(postId, v.PostId)
		}
	})
	return
}

func (s *sSysUserPost) AddUserPosts(ctx context.Context, tx gdb.TX, userId int64, PostIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧用户角色
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Where(dao.SysUserPost.Columns().UserId, userId).Delete()
		utility.WriteErrLog(ctx, err, "删除用户角色失败")
		if len(PostIds) == 0 {
			return
		}
		//添加用户角色信息
		data := g.List{}
		for _, v := range PostIds {
			data = append(data, g.Map{
				dao.SysUserPost.Columns().UserId: userId,
				dao.SysUserPost.Columns().PostId: v,
			})
		}
		_, err = dao.SysUserPost.Ctx(ctx).TX(tx).Data(data).Insert()
		utility.WriteErrLog(ctx, err, "添加用户职位失败")
	})
	return
}
