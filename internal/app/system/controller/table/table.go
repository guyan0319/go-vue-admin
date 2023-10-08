package table

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type tableController struct {
}

var Table = tableController{}

func (t *tableController) GetToolList(ctx context.Context, req *v1.GetToolListReq) (res *v1.GetToolListRes, err error) {
	var toolListCommon v1.ToolListCommon
	res = &v1.GetToolListRes{}
	toolListCommon = req.ToolListCommon
	toolListCommon.Status = "0"
	res.ToolListCommonRes, err = service.GenTable().GetTableList(ctx, toolListCommon)
	return
}
func (t *tableController) GetToolDbList(ctx context.Context, req *v1.GetToolDbListReq) (res *v1.GetToolDbListRes, err error) {

	var toolListCommon v1.ToolListCommon
	res = &v1.GetToolDbListRes{}
	toolListCommon = req.ToolListCommon
	toolListCommon.Status = "1"
	res.ToolListCommonRes, err = service.GenTable().GetTableList(ctx, toolListCommon)
	return
}

func (t *tableController) GetGenTableUpdate(ctx context.Context, req *v1.GetGenTablesUpdateReq) (res *v1.GetGenTablesUpdateRes, err error) {
	res, err = service.GenTable().GetGenTablesUpdate(ctx, req)
	return
}

func (t *tableController) InitTables(ctx context.Context, req *v1.GetInitTablesReq) (res *v1.GetInitTablesRes, err error) {
	err = service.GenTable().InitTables(ctx, "")
	return
}
func (t *tableController) Delete(ctx context.Context, req *v1.DeleteGenTableReq) (res *v1.DeleteGenTableRes, err error) {
	res, err = service.GenTable().Delete(ctx, req)
	return
}
func (t *tableController) Add(ctx context.Context, req *v1.PostGenTableAddReq) (res *v1.PostGenTableAddRes, err error) {
	res, err = service.GenTable().Add(ctx, req)
	return
}
func (t *tableController) Update(ctx context.Context, req *v1.PutGenTableUpdateReq) (res *v1.PutGenTableUpdateRes, err error) {
	res, err = service.GenTable().Update(ctx, req)
	return
}
func (t *tableController) BatchGenCode(ctx context.Context, req *v1.GetBatchGenCodeReq) (res *v1.GetBatchGenCodeRes, err error) {
	res, err = service.GenTable().BatchGenCode(ctx, req)
	return
}

// 同步
func (t *tableController) GenSynchDb(ctx context.Context, req *v1.GetGenSynchDbReq) (res *v1.GetGenSynchDbRes, err error) {
	err = service.GenTable().InitTables(ctx, req.Table)
	return
}

// 预览
func (t *tableController) GenPreview(ctx context.Context, req *v1.GetGenPreviewReq) (res *v1.GetGenPreviewRes, err error) {
	res, err = service.GenTable().GenPreview(ctx, req)
	return
}
