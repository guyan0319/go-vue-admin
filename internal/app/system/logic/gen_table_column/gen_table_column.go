package gen_table_column

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
)

func init() {
	service.RegisterGenTableColumn(New())
}

func New() *sGenTableColumn {
	return &sGenTableColumn{}
}

type sGenTableColumn struct {
}

func (s *sGenTableColumn) GetTableColumnByIds(ctx context.Context, ids []int64) (columnMap map[int64]*entity.GenTableColumn, err error) {
	return
}
func (s *sGenTableColumn) GetTableColumnByTableId(ctx context.Context, tableId int64) (tableColumn []*entity.GenTableColumn, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err := dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, tableId).OrderDesc(dao.GenTableColumn.Columns().Sort).Scan(&tableColumn)
		utility.WriteErrLog(ctx, err, "获取生成表字段数据失败")
	})
	return
}
func (s *sGenTableColumn) GetTableColumnByTableIds(ctx context.Context, tableIds []int64) (tableMap map[string][]*entity.GenTableColumn, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var tableColumn []*entity.GenTableColumn
		err := dao.GenTableColumn.Ctx(ctx).WhereIn(dao.GenTableColumn.Columns().TableId, tableIds).OrderDesc(dao.GenTableColumn.Columns().Sort).Scan(&tableColumn)
		utility.WriteErrLog(ctx, err, "获取生成表字段数据失败")
		tableMap = make(map[string][]*entity.GenTableColumn, 0)
		for _, column := range tableColumn {
			if column.ColumnName == "create_by" || column.ColumnName == "update_by" || column.ColumnName == "update_time" {
				continue
			}
			tableMap[column.TableId] = append(tableMap[column.TableId], column)
		}
	})
	return
}
