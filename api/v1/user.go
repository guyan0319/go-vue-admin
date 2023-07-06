package v1

import (
	"github.com/gogf/gf/v2/frame/g"
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
	g.Meta   `path:"/login" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	UserName string `p:"userName" v:"required"`
	Password string `p:"password" v:"required"`
	Code string `p:"code" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}
type SignInRes struct{}

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

type LogOutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type LogOutRes struct{}
