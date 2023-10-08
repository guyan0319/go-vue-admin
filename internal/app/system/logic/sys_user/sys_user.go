package sys_user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/utility"
	"go-vue-admin/utility/lib"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/service"
)

type (
	sSysUser struct{}
)

func init() {
	service.RegisterSysUser(New())
}

func New() service.ISysUser {
	return &sSysUser{}
}

// Create creates user account.
func (s *sSysUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {

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
	available, err = s.IsNicknameAvailable(ctx, in.Nickname, 0)
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

func (s *sSysUser) Add(ctx context.Context, req *v1.PostAddUserReq) (res *v1.PostAddUserRes, err error) {
	var (
		available bool
	)
	if req.NickName == "" {
		err = gerror.Newf(`昵称不能为空`)
		return
	}
	available, err = s.IsNicknameAvailable(ctx, req.NickName, 0)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`昵称 "%s" 已经存在请使用其他`, req.NickName)
		return
	}

	if req.UserName == "" {
		err = gerror.Newf(`用户名不能为空`)
		return
	}
	available, err = s.IsUserNameAvailable(ctx, req.UserName)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`用户名 "%s" 已经存在请使用其他`, req.UserName)
		return
	}
	if req.Password == "" || (len(req.Password) < 5 && len(req.Password) > 20) {
		err = gerror.Newf(`密码不能为空,且用户密码长度必须介于 5 和 20 之间`)
		return
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加用户信息
			userId, e := dao.SysUser.Ctx(ctx).TX(tx).InsertAndGetId(do.SysUser{
				UserName:    req.UserName,
				Password:    gsha1.Encrypt(req.Password),
				NickName:    req.NickName,
				DeptId:      req.DeptId,
				Phonenumber: req.Phonenumber,
				Email:       req.Email,
				Sex:         req.Sex,
				Status:      req.Status,
				Remark:      req.Remark,
				LoginIp:     lib.GetClientIp(ctx),
				CreateTime:  gtime.Now(),
				CreateBy:    adminName,
				LoginDate:   gtime.Now(),
				UpdateTime:  gtime.Now(),
				UpdateBy:    adminName,
			})
			utility.WriteErrLog(ctx, e, "获取用户数据失败")
			//设置用户角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, userId, req.RoleIds)
			utility.WriteErrLog(ctx, e, "设置用户角色失败")
			err = service.SysUserPost().AddUserPosts(ctx, tx, userId, req.PostIds)
			utility.WriteErrLog(ctx, e, "设置用户职位失败")
		})
		return err
	})
	return
}

func (s *sSysUser) Udate(ctx context.Context, req *v1.PutUpdateUserReq) (res *v1.PutUpdateUserRes, err error) {
	var (
		available bool
	)
	if req.NickName == "" {
		err = gerror.Newf(`昵称不能为空`)
		return
	}
	available, err = s.IsNicknameAvailable(ctx, req.NickName, req.UserId)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`昵称 "%s" 已经存在请使用其他`, req.NickName)
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加用户信息
			_, e := dao.SysUser.Ctx(ctx).TX(tx).WherePri(req.UserId).Update(do.SysUser{
				NickName:    req.NickName,
				DeptId:      req.DeptId,
				Phonenumber: req.Phonenumber,
				Email:       req.Email,
				Sex:         req.Sex,
				Status:      req.Status,
				Remark:      req.Remark,
				UpdateTime:  gtime.Now(),
				UpdateBy:    adminName,
			})
			utility.WriteErrLog(ctx, e, "获取用户数据失败")
			//设置用户角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, req.UserId, req.RoleIds)
			utility.WriteErrLog(ctx, e, "设置用户角色失败")
			err = service.SysUserPost().AddUserPosts(ctx, tx, req.UserId, req.PostIds)
			utility.WriteErrLog(ctx, e, "设置用户职位失败")
		})
		return err
	})
	return
}

// 假删除 支持批量删除
func (s *sSysUser) Delete(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		userIds := lib.ParamStrToSlice(req.UserId, ",")
		//删除用户信息
		_, e := dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().UserId, userIds).Update(do.SysUser{
			Status:     consts.SysUserStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除用户数据失败")
	})
	return
}

// 更改状态
func (s *sSysUser) ChangeStatus(ctx context.Context, req *v1.ChangeUserStatusReq) (res *v1.ChangeUserStatusRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//删除用户信息
		_, e := dao.SysUser.Ctx(ctx).WherePri(req.UserId).Update(do.SysUser{
			Status:     req.Status,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "更改用户数据失败")
	})
	return
}

// 更改密码
func (s *sSysUser) ResetPWD(ctx context.Context, req *v1.ResetPwdUserReq) (res *v1.ResetPwdUserRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//删除用户信息
		_, e := dao.SysUser.Ctx(ctx).WherePri(req.UserId).Update(do.SysUser{
			Password:   gsha1.Encrypt(req.Password),
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "更改用户数据失败")
	})
	return
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sSysUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// IsUserNameAvailable checks and returns given UserName is available for signing up.
func (s *sSysUser) IsUserNameAvailable(ctx context.Context, UserName string) (bool, error) {
	count, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: UserName,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sSysUser) IsNicknameAvailable(ctx context.Context, nickname string, userId int64) (bool, error) {
	sysUser := entity.SysUser{}
	err := dao.SysUser.Ctx(ctx).Where(do.SysUser{
		UserName: nickname,
	}).Scan(&sysUser)
	if err != nil {
		return false, err
	}
	if userId > 0 && userId != sysUser.UserId {
		return false, nil
	}
	if sysUser.UserId > 0 {
		return false, nil
	}
	return true, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sSysUser) GetProfile(ctx context.Context) *entity.SysUser {
	return service.Session().GetUser(ctx)
}
func (s *sSysUser) GetUserById(ctx context.Context, userId int64) (user *model.SysUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Scan(&user)
		utility.WriteErrLog(ctx, err, "获取用户数据失败")
	})
	return
}
func (s *sSysUser) GetOneUserById(ctx context.Context, id int64) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, id).Scan(&user)
		utility.WriteErrLog(ctx, err, "获取用户数据失败")
	})
	return
}

func (s *sSysUser) GetUserListByDeptId(ctx context.Context, req *v1.GetUserListReq) (userList *v1.GetUserListRes, err error) {
	userList = &v1.GetUserListRes{}
	uid := gconv.Int64(ctx.Value(consts.CtxAdminId))
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//user := []*model.UserList{}
		user := []*entity.SysUser{}
		//用户用户信息
		m := dao.SysUser.Ctx(ctx)
		if req.DeptId > 0 || uid != consts.ProAdminId {
			//获取部门列表
			deptIds := &model.DeptIds{}
			deptIds, err = service.SysDept().GetDeptId(ctx, uid, req.DeptId)
			m = m.WhereIn(dao.SysUser.Columns().DeptId, deptIds.Ids)
		}
		//手机号
		if req.Phonenumber != "" {
			m = m.Where(dao.SysUser.Columns().Phonenumber, req.Phonenumber)
		}
		//状态
		if req.Status != "" {
			m = m.Where(dao.SysUser.Columns().Status, req.Status)
		}
		//时间
		if len(req.Params) > 0 {
			//fmt.Println(req.Params["beginTime"])
			m = m.WhereBetween(dao.SysUser.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}
		//用户名不为空
		if req.UserName != "" {
			m = m.WhereLike(dao.SysUser.Columns().UserName, "%"+req.UserName+"%")
		}
		userList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&user)
		//fmt.Println(user)
		//获取所有部门数据
		deptList, err := service.SysDept().GetAllDeptList(ctx)
		if err != nil {
			return
		}
		userRows := make([]*model.UserList, len(user))
		for k, value := range user {
			ul := &model.UserList{}
			ul.SysUser = value
			ul.Dept = deptList[value.DeptId]
			userRows[k] = ul
		}
		userList.Rows = userRows
		utility.WriteErrLog(ctx, err, "获取用户数据失败")
	})

	return
}

func (s *sSysUser) GetAuthRole(ctx context.Context, req *v1.GetAuthRoleUserReq) (res *v1.GetAuthRoleUserRes, err error) {
	res = &v1.GetAuthRoleUserRes{}
	var user *entity.SysUser
	var roleList []*entity.SysRole
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	user, err = s.GetOneUserById(ctx, req.UserId)
	if err != nil {
		return
	}
	authRoleUser := &model.AuthRoleUser{}
	authRoleUser.SysUser = user
	if req.UserId == consts.ProAdminId {
		authRoleUser.Admin = true
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysRole.Ctx(ctx).As("r").Fields("r.*").LeftJoin(dao.SysUserRole.Table()+" ur", "r.role_id=ur.role_id").Where("ur.user_id", req.UserId).Page(req.PageNum, req.PageSize).Scan(&roleList)
		utility.WriteErrLog(ctx, err, "获取用户角色数据失败")
	})
	res.Roles = roleList
	res.User = authRoleUser
	return
}
func (s *sSysUser) PutAuthRole(ctx context.Context, req *v1.PutAuthRoleUserReq) (res *v1.PutAuthRoleUserRes, err error) {
	roleIds := lib.ParamStrToSlice(req.RoleIds, ",")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//设置角色
			err = service.SysUserRole().AddUserRoles(ctx, tx, req.UserId, roleIds)
			utility.WriteErrLog(ctx, err, "设置用户角色失败")
		})
		return err
	})
	return
}
