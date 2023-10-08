// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISysMenu interface {
		// 根据rold_id获取perms
		GetPermByRoleids(ctx context.Context, roleIds []int64) (permList *model.RolePerm, err error)
		// 根据rold_id获取menu
		GetMenuByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error)
		GetMenuTreeByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error)
		GetRoutersByRoleids(ctx context.Context, roleIds []int64) (userMenuRes []*model.UserMenuRes, err error)
		GetMenuTree(menuList []*entity.SysMenu, pid int64) (userMenuRes []*model.UserMenuRes)
		// 检测路由权限
		CheckUrlPerms(r *ghttp.Request) bool
		GetPermsUrlByUserId(ctx context.Context, userId int64) (perms map[string]bool)
		GetMenuTreeSelect(ctx context.Context) (menuTree []*model.SysMenuTreeRes, err error)
		MenuTreeSelect(menuList []*entity.SysMenu, pid int64) (tree []*model.SysMenuTreeRes)
		// 菜单列表
		GetMenuList(ctx context.Context, req *v1.GetMenuListReq) (menuList *v1.GetMenuListRes, err error)
		// 获取单条数据
		GetOneMenuById(ctx context.Context, menuId int64) (menu *entity.SysMenu, err error)
		// 修改数据
		UpdateMenu(ctx context.Context, req *v1.PutMenuUpdateReq) (res *v1.PutMenuUpdateRes, err error)
		// 初始化apipath
		InitApiPath(ctx context.Context) (err error)
		// 添加数据
		Add(ctx context.Context, req *v1.PostMenuAddReq) (err error)
		// 删除数据
		Delete(ctx context.Context, req *v1.MenuDeleteReq) (err error)
	}
)

var (
	localSysMenu ISysMenu
)

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}
