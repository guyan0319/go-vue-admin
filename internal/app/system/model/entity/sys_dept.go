// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId     int64       `json:"deptId"     description:"部门id"`
	ParentId   int64       `json:"parentId"   description:"父部门id"`
	Ancestors  string      `json:"ancestors"  description:"祖级列表"`
	DeptName   string      `json:"deptName"   description:"部门名称"`
	OrderNum   int         `json:"orderNum"   description:"显示顺序"`
	Leader     string      `json:"leader"     description:"负责人"`
	Phone      string      `json:"phone"      description:"联系电话"`
	Email      string      `json:"email"      description:"邮箱"`
	Status     string      `json:"status"     description:"部门状态（0正常 1停用）"`
	DelFlag    string      `json:"delFlag"    description:"删除标志（0代表存在 2代表删除）"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
}
