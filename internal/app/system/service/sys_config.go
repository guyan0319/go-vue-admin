// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
)

type (
	ISysConfig interface {
		// 获取参数配置表列表
		GetSysConfigList(ctx context.Context, req *v1.GetSysConfigListReq) (res *v1.GetSysConfigListRes, err error)
		// 添加参数配置表
		Add(ctx context.Context, req *v1.PostSysConfigReq) (res *v1.PostSysConfigRes, err error)
		// 修改参数配置表
		Update(ctx context.Context, req *v1.PutSysConfigReq) (res *v1.PutSysConfigRes, err error)
		// 删除参数配置表
		Delete(ctx context.Context, req *v1.DeleteSysConfigReq) (res *v1.DeleteSysConfigRes, err error)
		// 获取参数配置表
		GetSysConfig(ctx context.Context, req *v1.GetSysConfigReq) (res *v1.GetSysConfigRes, err error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
