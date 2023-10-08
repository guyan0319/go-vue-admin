package dict

import (
	"context"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

type dictController struct {
}

var Dict = dictController{}

func (d *dictController) DictData(ctx context.Context, req *v1.GetDictDataReq) (res *v1.GetDictDataRes, err error) {
	dictData, err := service.SysDictData().GetDictDataByType(ctx, req.DictType)
	res = &v1.GetDictDataRes{
		DictData: dictData,
	}
	return
}
