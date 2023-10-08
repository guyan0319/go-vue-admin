package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetDeptTreeReq struct {
	g.Meta `path:"/system/user/deptTree" method:"get" tags:"UserService" summary:"current Data"`
}
type GetDeptTreeRes struct {
	g.Meta   `mime:"application/json"`
	DeptTree []*model.SysDeptTreeRes `json:"deptTree"`
}
type GetDeptListReq struct {
	g.Meta   `path:"/system/dept/list" method:"get" tags:"DeptService" summary:"current Data"`
	Status   string `p:"status"`
	DeptName string `p:"deptName"`
}
type GetDeptListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.SysDept `json:"list"`
}
type GetDeptListUpdateReq struct {
	g.Meta `path:"/system/dept/list/exclude/{deptId}" method:"get" tags:"DeptService" summary:"current Data"`
	DeptId int64 `p:"deptId"`
}
type GetDeptListUpdateRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.SysDept `json:"list"`
}
type GetDeptUpdateReq struct {
	g.Meta `path:"/system/dept/{deptId}" method:"get" tags:"DeptService" summary:"current Data"`
	DeptId int64 `p:"deptId" v:"required"`
}
type GetDeptUpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDept
}
type DeleteDeptReq struct {
	g.Meta `path:"/system/dept/{deptId}" method:"delete" tags:"DeptService" summary:"current Data"`
	DeptId int64 `p:"deptId" v:"required"`
}
type DeleteDeptRes struct {
	g.Meta `mime:"application/json"`
}
type PutDeptUpdateReq struct {
	g.Meta   `path:"/system/dept" method:"put" tags:"DeptService" summary:"current Data"`
	DeptId   int64  `p:"deptId" v:"required"`
	ParentId int64  `p:"parentId"  v:"required" description:"父部门id"`
	DeptName string `p:"deptName"  v:"required"  description:"部门名称"`
	OrderNum int    `p:"orderNum"  v:"required" description:"显示顺序"`
	Leader   string `p:"leader"     description:"负责人"`
	Phone    string `p:"phone"      description:"联系电话"`
	Email    string `p:"email"      description:"邮箱"`
	Status   string `p:"status"     description:"部门状态（0正常 1停用）"`
}
type PutDeptUpdateRes struct {
	g.Meta `mime:"application/json"`
}
type PostDeptAddReq struct {
	g.Meta   `path:"/system/dept" method:"post" tags:"DeptService" summary:"current Data"`
	ParentId int64  `p:"parentId"  v:"required" description:"父部门id"`
	DeptName string `p:"deptName"  v:"required"  description:"部门名称"`
	OrderNum int    `p:"orderNum"  v:"required" description:"显示顺序"`
	Leader   string `p:"leader"     description:"负责人"`
	Phone    string `p:"phone"      description:"联系电话"`
	Email    string `p:"email"      description:"邮箱"`
	Status   string `p:"status"     description:"部门状态（0正常 1停用）"`
}
type PostDeptAddRes struct {
	g.Meta `mime:"application/json"`
}
