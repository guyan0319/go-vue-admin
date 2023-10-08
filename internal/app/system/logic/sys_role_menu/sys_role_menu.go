package sys_role_menu

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
)

func init() {
	service.RegisterSysRoleMenu(New())
}

func New() *sSysRoleMenu {
	return &sSysRoleMenu{}
}

type sSysRoleMenu struct {
}

func (s *sSysRoleMenu) AddRoleMenus(ctx context.Context, tx gdb.TX, roleId int64, MenuIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧角色菜单
		_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Delete()
		utility.WriteErrLog(ctx, err, "删除角色菜单失败")
		if len(MenuIds) == 0 {
			return
		}
		//添加角色菜单信息
		data := g.List{}
		for _, v := range MenuIds {
			data = append(data, g.Map{
				dao.SysRoleMenu.Columns().RoleId: roleId,
				dao.SysRoleMenu.Columns().MenuId: v,
			})
		}
		_, err = dao.SysRoleMenu.Ctx(ctx).TX(tx).Data(data).Insert()
		utility.WriteErrLog(ctx, err, "添加角色菜单失败")
	})
	return
}

func (s *sSysRoleMenu) GetMenuIdsByRoleId(ctx context.Context, roleId int64) (menuIds []int64, err error) {
	var roleMenu []*entity.SysRoleMenu
	err = dao.SysRoleMenu.Ctx(ctx).Fields(dao.SysRoleMenu.Columns().MenuId).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Scan(&roleMenu)
	utility.WriteErrLog(ctx, err, "获取角色菜单失败")
	for _, v := range roleMenu {
		menuIds = append(menuIds, v.MenuId)
	}
	return
}
