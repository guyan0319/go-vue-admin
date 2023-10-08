// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	RoleId            int64       `json:"roleId"            description:"角色ID"`
	RoleName          string      `json:"roleName"          description:"角色名称"`
	RoleKey           string      `json:"roleKey"           description:"角色权限字符串"`
	RoleSort          int         `json:"roleSort"          description:"显示顺序"`
	DataScope         string      `json:"dataScope"         description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	MenuCheckStrictly int         `json:"menuCheckStrictly" description:"菜单树选择项是否关联显示"`
	DeptCheckStrictly int         `json:"deptCheckStrictly" description:"部门树选择项是否关联显示"`
	Status            string      `json:"status"            description:"角色状态（0正常 1停用）"`
	DelFlag           string      `json:"delFlag"           description:"删除标志（0代表存在 2代表删除）"`
	CreateBy          string      `json:"createBy"          description:"创建者"`
	CreateTime        *gtime.Time `json:"createTime"        description:"创建时间"`
	UpdateBy          string      `json:"updateBy"          description:"更新者"`
	UpdateTime        *gtime.Time `json:"updateTime"        description:"更新时间"`
	Remark            string      `json:"remark"            description:"备注"`
}
