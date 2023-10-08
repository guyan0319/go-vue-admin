// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type (
	ISysDept interface {
		GetDeptById(ctx context.Context, id int64) (dept *entity.SysDept, err error)
		GetAllDeptList(ctx context.Context) (deptList map[int64]*entity.SysDept, err error)
		GetAllDept(ctx context.Context) (deptList []*entity.SysDept, err error)
		GetDeptListByUid(ctx context.Context, uid int64) (deptList []*entity.SysDept, err error)
		DeptTree(deptList []*entity.SysDept, pid int64) (tree []*model.SysDeptTreeRes)
		GetDeptTree(ctx context.Context, uid int64) (deptTree []*model.SysDeptTreeRes, err error)
		DeptTreeId(deptList []*entity.SysDept, deptIds *model.DeptIds, pid int64) (hasChild bool)
		GetDeptId(ctx context.Context, uid, pid int64) (deptIds *model.DeptIds, err error)
		// 根据roleid 获取dept
		GetDeptListByRoleId(ctx context.Context, roleId int64) (deptList []*entity.SysDept, err error)
		// 根据roleid 获取dept
		GetDeptListByDeptIds(ctx context.Context, deptIds []int64) (deptList []*entity.SysDept, err error)
		GetUpdateRoleDeptTree(ctx context.Context, req *v1.GetRoleDeptTreeReq) (res *v1.GetRoleDeptTreeRes, err error)
		HasChild(deptList map[int64]*entity.SysDept, deptId int64) bool
		// 部门列表
		GetDeptList(ctx context.Context, req *v1.GetDeptListReq) (deptList *v1.GetDeptListRes, err error)
		// 修改获取部门列表
		GetDeptListUpdate(ctx context.Context, req *v1.GetDeptListUpdateReq) (deptList *v1.GetDeptListUpdateRes, err error)
		// 删除数据
		Delete(ctx context.Context, req *v1.DeleteDeptReq) (err error)
		// 修改数据
		Update(ctx context.Context, req *v1.PutDeptUpdateReq) (err error)
		// 修改数据
		Add(ctx context.Context, req *v1.PostDeptAddReq) (err error)
		GetAncestors(ctx context.Context, deptId int64) (ancestors string, err error)
	}
)

var (
	localSysDept ISysDept
)

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}
