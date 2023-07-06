package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gmode"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SignUp is the API for user sign up.
func (c *Controller) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		UserName: req.UserName,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// Login is the API for user sign in.
func (c *Controller) Login(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	//判断验证码是否正确
	debug := gmode.IsDevelop()
	fmt.Println(req)
	if !debug {
		if !service.Captcha().VerifyString(req.VerifyKey, req.Code) {
			err = gerror.New("验证码输入错误")
			return
		}
	}

	err = service.User().Login(ctx, model.UserSignInInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *Controller) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

// LogOut is the API for user sign out.
func (c *Controller) LogOut(ctx context.Context, req *v1.LogOutReq) (res *v1.LogOutRes, err error) {
	err = service.User().LogOut(ctx)
	return
}

// CheckUserName checks and returns whether the user UserName is available.
func (c *Controller) CheckUserName(ctx context.Context, req *v1.CheckUserNameReq) (res *v1.CheckUserNameRes, err error) {
	available, err := service.User().IsUserNameAvailable(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`UserName "%s" is already token by others`, req.UserName)
	}
	return
}

// CheckNickName checks and returns whether the user nickname is available.
func (c *Controller) CheckNickName(ctx context.Context, req *v1.CheckNickNameReq) (res *v1.CheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

// Profile returns the user profile.
func (c *Controller) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{
		//SysUser: service.User().GetProfile(ctx),
	}
	return
}
