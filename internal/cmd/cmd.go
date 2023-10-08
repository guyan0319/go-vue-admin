package cmd

import (
	"context"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/controller/common"
	"go-vue-admin/internal/app/system/controller/dept"
	"go-vue-admin/internal/app/system/controller/dict"
	"go-vue-admin/internal/app/system/controller/menu"
	"go-vue-admin/internal/app/system/controller/post"
	"go-vue-admin/internal/app/system/controller/role"
	"go-vue-admin/internal/app/system/controller/sys_config"
	"go-vue-admin/internal/app/system/controller/sys_dict_type"
	"go-vue-admin/internal/app/system/controller/table"
	"go-vue-admin/internal/app/system/controller/user"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gfToken, err := GetGtoken(ctx)
			if err != nil {
				return err
			}
			s := g.Server()
			s.Use(service.Middleware().HandlerResponse) //返回数据处理
			//s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {

				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				// Register route handlers. 不需要登录
				group.Bind(
					common.Captcha,
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					err = gfToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Middleware(service.Middleware().Auth)
					//group.ALLMap(g.Map{
					//	"/user/profile": user.New().Profile,
					//})

					group.Bind(
						user.New(),
						dict.Dict,
						sys_dict_type.DictType,
						dept.Dept,
						role.Role,
						menu.Menu,
						post.Post,
						table.Table,
						sys_config.SysConfig,
					)

				})
			})
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}
func GetGtoken(ctx context.Context) (gfToken *gtoken.GfToken, err error) {
	// 启动gtoken
	gfToken = &gtoken.GfToken{
		ServerName:       consts.ServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  LoginFunc,
		LogoutPath:       "post:/logout",
		AuthPaths:        g.SliceStr{"/user", "/getInfo"}, // 这里是按照前缀拦截，拦截/user /user/list /user/add ...
		AuthExcludePaths: g.SliceStr{},                    // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc:    AuthAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	err = gfToken.Start()
	return
}
func LoginFunc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("userName").String()
	password := r.Get("password").String()
	code := r.Get("code").String()
	verifyKey := r.Get("verifyKey").String()

	//判断验证码是否正确
	//设置为生成环境
	//gmode.SetProduct()
	debug := gmode.IsDevelop()
	fmt.Println(debug)
	if !debug {
		if code == "" || verifyKey == "" {
			r.Response.WriteJson(gtoken.Fail(consts.ErrLoginCodeFailMsg))
			r.ExitAll()
		}
		if !service.Captcha().VerifyString(verifyKey, code) {
			r.Response.WriteJson(gtoken.Fail(consts.ErrLoginCodeFailMsg))
			r.ExitAll()
		}
	}
	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	ctx := context.TODO()
	var users *entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: username,
		//Password: password,
	}).Scan(&users)
	if err != nil || users == nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return fmt.Sprintf("%s%d", consts.GTokenAdminPrefix, users.UserId), users
}

func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var users entity.SysUser
	//fmt.Println("bbb保存",respData)
	err := gconv.Struct(respData.GetString("data"), &users)
	if err != nil {
		fmt.Println(err)
		r.Response.WriteJson(gtoken.Unauthorized(consts.ErrAuthFailMsg, nil))
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	//r.SetCtxVar(consts.CtxAdminId, 1)
	r.SetCtxVar(consts.CtxAdminId, users.UserId)
	r.SetCtxVar(consts.CtxAdminName, users.UserName)
	//fmt.Println("aa保存",user.UserId)
	r.Middleware.Next()
}
