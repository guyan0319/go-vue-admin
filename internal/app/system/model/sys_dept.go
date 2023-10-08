package model

import (
	"go-vue-admin/internal/app/system/model/entity"
)

type SysDeptSRes struct {
	*entity.SysDept
	Children []*SysDeptSRes `json:"children"`
}

type SysDeptTreeRes struct {
	Id       int64             `json:"id"     description:"部门id"`
	Label    string            `json:"label"   description:"部门名称"`
	Children []*SysDeptTreeRes `json:"children"   description:"子部门"`
}
type DeptIds struct {
	Ids []int64 `json:"ids"     description:"部门id"`
}
