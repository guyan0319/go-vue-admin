package sys_dict_type

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type sysDictTypeController struct {
}

var DictType = sysDictTypeController{}

func (c *sysDictTypeController) GetDictTypeOption(ctx context.Context, req *v1.GetDictTypeOptionSelectReq) (res *v1.GetDictTypeOptionSelectRes, err error) {
	res, err = service.SysDictType().GetDictTypeOption(ctx, req)
	return
}
func (s *sysDictTypeController) GetSysDictTypeList(ctx context.Context, req *v1.GetSysDictTypeListReq) (res *v1.GetSysDictTypeListRes, err error) {
	res, err = service.SysDictType().GetSysDictTypeList(ctx, req)
	return
}
func (s *sysDictTypeController) Add(ctx context.Context, req *v1.PostSysDictTypeReq) (res *v1.PostSysDictTypeRes, err error) {
	res, err = service.SysDictType().Add(ctx, req)
	return
}

func (s *sysDictTypeController) Update(ctx context.Context, req *v1.PutSysDictTypeReq) (res *v1.PutSysDictTypeRes, err error) {
	res, err = service.SysDictType().Update(ctx, req)
	return
}

func (s *sysDictTypeController) Delete(ctx context.Context, req *v1.DeleteSysDictTypeReq) (res *v1.DeleteSysDictTypeRes, err error) {
	res, err = service.SysDictType().Delete(ctx, req)
	return
}

func (s *sysDictTypeController) GetSysDictType(ctx context.Context, req *v1.GetSysDictTypeReq) (res *v1.GetSysDictTypeRes, err error) {
	res, err = service.SysDictType().GetSysDictType(ctx, req)
	return
}
