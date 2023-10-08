package sys_dept

import (
	"context"
	"fmt"
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
)

func init() {
	service.RegisterSysDept(New())
}

func New() *sSysDept {
	return &sSysDept{}
}

type sSysDept struct {
}

func (s *sSysDept) GetDeptById(ctx context.Context, id int64) (dept *entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//用户所属部门信息
		err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().DeptId, id).Scan(&dept)
		utility.WriteErrLog(ctx, err, "获取用户部门数据失败")
	})
	return
}

func (s *sSysDept) GetAllDeptList(ctx context.Context) (deptList map[int64]*entity.SysDept, err error) {
	cache := service.Cache().Driver()
	has, err := cache.Exists(ctx, consts.GetAllDeptListCache)
	if err != nil {
		return nil, err
	}
	if has != 0 {
		iDeptList, err := cache.Get(ctx, consts.GetAllDeptListCache)
		if err != nil {
			utility.WriteErrLog(ctx, err, "获取部门数据缓存失败")
		}
		err = iDeptList.Struct(&deptList)
		return deptList, err
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//deptList = make([]*entity.SysDept, )
		var deptRows []*entity.SysDept
		//字典数据表
		err = dao.SysDept.Ctx(ctx).Scan(&deptRows)
		deptList = make(map[int64]*entity.SysDept, 0)
		for _, v := range deptRows {
			deptList[v.DeptId] = v
		}
		cache.Set(ctx, consts.GetAllDeptListCache, deptList)
		utility.WriteErrLog(ctx, err, "获取部门数据失败")
	})
	return
}

func (s *sSysDept) GetAllDept(ctx context.Context) (deptList []*entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//字典数据表
		err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().Status, consts.SysDeptStatusOk).Scan(&deptList)
		utility.WriteErrLog(ctx, err, "获取部门数据失败")
	})
	return
}
func (s *sSysDept) GetDeptListByUid(ctx context.Context, uid int64) (deptList []*entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//是超级管理员
		if consts.ProAdminId == uid {
			deptList, err = s.GetAllDept(ctx)
			return
		}
		//获取用户角色
		roleList, err := service.SysRole().GetRolesByUid(ctx, uid)
		//字典数据表
		err = dao.SysDept.Ctx(ctx).As("d").Fields("d.*").LeftJoin(dao.SysRoleDept.Table()+" rd", "d.dept_id=rd.dept_id").WhereIn("rd.role_id", roleList.RoleIds).Distinct().Scan(&deptList)
		utility.WriteErrLog(ctx, err, "获取部门数据失败")
	})
	return
}

func (s *sSysDept) DeptTree(deptList []*entity.SysDept, pid int64) (tree []*model.SysDeptTreeRes) {
	tree = make([]*model.SysDeptTreeRes, 0)
	for _, v := range deptList {
		if v.ParentId == pid {
			//fmt.Printf("%+v\n", v)
			dd := &model.SysDeptTreeRes{Id: v.DeptId, Label: v.DeptName}
			child := s.DeptTree(deptList, v.DeptId)
			if child != nil {
				dd.Children = child
			}
			tree = append(tree, dd)
		}
	}
	return
}

func (s *sSysDept) GetDeptTree(ctx context.Context, uid int64) (deptTree []*model.SysDeptTreeRes, err error) {
	deptList, err := s.GetDeptListByUid(ctx, uid)
	if err != nil {
		return
	}
	//fmt.Printf("%+v\n", deptList)
	deptTree = s.DeptTree(deptList, 0)
	return
}

func (s *sSysDept) DeptTreeId(deptList []*entity.SysDept, deptIds *model.DeptIds, pid int64) (hasChild bool) {

	for _, v := range deptList {
		if v.ParentId == pid {
			child := s.DeptTreeId(deptList, deptIds, v.DeptId)
			if child {
				hasChild = child
				continue
			}
			deptIds.Ids = append(deptIds.Ids, v.DeptId)
		}
	}
	return
}
func (s *sSysDept) GetDeptId(ctx context.Context, uid, pid int64) (deptIds *model.DeptIds, err error) {
	deptIds = &model.DeptIds{}
	deptList, err := s.GetDeptListByUid(ctx, uid)
	if err != nil {
		return
	}
	//fmt.Printf("%+v\n", deptList)
	s.DeptTreeId(deptList, deptIds, pid)
	fmt.Printf("%+v\n", deptIds.Ids)
	if len(deptIds.Ids) < 1 {
		deptIds.Ids = append(deptIds.Ids, pid)
	}
	return
}

// 根据roleid 获取dept
func (s *sSysDept) GetDeptListByRoleId(ctx context.Context, roleId int64) (deptList []*entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//字典数据表
		err = dao.SysDept.Ctx(ctx).As("d").Fields("d.*").LeftJoin(dao.SysRoleDept.Table()+" rd", "d.dept_id=rd.dept_id").Where("rd.role_id", roleId).Distinct().Scan(&deptList)
		utility.WriteErrLog(ctx, err, "获取部门数据失败")
	})
	return
}

// 根据roleid 获取dept
func (s *sSysDept) GetDeptListByDeptIds(ctx context.Context, deptIds []int64) (deptList []*entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//字典数据表
		err = dao.SysDept.Ctx(ctx).As("d").WhereIn("dept_id", deptIds).Scan(&deptList)
		utility.WriteErrLog(ctx, err, "获取部门数据失败")
	})
	return
}

func (s *sSysDept) GetUpdateRoleDeptTree(ctx context.Context, req *v1.GetRoleDeptTreeReq) (res *v1.GetRoleDeptTreeRes, err error) {
	res = &v1.GetRoleDeptTreeRes{}
	uid := gconv.Int64(ctx.Value(consts.CtxAdminId))
	var deptList []*entity.SysDept
	if uid == consts.ProAdminId {
		deptList, err = s.GetAllDept(ctx)
		if err != nil {
			return
		}
	} else {
		deptList, err = s.GetDeptListByRoleId(ctx, req.RoleId)
		if err != nil {
			return
		}
	}
	res.CheckedKeys, err = service.SysRoleDept().GetDeptIdsByRoleid(ctx, req.RoleId)
	res.Depts = s.DeptTree(deptList, 0)
	return
}

func (s *sSysDept) HasChild(deptList map[int64]*entity.SysDept, deptId int64) bool {
	for _, dept := range deptList {
		if dept.ParentId == deptId {
			return true
		}
	}
	return false
}

// 部门列表
func (s *sSysDept) GetDeptList(ctx context.Context, req *v1.GetDeptListReq) (deptList *v1.GetDeptListRes, err error) {
	var list []*entity.SysDept
	deptList = &v1.GetDeptListRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDept.Ctx(ctx)
		if req.Status != "" {
			m = m.Where(dao.SysDept.Columns().Status, req.Status)
		}
		if req.DeptName != "" {
			m = m.WhereLike(dao.SysDept.Columns().DeptName, "%"+req.DeptName+"%")
		}
		err = m.Scan(&list)
		utility.WriteErrLog(ctx, err, "获取菜单失败")
	})
	deptList.List = list
	return
}

// 修改获取部门列表
func (s *sSysDept) GetDeptListUpdate(ctx context.Context, req *v1.GetDeptListUpdateReq) (deptList *v1.GetDeptListUpdateRes, err error) {
	var list []*entity.SysDept
	deptList = &v1.GetDeptListUpdateRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDept.Ctx(ctx).WhereNotIn(dao.SysDept.Columns().DeptId, req.DeptId).Scan(&list)
		utility.WriteErrLog(ctx, err, "获取菜单列表失败")
	})
	deptList.List = list
	return
}

// 删除数据
func (s sSysDept) Delete(ctx context.Context, req *v1.DeleteDeptReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysDept.Ctx(ctx).WherePri(req.DeptId).Update(do.SysDept{
			Status:     consts.SysDeptStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除部门数据失败")
	})
	return
}

// 修改数据
func (s sSysDept) Update(ctx context.Context, req *v1.PutDeptUpdateReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		ancestors, err := s.GetAncestors(ctx, req.ParentId)
		if err != nil {
			return
		}
		_, e := dao.SysDept.Ctx(ctx).WherePri(req.DeptId).Update(do.SysDept{
			ParentId:   req.ParentId,
			DeptName:   req.DeptName,
			Ancestors:  ancestors,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			Leader:     req.Leader,
			Phone:      req.Phone,
			Email:      req.Email,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "修改部门数据失败")
	})
	return
}

// 修改数据
func (s sSysDept) Add(ctx context.Context, req *v1.PostDeptAddReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		ancestors, err := s.GetAncestors(ctx, req.ParentId)
		if err != nil {
			return
		}
		_, e := dao.SysDept.Ctx(ctx).Data(do.SysDept{
			ParentId:   req.ParentId,
			DeptName:   req.DeptName,
			Ancestors:  ancestors,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			Leader:     req.Leader,
			Phone:      req.Phone,
			Email:      req.Email,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
			CreateTime: gtime.Now(),
			CreateBy:   adminName,
		}).Insert()
		utility.WriteErrLog(ctx, e, "添加部门数据失败")
	})
	return
}
func (s sSysDept) GetAncestors(ctx context.Context, deptId int64) (ancestors string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		parentDept, err := s.GetDeptById(ctx, deptId)
		if err != nil {
			return
		}
		ancestors = fmt.Sprintf("s%,s%", parentDept.Ancestors, deptId)
		fmt.Printf(ancestors, "kkkkkkk")
	})
	return
}
