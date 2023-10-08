package sys_dict_type

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
	service.RegisterSysDictType(New())
}

func New() *sSysDictType {
	return &sSysDictType{}
}

type sSysDictType struct {
}

// 字典类型表
func (s *sSysDictType) GetDictTypeOption(ctx context.Context, req *v1.GetDictTypeOptionSelectReq) (res *v1.GetDictTypeOptionSelectRes, err error) {
	res = &v1.GetDictTypeOptionSelectRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var dictType []*entity.SysDictType
		//字典类型表
		err = dao.SysDictType.Ctx(ctx).Scan(&dictType)
		utility.WriteErrLog(ctx, err, "获取字典类型数据表失败")
		res.DictType = dictType
	})
	return
}

// 获取字典类型表列表
func (s *sSysDictType) GetSysDictTypeList(ctx context.Context, req *v1.GetSysDictTypeListReq) (res *v1.GetSysDictTypeListRes, err error) {
	res = &v1.GetSysDictTypeListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysDictType
		m := dao.SysDictType.Ctx(ctx)
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysDictType.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		if req.DictName != "" {
			m = m.WhereLike(dao.SysDictType.Columns().DictName, "%"+req.DictName+"%")
		}
		if req.DictType != "" {
			m = m.Where(dao.SysDictType.Columns().DictType, req.DictType)
		}
		if req.Status != "" {
			m = m.Where(dao.SysDictType.Columns().Status, req.Status)
		}
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)

		utility.WriteErrLog(ctx, err, "获取字典类型表失败")
		res.List = list
	})

	return
}

// 添加字典类型表
func (s *sSysDictType) Add(ctx context.Context, req *v1.PostSysDictTypeReq) (res *v1.PostSysDictTypeRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysDictType.Ctx(ctx).Data(do.SysDictType{
			Remark:     req.Remark,
			DictName:   req.DictName,
			DictType:   req.DictType,
			Status:     req.Status,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
			CreateTime: gtime.Now(),
			CreateBy:   adminName,
		}).Insert()
		utility.WriteErrLog(ctx, e, "添加字典类型表数据失败")
	})

	return
}

// 修改字典类型表
func (s *sSysDictType) Update(ctx context.Context, req *v1.PutSysDictTypeReq) (res *v1.PutSysDictTypeRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysDictType.Ctx(ctx).WherePri(&req.DictId).Update(do.SysDictType{
			Remark:     req.Remark,
			DictName:   req.DictName,
			DictType:   req.DictType,
			Status:     req.Status,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "修改字典类型表数据失败")
	})

	return
}

// 删除字典类型表
func (s *sSysDictType) Delete(ctx context.Context, req *v1.DeleteSysDictTypeReq) (res *v1.DeleteSysDictTypeRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		postIds := lib.ParamStrToSlice(req.DictId, ",")
		_, e := dao.SysDictType.Ctx(ctx).WhereIn(dao.SysDictType.Columns().DictId, postIds).Update(do.SysDictType{
			Status:     consts.SysDictTypeStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除字典类型表数据失败")
	})

	return
}

// 获取字典类型表
func (s *sSysDictType) GetSysDictType(ctx context.Context, req *v1.GetSysDictTypeReq) (res *v1.GetSysDictTypeRes, err error) {
	res = &v1.GetSysDictTypeRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var table *entity.SysDictType
		err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictId, req.DictId).Scan(&table)
		res.SysDictType = table
		utility.WriteErrLog(ctx, err, "删除字典类型表数据失败")
	})

	return
}
