package sys_dict_data

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-vue-admin/internal/app/system/dao"
	"go-vue-admin/internal/app/system/model/entity"
	"go-vue-admin/internal/app/system/service"
	"go-vue-admin/utility"
)

func init() {
	service.RegisterSysDictData(New())
}

func New() *sSysDictData {
	return &sSysDictData{}
}

type sSysDictData struct {
}

func (s *sSysDictData) GetDictDataByType(ctx context.Context, dictType string) (dictdata []*entity.SysDictData, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//字典数据表
		err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictType, dictType).Scan(&dictdata)
		utility.WriteErrLog(ctx, err, "获取字典数据表失败")
	})
	return
}
