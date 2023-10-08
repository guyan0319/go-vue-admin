package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetSysDictTypeListReq struct {
	g.Meta     `path:"/system/dict/type/list" method:"Get" tags:"Get Get sys_dict_type" summary:"Get Get sys_dict_type"`
	CreateTime *gtime.Time `p:"createTime"`
	DictName   string      `p:"dictName"`
	DictType   string      `p:"dictType"`
	Status     string      `p:"status"`
	common.PageReq
}

type GetSysDictTypeListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.SysDictType `json:"list"`
	Total  int                   `json:"total"`
}

type PostSysDictTypeReq struct {
	g.Meta   `path:"/system/dict/type" method:"Post" tags:"Post sys_dict_type" summary:"Post sys_dict_type"`
	Remark   string `p:"remark" `
	DictName string `p:"dictName" `
	DictType string `p:"dictType" `
	Status   string `p:"status" `
}

type PostSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type PutSysDictTypeReq struct {
	g.Meta   `path:"/system/dict/type" method:"Put" tags:"Put sys_dict_type" summary:"Put sys_dict_type"`
	Remark   string `p:"remark" `
	DictName string `p:"dictName" `
	DictType string `p:"dictType" `
	DictId   int64  `p:"dictId"  v:"required" `
	Status   string `p:"status" `
}

type PutSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSysDictTypeReq struct {
	g.Meta `path:"/system/dict/type/{dictId}" method:"Delete" tags:"Delete sys_dict_type" summary:"Delete sys_dict_type"`
	DictId string `p:"dictId"  v:"required" `
}

type DeleteSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type GetSysDictTypeReq struct {
	g.Meta `path:"/system/dict/type/{dictId}" method:"Get" tags:"Get sys_dict_type" summary:"Get sys_dict_type"`
	DictId int64 `p:"dictId"  v:"required" `
}

type GetSysDictTypeRes struct {
	g.Meta `mime:"application/json"`
	*entity.SysDictType
}
