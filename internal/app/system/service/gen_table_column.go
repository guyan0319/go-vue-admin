// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-vue-admin/internal/app/system/model/entity"
)

type (
	IGenTableColumn interface {
		GetTableColumnByIds(ctx context.Context, ids []int64) (columnMap map[int64]*entity.GenTableColumn, err error)
		GetTableColumnByTableId(ctx context.Context, tableId int64) (tableColumn []*entity.GenTableColumn, err error)
		GetTableColumnByTableIds(ctx context.Context, tableIds []int64) (tableMap map[string][]*entity.GenTableColumn, err error)
	}
)

var (
	localGenTableColumn IGenTableColumn
)

func GenTableColumn() IGenTableColumn {
	if localGenTableColumn == nil {
		panic("implement not found for interface IGenTableColumn, forgot register?")
	}
	return localGenTableColumn
}

func RegisterGenTableColumn(i IGenTableColumn) {
	localGenTableColumn = i
}
