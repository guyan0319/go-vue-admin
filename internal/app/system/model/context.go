package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
}

type ContextUser struct {
	Id       int64  // User ID.
	UserName string // User UserName.
	Nickname string // User nickname.
}
