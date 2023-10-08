// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta     `orm:"table:sys_dept, do:true"`
	DeptId     interface{} // 部门id
	ParentId   interface{} // 父部门id
	Ancestors  interface{} // 祖级列表
	DeptName   interface{} // 部门名称
	OrderNum   interface{} // 显示顺序
	Leader     interface{} // 负责人
	Phone      interface{} // 联系电话
	Email      interface{} // 邮箱
	Status     interface{} // 部门状态（0正常 1停用）
	DelFlag    interface{} // 删除标志（0代表存在 2代表删除）
	CreateBy   interface{} // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   interface{} // 更新者
	UpdateTime *gtime.Time // 更新时间
}
