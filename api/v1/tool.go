package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/api/v1/common"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"
)

type DeleteGenTableReq struct {
	g.Meta  `path:"/tool/gen/{tableId}" method:"delete" tags:"Table Service" summary:"delete gen_table"`
	TableId string `json:"tableId"`
}
type PostGenTableAddRes struct {
	g.Meta `mime:"application/json"`
}
type PostGenTableAddReq struct {
	g.Meta `path:"/tool/gen/importTable" method:"post" tags:"Table Service" summary:"add gen_table"`
	Tables string `json:"tables"`
}

type DeleteGenTableRes struct {
	g.Meta `mime:"application/json"`
}

type GetToolListReq struct {
	g.Meta `path:"/tool/gen/list" method:"get" tags:"Table Service" summary:"table list"`
	ToolListCommon
}
type GetToolListRes struct {
	ToolListCommonRes
}

type GetGenTablesUpdateReq struct {
	g.Meta  `path:"/tool/gen/{tableId}" method:"get" tags:"Table Service" summary:"table list"`
	TableId int64 `json:"tableId"`
}
type GetGenTablesUpdateRes struct {
	g.Meta `mime:"application/json"`
	Info   *model.GenTableInfo      `json:"info"`
	Rows   []*entity.GenTableColumn `json:"rows"`
	Tables []*model.GenTableInfo    `json:"tables"`
}
type GetToolDbListReq struct {
	g.Meta `path:"/tool/gen/db/list" method:"get" tags:"Table Service" summary:"table list"`
	ToolListCommon
}
type ToolListCommon struct {
	TableName    string `p:"tableName" `
	TableComment string `p:"tableComment" `
	Status       string `p:"status" `
	common.PageReq
}
type GetToolDbListRes struct {
	ToolListCommonRes
}
type ToolListCommonRes struct {
	g.Meta `mime:"application/json"`
	Rows   []*model.GenTableList `json:"rows"`
	Total  int                   `json:"total"`
}
type GetInitTablesReq struct {
	g.Meta `path:"/tool/gen/initTable" method:"get" tags:"Table Service" summary:"gen table"`
}
type GetInitTablesRes struct {
	g.Meta `mime:"application/json"`
}

type PutGenTableUpdateReq struct {
	g.Meta `path:"/tool/gen" method:"put" tags:"Table Service" summary:"update gen_table"`
	*model.GenTableUpdate
}
type PutGenTableUpdateRes struct {
	g.Meta `mime:"application/json"`
}
type GetBatchGenCodeReq struct {
	g.Meta `path:"/tool/gen/batchGenCode" method:"get" tags:"Table Service" summary:"download gen_table"`
	Tables string `p:"tables" v:"required"`
}
type GetBatchGenCodeRes struct {
	//g.Meta `mime:"application/json"`
}

// 同步
type GetGenSynchDbReq struct {
	g.Meta `path:"/tool/gen/synchDb/{Tables}" method:"get" tags:"Table Service" summary:" synch gen_table"`
	Table  string `p:"tables" v:"required"`
}
type GetGenSynchDbRes struct {
	g.Meta `mime:"application/json"`
}
type GetGenPreviewReq struct {
	g.Meta  `path:"/tool/gen/preview/{tableId}" method:"get" tags:"Table Service" summary:" synch gen_table"`
	TableId int64 `p:"tables" v:"required"`
}
type GetGenPreviewRes struct {
	g.Meta `mime:"application/json"`
	List   map[string]string `json:"list"`
}
