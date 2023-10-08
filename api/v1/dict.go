package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/model/entity"
)

type GetDictDataReq struct {
	g.Meta   `path:"/system/dict/type" method:"get" tags:"UserService" summary:"current Data"`
	DictType string `p:"dictType" v:"required"`
}
type GetDictDataRes struct {
	g.Meta   `mime:"application/json"`
	DictData []*entity.SysDictData `json:"dictData"`
}

type GetDictTypeOptionSelectReq struct {
	g.Meta `path:"/system/dict/type/optionselect" method:"get" tags:"UserService" summary:"current Data"`
}
type GetDictTypeOptionSelectRes struct {
	g.Meta   `mime:"application/json"`
	DictType []*entity.SysDictType `json:"dictType"`
}
