package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"go-vue-admin/internal/app/system/model/entity"
)

type UserCreateInput struct {
	UserName string
	Password string
	Nickname string
}

type UserSignInInput struct {
	UserName string
	Password string
}
type SysUserRes struct {
	UserId      int64       `json:"userId"      description:"用户ID"`
	DeptId      int64       `json:"deptId"      description:"部门ID"`
	UserName    string      `json:"userName"    description:"用户账号"`
	NickName    string      `json:"nickName"    description:"用户昵称"`
	Email       string      `json:"email"       description:"用户邮箱"`
	Phonenumber string      `json:"phonenumber" description:"手机号码"`
	Sex         string      `json:"sex"         description:"用户性别（0男 1女 2未知）"`
	Avatar      string      `json:"avatar"      description:"头像地址"`
	Status      string      `json:"status"      description:"帐号状态（0正常 1停用）"`
	DelFlag     string      `json:"delFlag"     description:"删除标志（0代表存在 2代表删除）"`
	LoginIp     string      `json:"loginIp"     description:"最后登录IP"`
	LoginDate   *gtime.Time `json:"loginDate"   description:"最后登录时间"`
	CreateBy    string      `json:"createBy"    description:"创建者"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
	Remark      string      `json:"remark"      description:"备注"`
}
type SysUserInfoRes struct {
	*SysUserRes
	Dept    *SysDeptSRes  `json:"dept" `
	Roles   []*SysRoleRes `json:"roles" `
	RoleIds []int64       `json:"roleIds"    `
	PostIds []int64       `json:"postIds"    `
	Admin   bool          `json:"admin"            description:"是否是admin"`
}
type UserList struct {
	*entity.SysUser
	Dept *entity.SysDept `json:"dept" `
}
type AuthRoleUser struct {
	*entity.SysUser
	Admin bool            `json:"admin"            description:"是否是admin"`
	Dept  *entity.SysDept `json:"dept" `
}
