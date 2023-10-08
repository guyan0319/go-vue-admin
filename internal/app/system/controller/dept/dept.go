package dept

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/service"
)

type deptController struct {
}

var Dept = deptController{}

func (d *deptController) GetDeptTree(ctx context.Context, req *v1.GetDeptTreeReq) (res *v1.GetDeptTreeRes, err error) {
	uid := gconv.Int64(ctx.Value(consts.CtxAdminId))
	dictData, err := service.SysDept().GetDeptTree(ctx, uid)
	res = &v1.GetDeptTreeRes{
		DeptTree: dictData,
	}
	return
}

// 获取部门列表
func (c *deptController) GetDeptList(ctx context.Context, req *v1.GetDeptListReq) (res *v1.GetDeptListRes, err error) {
	res, err = service.SysDept().GetDeptList(ctx, req)
	return
}

// 获取所有部门列表
func (c *deptController) GetDeptlistUpdate(ctx context.Context, req *v1.GetDeptListUpdateReq) (res *v1.GetDeptListUpdateRes, err error) {
	res, err = service.SysDept().GetDeptListUpdate(ctx, req)
	return
}

// 获取修改部门数据
func (c *deptController) GetDeptUpdate(ctx context.Context, req *v1.GetDeptUpdateReq) (res *v1.GetDeptUpdateRes, err error) {
	res = &v1.GetDeptUpdateRes{}
	res.SysDept, err = service.SysDept().GetDeptById(ctx, req.DeptId)
	return
}

// 删除部门
func (c *deptController) Delete(ctx context.Context, req *v1.DeleteDeptReq) (res *v1.DeleteDeptRes, err error) {
	err = service.SysDept().Delete(ctx, req)
	return
}

// 修改部门
func (c *deptController) Update(ctx context.Context, req *v1.PutDeptUpdateReq) (res *v1.PutDeptUpdateRes, err error) {
	err = service.SysDept().Update(ctx, req)
	return
}

// 添加部门
func (c *deptController) Add(ctx context.Context, req *v1.PostDeptAddReq) (res *v1.PostDeptAddRes, err error) {
	err = service.SysDept().Add(ctx, req)
	return
}
