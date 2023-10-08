package sys_role

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/consts"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/do"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
	"go-vue-admin/utility/lib"
)

func init() {
	service.RegisterSysRole(New())
}

func New() *sSysRole {
	return &sSysRole{}
}

type sSysRole struct {
}

func (s *sSysRole) GetNomalRole(ctx context.Context) (roles []*entity.SysRole, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleId+" >", consts.ProAdminRoleId).Scan(&roles)
		utility.WriteErrLog(ctx, err, "获取用户角色数据失败")
	})
	return
}

func (s *sSysRole) GetRoleByUid(ctx context.Context, uid int64) (roles []*entity.SysRole, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		err = dao.SysUserRole.Ctx(ctx).As("ur").Fields("r.*").LeftJoin(dao.SysRole.Table()+" r", "ur.role_id=r.role_id").Where("ur.user_id", uid).Scan(&roles)
		utility.WriteErrLog(ctx, err, "获取用户角色数据失败")
	})
	return
}

func (s *sSysRole) GetRolesByUid(ctx context.Context, uid int64) (rolesList *model.SysRolesRes, err error) {
	var roles []*entity.SysRole
	roles, err = s.GetRoleByUid(ctx, uid)
	if err != nil {
		return
	}
	rolesList = new(model.SysRolesRes)
	rolesList.SysRole = roles
	for _, v := range roles {
		rolesList.RoleIds = append(rolesList.RoleIds, v.RoleId)
		rolesList.Roles = append(rolesList.Roles, v.RoleKey)
	}
	//fmt.Printf("角色%+v\n", rolesList)
	return
}

func (s *sSysRole) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (roleList *v1.GetRoleListRes, err error) {
	roleList = &v1.GetRoleListRes{}
	userId := gconv.Int64(ctx.Value(consts.CtxAdminId))
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	err = g.Try(ctx, func(ctx context.Context) {
		var roles []*entity.SysRole
		m := dao.SysRole.Ctx(ctx)
		//状态
		if req.Status != "" {
			m = m.Where(dao.SysRole.Columns().Status, req.Status)
		}
		//时间
		if len(req.Params) > 0 {
			m = m.WhereBetween(dao.SysRole.Columns().CreateTime, req.Params["beginTime"], req.Params["endTime"])
		}

		//角色名
		if req.RoleName != "" {
			m = m.WhereLike(dao.SysRole.Columns().RoleName, "%"+req.RoleName+"%")
		}
		//权限符
		if req.RoleKey != "" {
			m = m.WhereLike(dao.SysRole.Columns().RoleKey, "%"+req.RoleKey+"%")
		}
		if userId != consts.ProAdminId {
			//获取用户信息
			roleIds, err := service.SysUserRole().GetRoleIdByUid(ctx, userId)
			utility.WriteErrLog(ctx, err, "获取角色id数据失败")
			m = m.WhereIn(dao.SysRole.Columns().RoleId, roleIds)
		}
		roleList.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&roles)
		roleRows := make([]*model.RoleList, len(roles))
		for k, value := range roles {
			ul := &model.RoleList{}
			if value.RoleId == consts.ProAdminRoleId {
				ul.Admin = true
			}
			ul.SysRole = value
			roleRows[k] = ul
		}
		roleList.Rows = roleRows
	})
	return
}
func (s *sSysRole) IsRoleNameAvailable(ctx context.Context, roleName string, roleId int64) (bool, error) {
	var sysRole *entity.SysRole
	err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleName, roleName).Scan(&sysRole)
	if err != nil {
		return false, err
	}
	if sysRole == nil {
		return true, nil
	}
	if roleId > 0 && roleId == sysRole.RoleId {
		return true, nil
	}
	return false, nil

}

func (s *sSysRole) Add(ctx context.Context, req *v1.PostRoleReq) (res *v1.PostRoleRes, err error) {
	var (
		available bool
	)
	if req.RoleName == "" {
		err = gerror.Newf(`角色名不能为空`)
		return
	}
	available, err = s.IsRoleNameAvailable(ctx, req.RoleName, 0)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`角色名 "%s" 已经存在请使用其他`, req.RoleName)
		return
	}
	if req.RoleKey == "" {
		err = gerror.Newf(`权限字符不能为空`)
		return
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加角色信息
			roleId, err := dao.SysRole.Ctx(ctx).TX(tx).InsertAndGetId(do.SysRole{
				RoleKey:           req.RoleKey,
				RoleName:          req.RoleName,
				RoleSort:          req.RoleSort,
				DeptCheckStrictly: req.DeptCheckStrictly,
				MenuCheckStrictly: req.MenuCheckStrictly,
				Remark:            req.Remark,
				Status:            req.Status,
				CreateTime:        gtime.Now(),
				CreateBy:          adminName,
				UpdateTime:        gtime.Now(),
				UpdateBy:          adminName,
			})
			utility.WriteErrLog(ctx, err, "添加角色数据失败")
			err = service.SysRoleMenu().AddRoleMenus(ctx, tx, roleId, req.MenuIds)
			utility.WriteErrLog(ctx, err, "设置角色菜单失败")
		})
		return err
	})
	return
}

func (s *sSysRole) GetRoleUpdate(ctx context.Context, req *v1.GetRoleUpdateReq) (res *v1.GetRoleUpdateRes, err error) {
	res = &v1.GetRoleUpdateRes{}
	var role *entity.SysRole
	err = g.Try(ctx, func(ctx context.Context) {
		//添加角色信息
		e := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleId, req.RoleId).Scan(&role)
		utility.WriteErrLog(ctx, e, "获取角色失败")
	})
	res.SysRole = role
	return

}

// 获取菜单数
func (s *sSysRole) GetRoleUpdateTreeSelect(ctx context.Context, req *v1.GetRoleUpdateTreeSelectReq) (res *v1.GetRoleUpdateTreeSelectRes, err error) {
	res = &v1.GetRoleUpdateTreeSelectRes{}
	res.Menus, err = service.SysMenu().GetMenuTreeSelect(ctx)
	if err != nil {
		return
	}
	res.CheckedKeys, err = service.SysRoleMenu().GetMenuIdsByRoleId(ctx, req.RoleId)
	return
}

func (s *sSysRole) Update(ctx context.Context, req *v1.PutRoleUpdateReq) (res *v1.PutRoleUpdateRes, err error) {
	var (
		available bool
	)
	if req.RoleName == "" {
		err = gerror.Newf(`角色名不能为空`)
		return
	}
	available, err = s.IsRoleNameAvailable(ctx, req.RoleName, req.RoleId)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`角色名 "%s" 已经存在请使用其他`, req.RoleName)
		return
	}
	if req.RoleKey == "" {
		err = gerror.Newf(`权限字符不能为空`)
		return
	}
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加角色信息
			_, err := dao.SysRole.Ctx(ctx).TX(tx).WherePri(req.RoleId).Update(do.SysRole{
				RoleKey:           req.RoleKey,
				RoleName:          req.RoleName,
				RoleSort:          req.RoleSort,
				DeptCheckStrictly: req.DeptCheckStrictly,
				MenuCheckStrictly: req.MenuCheckStrictly,
				Remark:            req.Remark,
				Status:            req.Status,
				UpdateTime:        gtime.Now(),
				UpdateBy:          adminName,
			})
			utility.WriteErrLog(ctx, err, "修改角色数据失败")
			err = service.SysRoleMenu().AddRoleMenus(ctx, tx, req.RoleId, req.MenuIds)
			utility.WriteErrLog(ctx, err, "设置角色菜单失败")
		})
		return err
	})
	return
}

func (s *sSysRole) UpdateDataScope(ctx context.Context, req *v1.PutRoleDataScopeReq) (res *v1.PutRoleDataScopeRes, err error) {

	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//添加角色信息
			_, err := dao.SysRole.Ctx(ctx).TX(tx).WherePri(req.RoleId).Update(do.SysRole{
				DeptCheckStrictly: req.DeptCheckStrictly,
				Remark:            req.Remark,
				DataScope:         req.DataScope,
				UpdateTime:        gtime.Now(),
				UpdateBy:          adminName,
			})
			utility.WriteErrLog(ctx, err, "修改角色数据失败")
			//全部数据权限
			if req.DataScope == consts.SysRoleDataScopeCustom {
				err = service.SysRoleDept().AddRoleDepts(ctx, tx, req.RoleId, req.DeptIds)
				utility.WriteErrLog(ctx, err, "设置角色失败")
			}
		})
		return err
	})
	return
}

// 更角色改状态
func (s *sSysRole) ChangeStatus(ctx context.Context, req *v1.ChangeStatusRoleReq) (res *v1.ChangeStatusRoleRes, err error) {

	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		//删除用户信息
		_, e := dao.SysRole.Ctx(ctx).WherePri(req.RoleId).Update(do.SysUser{
			Status:     req.Status,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "更改角色数据失败")
	})
	return
}

// 更角色改状态
func (s *sSysRole) Delete(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {

	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		roleIds := lib.ParamStrToSlice(req.RoleId, ",")
		//删除用户信息
		_, e := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().RoleId, roleIds).Update(do.SysUser{
			Status:     consts.SysRoleStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除角色数据失败")
	})
	return
}

// 获取分配用户
func (s *sSysRole) GetRoleAuthUser(ctx context.Context, req *v1.GetRoleAuthUserReq) (res *v1.GetRoleAuthUserRes, err error) {
	res = &v1.GetRoleAuthUserRes{}
	var users []*entity.SysUser
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//用户用户信息
		m := dao.SysUser.Ctx(ctx).As("u").Fields("u.*").LeftJoin(dao.SysUserRole.Table()+" ur", "u.user_id=ur.user_id").Where("ur.role_id", req.RoleId)
		if req.UserName != "" {
			m = m.WhereLike("u.user_name", "%"+req.UserName+"%")
		}
		if req.Phonenumber != "" {
			m = m.WhereLike("u.phonenumber", "%"+req.Phonenumber+"%")
		}
		err = m.Page(req.PageNum, req.PageSize).Scan(&users)
		utility.WriteErrLog(ctx, err, "获取角色用户数据失败")
	})
	res.Users = users
	return
}

// 获取添加分配用户
func (s *sSysRole) GetRoleAddAuthUser(ctx context.Context, req *v1.GetRoleAddAuthUserReq) (res *v1.GetRoleAddAuthUserRes, err error) {
	res = &v1.GetRoleAddAuthUserRes{}
	var users []*entity.SysUser
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	err = g.Try(ctx, func(ctx context.Context) {
		userIds, err := service.SysUserRole().GetUserIdByRoleId(ctx, req.RoleId)
		if err != nil {
			return
		}
		//用户用户信息
		m := dao.SysUser.Ctx(ctx)
		if req.UserName != "" {
			m = m.WhereLike("user_name", "%"+req.UserName+"%")
		}
		if req.Phonenumber != "" {
			m = m.WhereLike("phonenumber", "%"+req.Phonenumber+"%")
		}
		fmt.Println(userIds, "uuuuuu")
		if len(userIds) > 0 {
			m = m.WhereNotIn(dao.SysUser.Columns().UserId, userIds)
		}
		res.Total, err = m.Count()
		if err != nil {
			return
		}
		err = m.Page(req.PageNum, req.PageSize).Scan(&users)
		utility.WriteErrLog(ctx, err, "获取角色用户数据失败")
	})
	res.Users = users
	return
}

// 添加分配用户
func (s *sSysRole) PutRoleAddAuthUser(ctx context.Context, req *v1.PutRoleAddAuthUserReq) (res *v1.PutRoleAddAuthUserRes, err error) {
	userIds := lib.ParamStrToSlice(req.UserIds, ",")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//设置角色
			err = service.SysUserRole().AddRoleUsers(ctx, tx, req.RoleId, userIds)
			utility.WriteErrLog(ctx, err, "设置用户角色失败")
		})
		return err
	})
	return
}

// 取消分配用户
func (s *sSysRole) PutRoleCancelAuthUser(ctx context.Context, req *v1.PutRoleCancelAuthUserReq) (res *v1.PutRoleCancelAuthUserRes, err error) {
	userIds := []int64{req.UserId}
	fmt.Println(userIds, "dddddd")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//设置角色
			err = service.SysUserRole().CancelRoleUsers(ctx, tx, req.RoleId, userIds)
			utility.WriteErrLog(ctx, err, "取消分配用户失败")
		})
		return err
	})
	return
}

// 批量取消分配用户
func (s *sSysRole) PutRoleCancelAllAuthUser(ctx context.Context, req *v1.PutRoleCancelAllAuthUserReq) (res *v1.PutRoleCancelAllAuthUserRes, err error) {
	userIds := lib.ParamStrToSlice(req.UserIds, ",")
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			//设置角色
			err = service.SysUserRole().CancelRoleUsers(ctx, tx, req.RoleId, userIds)
			utility.WriteErrLog(ctx, err, "批量取消分配用户失败")
		})
		return err
	})
	return
}
