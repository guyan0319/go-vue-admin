package sys_role_dept

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
	service.RegisterSysRoleDept(New())
}

func New() *sSysRoleDept {
	return &sSysRoleDept{}
}

type sSysRoleDept struct {
}

func (s *sSysRoleDept) GetDeptIdsByRoleid(ctx context.Context, roleId int64) (deptIds []int64, err error) {
	var roleDept []*entity.SysRoleDept
	err = dao.SysRoleDept.Ctx(ctx).Fields(dao.SysRoleDept.Columns().DeptId).Where(dao.SysRoleDept.Columns().RoleId, roleId).Scan(&roleDept)
	deptList, err := service.SysDept().GetAllDeptList(ctx)
	if err != nil {
		return
	}
	utility.WriteErrLog(ctx, err, "获取角色菜单失败")
	for _, v := range roleDept {
		//过滤有子部门节点
		has := service.SysDept().HasChild(deptList, v.DeptId)
		if has {
			continue
		}
		deptIds = append(deptIds, v.DeptId)
	}
	return
}

func (s *sSysRoleDept) AddRoleDepts(ctx context.Context, tx gdb.TX, roleId int64, DeptIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//删除旧角色菜单
		_, err = dao.SysRoleDept.Ctx(ctx).TX(tx).Where(dao.SysRoleDept.Columns().RoleId, roleId).Delete()
		utility.WriteErrLog(ctx, err, "删除角色部门失败")
		if len(DeptIds) == 0 {
			return
		}
		//添加角色菜单信息
		data := g.List{}
		for _, v := range DeptIds {
			data = append(data, g.Map{
				dao.SysRoleDept.Columns().RoleId: roleId,
				dao.SysRoleDept.Columns().DeptId: v,
			})
		}
		_, err = dao.SysRoleDept.Ctx(ctx).TX(tx).Data(data).Insert()
		utility.WriteErrLog(ctx, err, "添加角色部门失败")
	})
	return
}
