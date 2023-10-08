package role

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type roleController struct {
}

var Role = roleController{}

func (c *roleController) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error) {
	res, err = service.SysRole().GetRoleList(ctx, req)
	return
}

func (c *roleController) Add(ctx context.Context, req *v1.PostRoleReq) (res *v1.PostRoleRes, err error) {
	res, err = service.SysRole().Add(ctx, req)
	return
}

func (c *roleController) GetRoleUpdate(ctx context.Context, req *v1.GetRoleUpdateReq) (res *v1.GetRoleUpdateRes, err error) {
	res, err = service.SysRole().GetRoleUpdate(ctx, req)
	return
}

func (c *roleController) GetRoleUpdateTreeSelect(ctx context.Context, req *v1.GetRoleUpdateTreeSelectReq) (res *v1.GetRoleUpdateTreeSelectRes, err error) {
	res, err = service.SysRole().GetRoleUpdateTreeSelect(ctx, req)
	return
}
func (c *roleController) Update(ctx context.Context, req *v1.PutRoleUpdateReq) (res *v1.PutRoleUpdateRes, err error) {
	res, err = service.SysRole().Update(ctx, req)
	return
}
func (c *roleController) GetUpdateRoleDeptTree(ctx context.Context, req *v1.GetRoleDeptTreeReq) (res *v1.GetRoleDeptTreeRes, err error) {
	res, err = service.SysDept().GetUpdateRoleDeptTree(ctx, req)
	return
}
func (c *roleController) UpdateRoleDataScope(ctx context.Context, req *v1.PutRoleDataScopeReq) (res *v1.PutRoleDataScopeRes, err error) {
	res, err = service.SysRole().UpdateDataScope(ctx, req)
	return
}
func (c *roleController) ChangeStatusRole(ctx context.Context, req *v1.ChangeStatusRoleReq) (res *v1.ChangeStatusRoleRes, err error) {
	res, err = service.SysRole().ChangeStatus(ctx, req)
	return
}
func (c *roleController) Delete(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	res, err = service.SysRole().Delete(ctx, req)
	return
}
func (c *roleController) GetRoleAuthUser(ctx context.Context, req *v1.GetRoleAuthUserReq) (res *v1.GetRoleAuthUserRes, err error) {
	res, err = service.SysRole().GetRoleAuthUser(ctx, req)
	return
}
func (c *roleController) GetRoleAddAuthUser(ctx context.Context, req *v1.GetRoleAddAuthUserReq) (res *v1.GetRoleAddAuthUserRes, err error) {
	res, err = service.SysRole().GetRoleAddAuthUser(ctx, req)
	return
}
func (c *roleController) PutRoleAddAuthUser(ctx context.Context, req *v1.PutRoleAddAuthUserReq) (res *v1.PutRoleAddAuthUserRes, err error) {
	res, err = service.SysRole().PutRoleAddAuthUser(ctx, req)
	return
}
func (c *roleController) PutRoleCancelAuthUser(ctx context.Context, req *v1.PutRoleCancelAuthUserReq) (res *v1.PutRoleCancelAuthUserRes, err error) {
	res, err = service.SysRole().PutRoleCancelAuthUser(ctx, req)
	return
}
func (c *roleController) PutRoleCancelAllAuthUser(ctx context.Context, req *v1.PutRoleCancelAllAuthUserReq) (res *v1.PutRoleCancelAllAuthUserRes, err error) {
	res, err = service.SysRole().PutRoleCancelAllAuthUser(ctx, req)
	return
}
