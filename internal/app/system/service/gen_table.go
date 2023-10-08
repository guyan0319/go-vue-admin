// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/model"
	"go-vue-admin/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IGenTable interface {
		GetTableList(ctx context.Context, req v1.ToolListCommon) (tableList v1.ToolListCommonRes, err error)
		GetGenTablesUpdate(ctx context.Context, req *v1.GetGenTablesUpdateReq) (tableList *v1.GetGenTablesUpdateRes, err error)
		Delete(ctx context.Context, req *v1.DeleteGenTableReq) (tableList *v1.DeleteGenTableRes, err error)
		Add(ctx context.Context, req *v1.PostGenTableAddReq) (tableList *v1.PostGenTableAddRes, err error)
		Update(ctx context.Context, req *v1.PutGenTableUpdateReq) (tableList *v1.PutGenTableUpdateRes, err error)
		// 删除指定表名数据
		DeleteTable(ctx context.Context, tableName string) (err error)
		// 生成表数据
		InitTables(ctx context.Context, tableName string) (err error)
		GetGolangTypeBySqlType(ctx context.Context, field *gdb.TableField) (typeName string, err error)
		CheckTable(ctx context.Context, tableName string) (bool, error)
		CheckTableColumn(ctx context.Context, columnName string, tableId int64) (bool, error)
		GenPreview(ctx context.Context, req *v1.GetGenPreviewReq) (res *v1.GetGenPreviewRes, err error)
		BatchGenCode(ctx context.Context, req *v1.GetBatchGenCodeReq) (tableList *v1.GetBatchGenCodeRes, err error)
		// 生成controller文件
		GenControllerFile(codeDir string, table *entity.GenTable, tplCtrlContent []model.TplCtrlContent) (err error)
		LogicTplContent(funcType string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (tlc model.TplLogicContent)
		ApiJsTplContent(funcType string, table *entity.GenTable) (tlc model.TplApiJsContent)
		// 生成logic文件
		GenLogicFile(codeDir string, table *entity.GenTable, tplLogicContent []model.TplLogicContent) (err error)
		GetApiPath(tableName string) string
		// 生成api文件
		GenApiFile(codeDir string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (err error)
		GetStatusDictType(tableName, columnName string) (res string, ok bool)
		VueTplContent(funcSilce []string, table *entity.GenTable, tableColumnMap []*entity.GenTableColumn) (tvc model.TplVueContent)
		GetDictTypeVue(dictType []string) string
		// 生成vue文件
		GenVueFile(codeDir string, table *entity.GenTable, tvc model.TplVueContent) (err error)
		// 生成consts文件
		GenConstsFile(codeDir string, table *entity.GenTable) (err error)
		// 生成apijs文件
		GenApiJsFile(codeDir string, table *entity.GenTable, tplApiJsContent []model.TplApiJsContent) (err error)
		// 生成文件
		WriteDownFile(src, dst string, params g.Map, fileName string) (err error)
		GetReqMeta(url, method, tableName string) (gMeta string)
		CreateDir(userId string) (dir string)
		GetGenTableSilceByName(ctx context.Context, tableName []string) (tableSilce []*entity.GenTable, err error)
		GetGenTableSilceByTableId(ctx context.Context, tableId int64) (tableSilce *entity.GenTable, err error)
	}
)

var (
	localGenTable IGenTable
)

func GenTable() IGenTable {
	if localGenTable == nil {
		panic("implement not found for interface IGenTable, forgot register?")
	}
	return localGenTable
}

func RegisterGenTable(i IGenTable) {
	localGenTable = i
}
