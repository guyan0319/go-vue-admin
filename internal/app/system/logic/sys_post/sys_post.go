package sys_post

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
	"go-vue-admin/utility/lib"
)

func init() {
	service.RegisterSysPost(New())
}

func New() *sSysPost {
	return &sSysPost{}
}

type sSysPost struct {
}

func (s *sSysPost) GetAllPostByStatus(ctx context.Context, status string) (posts []*entity.SysPost, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysPost.Ctx(ctx)
		if status != "" {
			m = m.Where(dao.SysPost.Columns().Status, status)
		}
		err = m.Scan(&posts)
		utility.WriteErrLog(ctx, err, "获取用户岗位数据失败")
	})
	return
}

// 部门列表
func (s *sSysPost) GetPostList(ctx context.Context, req *v1.GetPostListReq) (postList *v1.GetPostListRes, err error) {

	postList = &v1.GetPostListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysPost
		m := dao.SysPost.Ctx(ctx)
		if req.Status != "" {
			m = m.Where(dao.SysPost.Columns().Status, req.Status)
		}
		//时间
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysPost.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		if req.PostName != "" {
			m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+req.PostName+"%")
		}
		if req.PostCode != "" {
			m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+req.PostCode+"%")
		}
		postList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)
		utility.WriteErrLog(ctx, err, "获取部门失败")
		postList.List = list
	})

	return
}

// 添加岗位数据
func (s sSysPost) Add(ctx context.Context, req *v1.PostPostAddReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysPost.Ctx(ctx).Data(do.SysPost{
			PostCode:   req.PostCode,
			PostName:   req.PostName,
			PostSort:   req.PostSort,
			Status:     req.Status,
			Remark:     req.Remark,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
			CreateTime: gtime.Now(),
			CreateBy:   adminName,
		}).Insert()
		utility.WriteErrLog(ctx, e, "添加岗位数据失败")
	})
	return
}

// 修改数据
func (s sSysPost) Update(ctx context.Context, req *v1.PutPostUpdateReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysPost.Ctx(ctx).WherePri(&req.PostId).Update(do.SysPost{
			PostCode:   req.PostCode,
			PostName:   req.PostName,
			PostSort:   req.PostSort,
			Status:     req.Status,
			Remark:     req.Remark,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "修改岗位数据失败")
	})
	return
}

// 删除数据
func (s sSysPost) Delete(ctx context.Context, req *v1.DeletePostReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		postIds := lib.ParamStrToSlice(req.PostId, ",")
		_, e := dao.SysPost.Ctx(ctx).WhereIn(dao.SysPost.Columns().PostId, postIds).Update(do.SysPost{
			Status:     consts.SysPostStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除岗位数据失败")
	})
	return
}

func (s *sSysPost) GetPostById(ctx context.Context, id int64) (post *entity.SysPost, err error) {
	err = g.Try(ctx, func(ctx context.Context) {

		//用户岗位信息
		err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().PostId, id).Scan(&post)
		utility.WriteErrLog(ctx, err, "获取用户岗位数据失败")
	})
	return
}
