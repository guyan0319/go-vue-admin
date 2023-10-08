package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type ProfileRes struct {
	*entity.SysUser
}

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	UserName  string `p:"userName" v:"required|length:6,16"`
	Password  string `p:"password" v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
}
type SignUpRes struct{}

type SignInReq struct {
	g.Meta    `path:"/login" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	UserName  string `p:"userName" v:"required"`
	Password  string `p:"password" v:"required"`
	Code      string `p:"code" v:"required#验证码不能为空"`
	VerifyKey string `p:"verifyKey"`
}
type SignInRes struct {
	Token string `p:"token"`
}

type CheckUserNameReq struct {
	g.Meta   `path:"/user/check-UserName" method:"post" tags:"UserService" summary:"Check UserName available"`
	UserName string `v:"required"`
}
type CheckUserNameRes struct{}

type CheckNickNameReq struct {
	g.Meta   `path:"/user/check-nick-name" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type CheckNickNameRes struct{}

type IsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type IsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

//type LogOutReq struct {
//	g.Meta `path:"/logout" method:"post" tags:"UserService" summary:"Sign out current user"`
//}
//type LogOutRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/getInfo" method:"get" tags:"UserService" summary:"current user info"`
}

type GetInfoRes struct {
	g.Meta      `mime:"application/json"`
	Permissions []string              `json:"permissions"`
	Roles       []string              `json:"roles"`
	User        *model.SysUserInfoRes `json:"user"`
}
type GetRoutersReq struct {
	g.Meta `path:"/getRouters" method:"get" tags:"UserService" summary:"current user routers"`
}
type GetRoutersRes struct {
	g.Meta   `mime:"application/json"`
	MenuList []*model.UserMenuRes `json:"menuList"`
}
type GetAddUserReq struct {
	g.Meta `path:"/system/user" method:"get" tags:"UserService" summary:"add user"`
	UserId int64 `p:"userId"`
}
type GetAddUserRes struct {
	g.Meta  `mime:"application/json"`
	Roles   []*entity.SysRole `json:"roles"`
	Posts   []*entity.SysPost `json:"posts"`
	RoleIds []int64           `json:"roleIds"`
	User    *model.UserList   `json:"user"`
	PostIds []int64           `json:"postIds"`
}
type PostAddUserReq struct {
	g.Meta      `path:"/system/user" method:"post" tags:"UserService" summary:"add user"`
	DeptId      int64   `p:"deptId"`
	UserName    string  `p:"userName"`
	NickName    string  `p:"nickName"`
	Password    string  `p:"password"`
	Phonenumber string  `p:"phonenumber"`
	Email       string  `p:"email"`
	Sex         string  `p:"sex"`
	Status      string  `p:"status"`
	Remark      string  `p:"remark"`
	PostIds     []int64 `p:"postIds"`
	RoleIds     []int64 `p:"roleIds"`
}
type PostAddUserRes struct {
	g.Meta `mime:"application/json"`
}
type DeleteUserReq struct {
	g.Meta `path:"/system/user/{userId}" method:"DELETE" tags:"UserService" summary:"delete user"`
	UserId string `p:"userId" v:"required"`
}
type DeleteUserRes struct {
	g.Meta `mime:"application/json"`
}
type ChangeUserStatusReq struct {
	g.Meta `path:"/system/user/changeStatus" method:"PUT" tags:"UserService" summary:"change user status"`
	UserId int64 `p:"userId" v:"required"`
	Status int64 `p:"status" v:"required"`
}
type ChangeUserStatusRes struct {
	g.Meta `mime:"application/json"`
}
type ResetPwdUserReq struct {
	g.Meta   `path:"/system/user/resetPwd" method:"PUT" tags:"UserService" summary:"change user status"`
	UserId   int64 `p:"userId" v:"required"`
	Password int64 `p:"password" v:"required"`
}
type ResetPwdUserRes struct {
	g.Meta `mime:"application/json"`
}

type PutUpdateUserReq struct {
	g.Meta      `path:"/system/user" method:"put" tags:"UserService" summary:"update user"`
	UserId      int64   `p:"userId" v:"required"`
	DeptId      int64   `p:"deptId"`
	NickName    string  `p:"nickName"`
	Phonenumber string  `p:"phonenumber"`
	Email       string  `p:"email"`
	Sex         string  `p:"sex"`
	Status      string  `p:"status"`
	Remark      string  `p:"remark"`
	PostIds     []int64 `p:"postIds"`
	RoleIds     []int64 `p:"roleIds"`
}
type PutUpdateUserRes struct {
	g.Meta `mime:"application/json"`
}

type GetUserListReq struct {
	g.Meta `path:"/system/user/list" method:"get" tags:"UserService" summary:"current user list"`
	common.PageReq
	DeptId      int64  `p:"deptId"`
	UserName    string `p:"userName"`
	Phonenumber string `p:"phonenumber"`
	Status      string `p:"status"`
}
type GetUserListRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.UserList `json:"rows"`
	Total  int               `json:"total"`
}
type GetAuthRoleUserReq struct {
	g.Meta `path:"/system/user/authRole/{userId}" method:"get" tags:"UserService" summary:"get user role"`
	UserId int64 `p:"userId" v:"required"`
	common.PageReq
}
type GetAuthRoleUserRes struct {
	g.Meta `mime:"application/json"`
	Roles  []*entity.SysRole   `json:"roles"`
	User   *model.AuthRoleUser `json:"user"`
}
type PutAuthRoleUserReq struct {
	g.Meta  `path:"/system/user/authRole" method:"put" tags:"UserService" summary:"update user role"`
	UserId  int64  `p:"userId" v:"required"`
	RoleIds string `p:"roleIds" v:"required"`
}
type PutAuthRoleUserRes struct {
	g.Meta `mime:"application/json"`
}
