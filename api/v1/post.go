package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetPostListReq struct {
	g.Meta   `path:"/system/post/list" method:"get" tags:"Post Service" summary:"post Data"`
	PostName string `p:"postName" `
	PostCode string `p:"postCode" `
	Status   string `p:"status" `
	common.PageReq
}
type GetPostListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.SysPost `json:"list"`
	Total  int               `json:"total"`
}
type PostPostAddReq struct {
	g.Meta   `path:"/system/post" method:"post" tags:"Post Service" summary:"post Data"`
	PostName string `p:"postName"  v:"required"`
	PostCode string `p:"postCode"  v:"required"`
	PostSort string `p:"postSort"  v:"required"`
	Status   string `p:"status" v:"required"`
	Remark   string `p:"remark" `
}
type PostPostAddRes struct {
	g.Meta `mime:"application/json"`
}
type PutPostUpdateReq struct {
	g.Meta   `path:"/system/post" method:"put" tags:"Post Service" summary:"post Data"`
	PostId   string `p:"postId"  v:"required"`
	PostName string `p:"postName"  v:"required"`
	PostCode string `p:"postCode"  v:"required"`
	PostSort string `p:"postSort"  v:"required"`
	Status   string `p:"status" v:"required"`
	Remark   string `p:"remark" `
}
type PutPostUpdateRes struct {
	g.Meta `mime:"application/json"`
}

type GetPostUpdateReq struct {
	g.Meta `path:"/system/post/{postId}" method:"get" tags:"Post Service" summary:"post Data"`
	PostId int64 `p:"postId"  v:"required"`
}
type GetPostUpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysPost
}

type DeletePostReq struct {
	g.Meta `path:"/system/post/{postId}" method:"delete" tags:"Post Service" summary:"current Data"`
	PostId string `p:"postId" v:"required"`
}
type DeletePostRes struct {
	g.Meta `mime:"application/json"`
}
