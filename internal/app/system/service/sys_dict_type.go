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
	ISysDictType interface {
		// 字典类型表
		GetDictTypeOption(ctx context.Context, req *v1.GetDictTypeOptionSelectReq) (res *v1.GetDictTypeOptionSelectRes, err error)
		// 获取字典类型表列表
		GetSysDictTypeList(ctx context.Context, req *v1.GetSysDictTypeListReq) (res *v1.GetSysDictTypeListRes, err error)
		// 添加字典类型表
		Add(ctx context.Context, req *v1.PostSysDictTypeReq) (res *v1.PostSysDictTypeRes, err error)
		// 修改字典类型表
		Update(ctx context.Context, req *v1.PutSysDictTypeReq) (res *v1.PutSysDictTypeRes, err error)
		// 删除字典类型表
		Delete(ctx context.Context, req *v1.DeleteSysDictTypeReq) (res *v1.DeleteSysDictTypeRes, err error)
		// 获取字典类型表
		GetSysDictType(ctx context.Context, req *v1.GetSysDictTypeReq) (res *v1.GetSysDictTypeRes, err error)
	}
)

var (
	localSysDictType ISysDictType
)

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}
