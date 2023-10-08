package sys_menu

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
	"golang.org/x/exp/maps"
	"strings"
)

func init() {
	service.RegisterSysMenu(New())
}

func New() *sSysMenu {
	return &sSysMenu{}
}

type sSysMenu struct {
}

//特殊路由

var SpecialApiPath = map[string]bool{
	//user
	"post/user/sign-up":                 true,
	"get/getInfo":                       true,
	"get/getRouters":                    true,
	"put/system/user/changeStatus":      true,
	"get/system/user/authRole/{userId}": true,
	"put/system/user/authRole":          true,
	//系统工具
	"get/tool/gen/db/list":                        true,
	"get/tool/gen/initTable":                      true,
	"get/system/menu/roleMenuTreeselect/{roleId}": true,
	//菜单
	"get/system/menu/treeselect": true,
	//字典
	"get/system/dict/type/optionselect": true,
	"get/system/dict/type":              true,
	//部门
	"get/system/user/deptTree":              true,
	"get/system/dept/list/exclude/{deptId}": true,
}

// 根据rold_id获取perms
func (s *sSysMenu) GetPermByRoleids(ctx context.Context, roleIds []int64) (permList *model.RolePerm, err error) {
	var mapPerms map[int64][]string
	mapPerms = make(map[int64][]string, 0)
	permList = new(model.RolePerm)
	//fmt.Printf("角色id%+v\n", roleIds)
	if lib.InSliceInt64(consts.ProAdminRoleId, &roleIds) { //如果包含超级管理员
		permList.AllPerm = []string{"*:*:*"}
		mapPerms[consts.ProAdminRoleId] = []string{}
		permList.MapPerms = mapPerms
		//fmt.Println("超级管理")
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		perms := []model.PermsData{}
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("rm.role_id,m.perms").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).Scan(&perms)
		utility.WriteErrLog(ctx, err, "获取用户角色数据失败")
		for _, v := range perms {
			if v.Perms == "" {
				continue
			}
			mapPerms[v.RoleId] = append(mapPerms[v.RoleId], v.Perms)
			if !lib.InSliceString(v.Perms, &permList.AllPerm) {
				permList.AllPerm = append(permList.AllPerm, v.Perms)
			}
		}
		permList.MapPerms = mapPerms

	})
	return
}

// 根据rold_id获取menu
func (s *sSysMenu) GetMenuByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error) {
	//fmt.Printf("菜单%+v\n", roleIds)
	menuType := []string{"M", "C"}
	if lib.InSliceInt64(consts.ProAdminRoleId, &roleIds) { //包含超级管理员角色，返回所有的菜单
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, consts.SysMenuStatusOk).WhereIn(dao.SysMenu.Columns().MenuType, menuType).OrderAsc("parent_id").Scan(&menu)
		utility.WriteErrLog(ctx, err, "获取用户菜单数据失败")
		//fmt.Printf("%+v\n", menu)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("m.*").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).WhereIn("m.menu_type", menuType).OrderAsc("m.parent_id").Scan(&menu)
		utility.WriteErrLog(ctx, err, "获取用户菜单数据失败")
	})
	return
}
func (s *sSysMenu) GetMenuTreeByRoleids(ctx context.Context, roleIds []int64) (menu []*entity.SysMenu, err error) {
	//fmt.Printf("菜单%+v\n", roleIds)
	if lib.InSliceInt64(consts.ProAdminRoleId, &roleIds) { //包含超级管理员角色，返回所有的菜单
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, consts.SysMenuStatusOk).OrderAsc("parent_id").Scan(&menu)
		utility.WriteErrLog(ctx, err, "获取用户菜单数据失败")
		//fmt.Printf("%+v\n", menu)
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//权限
		err = dao.SysRoleMenu.Ctx(ctx).As("rm").Fields("m.*").LeftJoin(dao.SysMenu.Table()+" m", "rm.menu_id=m.menu_id").WhereIn("rm.role_id", roleIds).Where("m.status", consts.SysMenuStatusOk).OrderAsc("m.parent_id").Scan(&menu)
		utility.WriteErrLog(ctx, err, "获取用户菜单数据失败")
	})
	return
}

func (s *sSysMenu) GetRoutersByRoleids(ctx context.Context, roleIds []int64) (userMenuRes []*model.UserMenuRes, err error) {
	menuList, err := s.GetMenuByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	userMenuRes = s.GetMenuTree(menuList, 0)
	return
}
func (s *sSysMenu) GetMenuTree(menuList []*entity.SysMenu, pid int64) (userMenuRes []*model.UserMenuRes) {
	tree := make([]*model.UserMenuRes, 0)
	for _, v := range menuList {
		if v.ParentId == pid {
			m := &model.UserMenuRes{}
			meta := &model.MenuMeta{}
			meta.Icon = v.Icon
			if v.IsFrame == 0 {
				meta.Link = v.Perms
			}
			if v.IsCache == 1 {
				meta.NoCache = true
			}
			meta.Title = v.MenuName
			m.Meta = meta
			m.Component = v.Component

			//m.AlwaysShow
			if v.Visible == "0" {
				m.Hidden = false
			}
			m.Path = v.Path
			if v.MenuType == "M" {
				m.Redirect = "noRedirect"
				m.Component = "Layout"
				m.Path = "/" + m.Path
				//只有是目录才设置为true
				m.AlwaysShow = true
			}
			m.Name = strings.ReplaceAll(v.Path, "/", "")
			m.Name = lib.StrFirstToUpper(m.Name)
			child := s.GetMenuTree(menuList, v.MenuId)
			if child != nil {
				m.Children = child
			}
			tree = append(tree, m)
		}
	}
	return tree
}

// 检测路由权限
func (s *sSysMenu) CheckUrlPerms(r *ghttp.Request) bool {
	//fmt.Println(strings.ToLower(r.Router.Method))
	//fmt.Println(r.Router.Uri)
	url := strings.ToLower(r.Router.Method) + r.Router.Uri
	//获取userid
	userId := gconv.Int64(r.Context().Value(consts.CtxAdminId))
	//如果超级管理员，全部放行
	if userId == consts.ProAdminId {
		return true
	}
	perms := s.GetPermsUrlByUserId(r.Context(), userId)
	//fmt.Println(url)
	//fmt.Println(perms)
	if _, ok := perms[url]; !ok {
		return false
	}
	return true
}

func (s *sSysMenu) GetPermsUrlByUserId(ctx context.Context, userId int64) (perms map[string]bool) {
	cache := service.Cache().Driver()
	//cache.Del(ctx, consts.CacheKeyPermsUrl)
	has, err := cache.Exists(ctx, consts.CacheKeyPermsUrl)
	if err != nil {
		return
	}
	if has != 0 {
		iPerms, err := cache.Get(ctx, consts.CacheKeyPermsUrl)
		if err != nil {
			utility.WriteErrLog(ctx, err, "获取用户权限数据缓存失败")
		}
		err = iPerms.Struct(&perms)
		return
	}
	perms = make(map[string]bool, 0)
	roleIds := []int64{}
	rolesList, err := service.SysRole().GetRolesByUid(ctx, userId)
	if err != nil {
		return
	}
	for _, v := range rolesList.SysRole {
		roleIds = append(roleIds, v.RoleId)
	}
	//获取菜单
	menuList, err := s.GetMenuTreeByRoleids(ctx, roleIds)
	for _, menu := range menuList {
		perms[menu.ApiPath] = true
	}
	maps.Copy(perms, SpecialApiPath)
	//放入缓存
	cache.Set(ctx, consts.CacheKeyPermsUrl, perms)
	return
}

func (s *sSysMenu) GetMenuTreeSelect(ctx context.Context) (menuTree []*model.SysMenuTreeRes, err error) {
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
	menuList, err := s.GetMenuTreeByRoleids(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	menuTree = s.MenuTreeSelect(menuList, 0)
	return
}

func (s *sSysMenu) MenuTreeSelect(menuList []*entity.SysMenu, pid int64) (tree []*model.SysMenuTreeRes) {
	tree = make([]*model.SysMenuTreeRes, 0)
	for _, v := range menuList {
		if v.ParentId == pid {
			//fmt.Printf("%+v\n", v)
			dd := &model.SysMenuTreeRes{Id: v.MenuId, Label: v.MenuName}
			child := s.MenuTreeSelect(menuList, v.MenuId)
			if child != nil {
				dd.Children = child
			}
			tree = append(tree, dd)
		}
	}
	return
}

// 菜单列表
func (s *sSysMenu) GetMenuList(ctx context.Context, req *v1.GetMenuListReq) (menuList *v1.GetMenuListRes, err error) {
	var list []*entity.SysMenu
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysMenu.Ctx(ctx)
		if req.Status != "" {
			m = m.Where(dao.SysMenu.Columns().Status, req.Status)
		}
		if req.MenuName != "" {
			m = m.WhereLike(dao.SysMenu.Columns().MenuName, "%"+req.MenuName+"%")
		}
		err = m.Scan(&list)
		utility.WriteErrLog(ctx, err, "获取菜单失败")
	})
	menuArr := make([]*model.SysMenuList, len(list))
	for k, menu := range list {
		m := &model.SysMenuList{}
		m.SysMenu = menu
		menuArr[k] = m
	}
	menuList = &v1.GetMenuListRes{}
	menuList.List = menuArr
	return
}

// 获取单条数据
func (s sSysMenu) GetOneMenuById(ctx context.Context, menuId int64) (menu *entity.SysMenu, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//菜单信息
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, menuId).Scan(&menu)
		utility.WriteErrLog(ctx, err, "获取菜单数据失败")
	})
	return
}

// 修改数据
func (s sSysMenu) UpdateMenu(ctx context.Context, req *v1.PutMenuUpdateReq) (res *v1.PutMenuUpdateRes, err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).WherePri(req.MenuId).Update(do.SysMenu{
			ParentId:   req.ParentId,
			MenuType:   req.MenuType,
			Icon:       req.Icon,
			MenuName:   req.MenuName,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			IsFrame:    req.IsFrame,
			Path:       req.Path,
			ApiPath:    req.ApiPath,
			Component:  req.Component,
			Query:      req.Query,
			IsCache:    req.IsCache,
			Perms:      req.Perms,
			Remark:     req.Remark,
			Visible:    req.Visible,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "修改菜单数据失败")
	})
	return
}

// 初始化apipath
func (s *sSysMenu) InitApiPath(ctx context.Context) (err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		var menu []*entity.SysMenu
		err = dao.SysMenu.Ctx(ctx).Where("api_path=''").Scan(&menu)
		for _, m := range menu {
			if m.Perms == "" {
				continue
			}
			perms := strings.Split(m.Perms, ":")
			apiPath := "/" + perms[0] + "/" + perms[1]
			switch perms[2] {
			case "add":
				apiPath = "post" + apiPath
			case "list":
				apiPath = "get" + apiPath + "/list"
			case "query":
				apiPath = "get" + apiPath + "/{" + perms[1] + "Id}"
			case "edit":
				apiPath = "put" + apiPath
			case "remove":
				apiPath = "delete" + apiPath + "/{" + perms[1] + "Id}"
			default:
				continue
			}
			fmt.Println(apiPath)
			_, err = dao.SysMenu.Ctx(ctx).WherePri(&m.MenuId).Update(do.SysMenu{
				ApiPath: apiPath,
			})
		}
	})
	return
}

// 添加数据
func (s sSysMenu) Add(ctx context.Context, req *v1.PostMenuAddReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).Data(do.SysMenu{
			ParentId:   req.ParentId,
			MenuType:   req.MenuType,
			Icon:       req.Icon,
			MenuName:   req.MenuName,
			OrderNum:   req.OrderNum,
			Status:     req.Status,
			IsFrame:    req.IsFrame,
			Path:       req.Path,
			ApiPath:    req.ApiPath,
			Visible:    req.Visible,
			Component:  req.Component,
			Query:      req.Query,
			IsCache:    req.IsCache,
			Perms:      req.Perms,
			Remark:     req.Remark,
			UpdateTime: gtime.Now(),
			CreateTime: gtime.Now(),
			CreateBy:   adminName,
			UpdateBy:   adminName,
		}).Insert()
		utility.WriteErrLog(ctx, e, "添加菜单数据失败")
	})
	return
}

// 删除数据
func (s sSysMenu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (err error) {
	adminName := gconv.String(ctx.Value(consts.CtxAdminName))
	err = g.Try(ctx, func(ctx context.Context) {
		_, e := dao.SysMenu.Ctx(ctx).WherePri(req.MenuId).Update(do.SysMenu{
			Status:     consts.SysMenuStatusNo,
			UpdateTime: gtime.Now(),
			UpdateBy:   adminName,
		})
		utility.WriteErrLog(ctx, e, "删除菜单数据失败")
	})
	return
}
