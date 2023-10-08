package menu

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type menuController struct {
}

var Menu = menuController{}

func (c *menuController) GetMenuTreeSelect(ctx context.Context, req *v1.GetTreeSelectReq) (res *v1.GetTreeSelectRes, err error) {
	res = &v1.GetTreeSelectRes{}
	treeSelect, err := service.SysMenu().GetMenuTreeSelect(ctx)
	if err != nil {
		return
	}
	res.Tree = treeSelect
	return
}

// 获取菜单列表
func (c *menuController) GetMenuList(ctx context.Context, req *v1.GetMenuListReq) (res *v1.GetMenuListRes, err error) {
	res, err = service.SysMenu().GetMenuList(ctx, req)
	return
}

// 获取修改菜单
func (c *menuController) GetMenuUpdate(ctx context.Context, req *v1.GetMenuUpdateReq) (res *v1.GetMenuUpdateRes, err error) {
	res = &v1.GetMenuUpdateRes{}
	res.SysMenu, err = service.SysMenu().GetOneMenuById(ctx, req.MenuId)
	return
}

// 修改菜单
func (c *menuController) PutMenuUpdate(ctx context.Context, req *v1.PutMenuUpdateReq) (res *v1.PutMenuUpdateRes, err error) {
	res, err = service.SysMenu().UpdateMenu(ctx, req)
	return
}

// 添加菜单
func (c *menuController) Add(ctx context.Context, req *v1.PostMenuAddReq) (res *v1.PostMenuAddRes, err error) {
	err = service.SysMenu().Add(ctx, req)
	return
}

// 添加菜单
func (c *menuController) Delete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	err = service.SysMenu().Delete(ctx, req)
	return
}
