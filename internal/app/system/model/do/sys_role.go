// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta            `orm:"table:sys_role, do:true"`
	RoleId            interface{} // 角色ID
	RoleName          interface{} // 角色名称
	RoleKey           interface{} // 角色权限字符串
	RoleSort          interface{} // 显示顺序
	DataScope         interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly interface{} // 菜单树选择项是否关联显示
	DeptCheckStrictly interface{} // 部门树选择项是否关联显示
	Status            interface{} // 角色状态（0正常 1停用）
	DelFlag           interface{} // 删除标志（0代表存在 2代表删除）
	CreateBy          interface{} // 创建者
	CreateTime        *gtime.Time // 创建时间
	UpdateBy          interface{} // 更新者
	UpdateTime        *gtime.Time // 更新时间
	Remark            interface{} // 备注
}
