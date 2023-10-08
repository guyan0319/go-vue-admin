package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetRoleListReq struct {
	g.Meta `path:"/system/role/list" method:"get" tags:"RoleService" summary:"current role list"`
	common.PageReq
	RoleName string `p:"roleName"`
	RoleKey  string `p:"roleKey"`
	Status   string `p:"status"`
}
type GetRoleListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.RoleList `json:"rows"`
	Total  int               `json:"total"`
}
type PostRoleReq struct {
	g.Meta            `path:"/system/role" method:"post" tags:"RoleService" summary:"add role "`
	RoleName          string  `p:"roleName"  v:"required"`
	RoleKey           string  `p:"roleKey"  v:"required"`
	Status            string  `p:"status"  v:"required"`
	RoleSort          string  `p:"roleSort"`
	Remark            string  `p:"remark"`
	DeptCheckStrictly int     `p:"deptCheckStrictly"`
	MenuCheckStrictly int     `p:"menuCheckStrictly"`
	DeptIds           []int64 `p:"deptIds"`
	MenuIds           []int64 `p:"menuIds"`
}
type PostRoleRes struct {
	g.Meta `mime:"application/json"`
}
type GetRoleUpdateReq struct {
	g.Meta `path:"/system/role/{roleId}" method:"get" tags:"RoleService" summary:"update role "`
	RoleId int64 `p:"roleId"  v:"required"`
}
type GetRoleUpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysRole
}
type PutRoleUpdateReq struct {
	g.Meta            `path:"/system/role" method:"put" tags:"RoleService" summary:"update role "`
	RoleId            int64   `p:"roleId"  v:"required"`
	RoleName          string  `p:"roleName"  v:"required"`
	RoleKey           string  `p:"roleKey"  v:"required"`
	Status            string  `p:"status"  v:"required"`
	RoleSort          string  `p:"roleSort"`
	Remark            string  `p:"remark"`
	DeptCheckStrictly int     `p:"deptCheckStrictly"`
	MenuCheckStrictly int     `p:"menuCheckStrictly"`
	DeptIds           []int64 `p:"deptIds"`
	MenuIds           []int64 `p:"menuIds"`
}
type PutRoleUpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysRole
}

// 配置分配数据权限
type PutRoleDataScopeReq struct {
	g.Meta            `path:"/system/role/dataScope" method:"put" tags:"RoleService" summary:"update role "`
	RoleId            int64   `p:"roleId"  v:"required"`
	DataScope         string  `p:"dataScope"  v:"required"`
	Remark            string  `p:"remark"`
	DeptCheckStrictly int     `p:"deptCheckStrictly"`
	MenuCheckStrictly int     `p:"menuCheckStrictly"`
	DeptIds           []int64 `p:"deptIds"`
}
type PutRoleDataScopeRes struct {
	g.Meta `mime:"application/json"`
}
type GetRoleUpdateTreeSelectReq struct {
	g.Meta `path:"/system/menu/roleMenuTreeselect/{roleId}" method:"get" tags:"RoleService" summary:"update role "`
	RoleId int64 `p:"roleId"  v:"required"`
}
type GetRoleUpdateTreeSelectRes struct {
	g.Meta      `mime:"application/json"`
	Menus       []*model.SysMenuTreeRes `json:"menus"`
	CheckedKeys []int64                 `json:"checkedKeys"`
}

// 分配数据权限
type GetRoleDeptTreeReq struct {
	g.Meta `path:"/system/role/deptTree/{roleId}" method:"get" tags:"DeptService" summary:"current Data"`
	RoleId int64 `p:"roleId"  v:"required"`
}
type GetRoleDeptTreeRes struct {
	g.Meta      `mime:"application/json"`
	Depts       []*model.SysDeptTreeRes `json:"depts"`
	CheckedKeys []int64                 `json:"checkedKeys"`
}
type ChangeStatusRoleReq struct {
	g.Meta `path:"/system/role/changeStatus" method:"PUT" tags:"DeptService" summary:"current Data"`
	RoleId int64  `p:"roleId"  v:"required"`
	Status string `p:"status"  v:"required"`
}
type ChangeStatusRoleRes struct {
	g.Meta `mime:"application/json"`
}
type DeleteRoleReq struct {
	g.Meta `path:"/system/role/{roleId}" method:"delete" tags:"DeptService" summary:"current Data"`
	RoleId string `p:"roleId"  v:"required"`
}
type DeleteRoleRes struct {
	g.Meta `mime:"application/json"`
}

// 分配用户
type GetRoleAuthUserReq struct {
	g.Meta      `path:"/system/role/authUser/allocatedList" method:"get" tags:"User Service" summary:"role user Data"`
	RoleId      int64  `p:"roleId"  v:"required"`
	UserName    string `p:"userName" `
	Phonenumber string `p:"phonenumber" `
	common.PageReq
}
type GetRoleAuthUserRes struct {
	g.Meta `mime:"application/json"`
	Users  []*entity.SysUser `json:"users"`
}

// 取消授权
type PutRoleCancelAuthUserReq struct {
	g.Meta `path:"/system/role/authUser/cancel" method:"put" tags:"UserService" summary:"put Data"`
	RoleId int64 `p:"roleId"  v:"required"`
	UserId int64 `p:"userId"  v:"required"`
}
type PutRoleCancelAuthUserRes struct {
	g.Meta `mime:"application/json"`
}
type PutRoleCancelAllAuthUserReq struct {
	g.Meta  `path:"/system/role/authUser/cancelAll" method:"put" tags:"UserService" summary:"put Data"`
	RoleId  int64  `p:"roleId"  v:"required"`
	UserIds string `p:"userIds"  v:"required"`
}
type PutRoleCancelAllAuthUserRes struct {
	g.Meta `mime:"application/json"`
}

type GetRoleAddAuthUserReq struct {
	g.Meta      `path:"/system/role/authUser/unallocatedList" method:"get" tags:"UserService" summary:"get Data"`
	RoleId      int64  `p:"roleId"  v:"required"`
	UserName    string `p:"userName" `
	Phonenumber string `p:"phonenumber" `
	common.PageReq
}
type GetRoleAddAuthUserRes struct {
	g.Meta `mime:"application/json"`
	Users  []*entity.SysUser `json:"users"`
	Total  int               `json:"total"`
}
type PutRoleAddAuthUserReq struct {
	g.Meta  `path:"/system/role/authUser/selectAll" method:"put" tags:"UserService" summary:"get Data"`
	RoleId  int64  `p:"roleId"  v:"required"`
	UserIds string `p:"userIds"  v:"required" `
}
type PutRoleAddAuthUserRes struct {
	g.Meta `mime:"application/json"`
}
