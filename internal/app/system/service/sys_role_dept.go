// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysRoleDept interface {
		GetDeptIdsByRoleid(ctx context.Context, roleId int64) (deptIds []int64, err error)
		AddRoleDepts(ctx context.Context, tx gdb.TX, roleId int64, DeptIds []int64) (err error)
	}
)

var (
	localSysRoleDept ISysRoleDept
)

func SysRoleDept() ISysRoleDept {
	if localSysRoleDept == nil {
		panic("implement not found for interface ISysRoleDept, forgot register?")
	}
	return localSysRoleDept
}

func RegisterSysRoleDept(i ISysRoleDept) {
	localSysRoleDept = i
}
