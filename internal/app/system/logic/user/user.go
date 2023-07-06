package user

import (
	"context"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses UserName as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.UserName
	}
	var (
		available bool
	)
	// UserName checks.
	available, err = s.IsUserNameAvailable(ctx, in.UserName)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`UserName "%s" is already token by others`, in.UserName)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.SysUser.Ctx(ctx).Data(do.SysUser{
			//UserName: in.UserName,
			Password: in.Password,
			//Nickname: in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// Login creates session for given user account.
func (s *sUser) Login(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: in.UserName,
		//Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`UserName or Password not correct`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id: user.UserId,
		UserName: user.UserName,
		Nickname: user.UserName,
	})
	return nil
}

// LogOut removes the session for current signed-in user.
func (s *sUser) LogOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsUserNameAvailable checks and returns given UserName is available for signing up.
func (s *sUser) IsUserNameAvailable(ctx context.Context, UserName string) (bool, error) {
	count, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		//UserName: UserName,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.SysUser {
	return service.Session().GetUser(ctx)
}
