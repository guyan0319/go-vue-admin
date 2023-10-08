package sys_config

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
	service.RegisterSysConfig(New())
}

func New() *sSysConfig {
	return &sSysConfig{}
}

type sSysConfig struct {
}

// 获取参数配置表列表
func (s *sSysConfig) GetSysConfigList(ctx context.Context, req *v1.GetSysConfigListReq) (res *v1.GetSysConfigListRes, err error) {
	res = &v1.GetSysConfigListRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var list []*entity.SysConfig
		m := dao.SysConfig.Ctx(ctx)
		if req.ConfigName != "" {
			m = m.WhereLike(dao.SysConfig.Columns().ConfigName, "%"+req.ConfigName+"%")
		}
		if req.ConfigKey != "" {
			m = m.Where(dao.SysConfig.Columns().ConfigKey, req.ConfigKey)
		}
		if req.ConfigType != "" {
			m = m.Where(dao.SysConfig.Columns().ConfigType, req.ConfigType)
		}
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysConfig.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&list)

		utility.WriteErrLog(ctx, err, "获取参数配置表失败")
		res.Rows = list
	})

	return
}

// 添加参数配置表
func (s *sSysConfig) Add(ctx context.Context, req *v1.PostSysConfigReq) (res *v1.PostSysConfigRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysConfig.Ctx(ctx).Data(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			Remark:      req.Remark,
			UpdateTime:  gtime.Now(),
			UpdateBy:    adminName,
			CreateTime:  gtime.Now(),
			CreateBy:    adminName,
		}).Insert()
		utility.WriteErrLog(ctx, e, "添加参数配置表数据失败")
	})

	return
}

// 修改参数配置表
func (s *sSysConfig) Update(ctx context.Context, req *v1.PutSysConfigReq) (res *v1.PutSysConfigRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysConfig.Ctx(ctx).WherePri(&req.ConfigId).Update(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			Remark:      req.Remark,
			UpdateTime:  gtime.Now(),
			UpdateBy:    adminName,
		})
		utility.WriteErrLog(ctx, e, "修改参数配置表数据失败")
	})

	return
}

// 删除参数配置表
func (s *sSysConfig) Delete(ctx context.Context, req *v1.DeleteSysConfigReq) (res *v1.DeleteSysConfigRes, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		postIds := lib.ParamStrToSlice(req.ConfigId, ",")
		_, e := dao.SysConfig.Ctx(ctx).WhereIn(dao.SysConfig.Columns().ConfigId, postIds).Delete()
		utility.WriteErrLog(ctx, e, "删除参数配置表数据失败")
	})
	return
}

// 获取参数配置表
func (s *sSysConfig) GetSysConfig(ctx context.Context, req *v1.GetSysConfigReq) (res *v1.GetSysConfigRes, err error) {
	res = &v1.GetSysConfigRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var table *entity.SysConfig
		err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigId, req.ConfigId).Scan(&table)
		res.SysConfig = table
		utility.WriteErrLog(ctx, err, "删除参数配置表数据失败")
	})

	return
}
