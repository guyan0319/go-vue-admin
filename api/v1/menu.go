package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetTreeSelectReq struct {
	g.Meta `path:"/system/menu/treeselect" method:"get" tags:"MenuService" summary:"menu treeselect"`
}
type GetTreeSelectRes struct {
	g.Meta `mime:"application/json"`
	Tree   []*model.SysMenuTreeRes `json:"tree"`
}

type GetMenuListReq struct {
	g.Meta   `path:"/system/menu/list" method:"get" tags:"MenuService" summary:"menu list"`
	MenuName string `p:"menuName"`
	Status   string `p:"status"`
}
type GetMenuListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SysMenuList `json:"list"`
}
type GetMenuUpdateReq struct {
	g.Meta `path:"/system/menu/{menuId}" method:"get" tags:"MenuService" summary:"menu list"`
	MenuId int64 `p:"menuId" v:"required"`
}
type GetMenuUpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysMenu
}
type MenuDeleteReq struct {
	g.Meta `path:"/system/menu/{menuId}" method:"DELETE" tags:"MenuService" summary:"delete menu"`
	MenuId int64 `p:"menuId" v:"required"`
}
type MenuDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type PutMenuUpdateReq struct {
	g.Meta    `path:"/system/menu" method:"put" tags:"MenuService" summary:"menu list"`
	MenuId    int64  `p:"menuId" v:"required"`
	MenuName  string `p:"menuName"  v:"required"  description:"菜单名称"`
	ParentId  int64  `p:"parentId"   description:"父菜单ID"`
	OrderNum  int    `p:"orderNum"   v:"required" description:"显示顺序"`
	Path      string `p:"path"      description:"路由地址"`
	ApiPath   string `p:"apiPath"    description:"后台api地址"`
	IsFrame   int    `p:"isFrame"    description:"是否为外链（0是 1否）"`
	MenuType  string `p:"menuType"   description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible   string `p:"visible"    description:"菜单状态（0显示 1隐藏）"`
	Status    string `p:"status"     description:"菜单状态（0正常 1停用）"`
	Icon      string `p:"icon"       description:"菜单图标"`
	Component string `p:"component"  description:"组件路径"`
	Query     string `p:"query"      description:"路由参数"`
	IsCache   int    `p:"isCache"    description:"是否缓存（0缓存 1不缓存）"`
	Perms     string `p:"perms"      description:"权限标识"`
	Remark    string `p:"remark"     description:"备注"`
}
type PutMenuUpdateRes struct {
	g.Meta `mime:"application/json"`
}
type PostMenuAddReq struct {
	g.Meta    `path:"/system/menu" method:"post" tags:"MenuService" summary:"menu list"`
	MenuName  string `p:"menuName"  v:"required"  description:"菜单名称"`
	ParentId  int64  `p:"parentId"   description:"父菜单ID"`
	OrderNum  int    `p:"orderNum"   v:"required" description:"显示顺序"`
	Path      string `p:"path"     description:"路由地址"`
	ApiPath   string `p:"apiPath"      description:"后台api地址"`
	IsFrame   int    `p:"isFrame"    description:"是否为外链（0是 1否）"`
	MenuType  string `p:"menuType"   description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible   string `p:"visible"    description:"菜单状态（0显示 1隐藏）"`
	Status    string `p:"status"     description:"菜单状态（0正常 1停用）"`
	Icon      string `p:"icon"       description:"菜单图标"`
	Component string `p:"component"  description:"组件路径"`
	Query     string `p:"query"      description:"路由参数"`
	IsCache   int    `p:"isCache"    description:"是否缓存（0缓存 1不缓存）"`
	Perms     string `p:"perms"      description:"权限标识"`
	Remark    string `p:"remark"     description:"备注"`
}
type PostMenuAddRes struct {
	g.Meta `mime:"application/json"`
}
