package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SignUp is the API for user sign up.
func (c *Controller) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.SysUser().Create(ctx, model.UserCreateInput{
		UserName: req.UserName,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *Controller) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.SysUser().IsSignedIn(ctx),
	}
	return
}

// CheckUserName checks and returns whether the user UserName is available.
func (c *Controller) CheckUserName(ctx context.Context, req *v1.CheckUserNameReq) (res *v1.CheckUserNameRes, err error) {
	available, err := service.SysUser().IsUserNameAvailable(ctx, req.UserName)
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
	available, err := service.SysUser().IsNicknameAvailable(ctx, req.Nickname, 0)
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

// GetInfo returns the user info.
func (c *Controller) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	var (
		user       = &model.SysUserInfoRes{}
		sysUserRes = &model.SysUserRes{}
	)
	userId := gconv.Int64(ctx.Value(consts.CtxAdminId))
	sysUserRes, err = service.SysUser().GetUserById(ctx, userId)
	user.SysUserRes = sysUserRes

	//获取用户角色
	roleIds := []int64{}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	for _, v := range rolesList.SysRole {
		sysRoleRes := &model.SysRoleRes{}
		if v.RoleId == consts.ProAdminRoleId {
			sysRoleRes.Admin = true
		}
		sysRoleRes.SysRole = v
		user.Roles = append(user.Roles, sysRoleRes)
		roleIds = append(roleIds, v.RoleId)
		//fmt.Printf("%+v\n", v)
	}
	//当前用户是否超级管理员
	if user.UserId == consts.ProAdminId {
		user.Admin = true
	}
	//获取权限
	permlist := &model.RolePerm{}
	permlist, err = service.SysMenu().GetPermByRoleids(ctx, roleIds)
	//fmt.Printf("权限%+v\n", permlist)
	//设置每个角色权限
	for _, val := range user.Roles {
		val.Permissions = permlist.MapPerms[val.RoleId]
	}
	//fmt.Printf("%+v\n", roles)
	//fmt.Println(userId)
	res = &v1.GetInfoRes{
		User:        user,
		Roles:       rolesList.Roles,
		Permissions: permlist.AllPerm,
	}
	return
}

// Profile returns the user profile.
func (c *Controller) GetRoutters(ctx context.Context, req *v1.GetRoutersReq) (res *v1.GetRoutersRes, err error) {

	userId := gconv.Int64(ctx.Value(consts.CtxAdminId))
	//获取roleids
	roleIds := []int64{}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, v := range rolesList.SysRole {
		roleIds = append(roleIds, v.RoleId)
		//fmt.Printf("%+v\n", v)
	}
	//获取菜单
	menuList, err := service.SysMenu().GetRoutersByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	//menuTree
	res = &v1.GetRoutersRes{
		MenuList: menuList,
		//SysUser: service.User().GetProfile(ctx),
	}
	return
}
func (c *Controller) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	res, err = service.SysUser().GetUserListByDeptId(ctx, req)
	return
}
func (c *Controller) GetAddUser(ctx context.Context, req *v1.GetAddUserReq) (res *v1.GetAddUserRes, err error) {
	res = &v1.GetAddUserRes{}
	res.Roles, err = service.SysRole().GetNomalRole(ctx)
	if err != nil {
		return
	}
	if req.UserId > 0 {
		user := &model.UserList{}
		user.SysUser, err = service.SysUser().GetOneUserById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		user.Dept, err = service.SysDept().GetDeptById(ctx, user.DeptId)
		if err != nil {
			return nil, err
		}
		res.User = user
		res.PostIds, _ = service.SysUserPost().GetPostIdByUid(ctx, req.UserId)
		res.RoleIds, _ = service.SysUserRole().GetRoleIdByUid(ctx, req.UserId)
	}
	res.Posts, err = service.SysPost().GetAllPostByStatus(ctx, consts.SysPostStatusOk)
	return
}
func (c *Controller) Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error) {
	res, err = service.SysUser().Add(ctx, req)
	return
}
func (c *Controller) Udate(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error) {
	res, err = service.SysUser().Udate(ctx, req)
	return
}
func (c *Controller) Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	res, err = service.SysUser().Delete(ctx, req)
	return
}
func (c *Controller) ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error) {
	res, err = service.SysUser().ChangeStatus(ctx, req)
	return
}
func (c *Controller) ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error) {
	res, err = service.SysUser().ResetPWD(ctx, req)
	return
}
func (c *Controller) GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error) {
	res, err = service.SysUser().GetAuthRole(ctx, req)
	return
}
func (c *Controller) PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error) {
	res, err = service.SysUser().PutAuthRole(ctx, req)
	return
}
