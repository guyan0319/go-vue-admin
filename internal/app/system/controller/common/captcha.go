package common

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/google/uuid"
	v1 "go-vue-admin/api/v1"
	"go-vue-admin/internal/app/system/service"
)

var Captcha = captchaController{}

type captchaController struct {
}

// CaptchaImage 获取验证码
func (c *captchaController) CaptchaImage(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	guid := uuid.New()
	res = &v1.CaptchaRes{
		Key:            idKeyC,
		Img:            base64stringC,
		Uuid:           guid.String(),
		CaptchaEnabled: true,
	}
	return
}

// CaptchaImage 获取验证码
func (c *captchaController) Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	err = service.SysMenu().InitApiPath(ctx)
	//ghttp.RequestFromCtx(ctx).Response.ServeFileDownload("temp/down/1/goadmin.zip")
	//ghttp.RequestFromCtx(ctx).Response.Write("nihao")
	//return
	//db := g.DB()
	//db.GetSchema()
	////fmt.Println(db.GetSchema(), "aaaaaa")
	////fmt.Println(db.GetGroup())
	////fmt.Println(db.GetConfig())
	//
	//list, _ := db.GetAll(ctx, "SELECT * FROM information_schema.Tables WHERE TABLE_SCHEMA='"+db.GetSchema()+"'")
	//////list, err := db.GetAll(ctx, "show tables")
	////
	//for _, tableName := range list {
	//	fmt.Println(tableName["TABLE_COMMENT"])
	//	fmt.Println(tableName["TABLE_NAME"])
	//	//fmt.Println(tableName["Tables_in_gvadmindb"])
	//}
	////fmt.Println(list)
	//////获取表字段
	//tableList, err := db.Ctx(ctx).TableFields(ctx, "sys_user")
	//fmt.Println(tableList)
	//for _, field := range tableList {
	//	fmt.Println(field.Key)
	//	fmt.Println(field.Index)
	//	fmt.Println(field.Type)
	//	fmt.Println(field.Name)
	//	fmt.Println(field.Default)
	//	fmt.Println(field.Null)
	//	fmt.Println(field.Extra)
	//	fmt.Println("aaaaa")
	//	//fmt.Println(s, field.Type)
	//	//typeName, _ := service.GenTable().GetGolangTypeBySqlType(ctx, field)
	//	//fmt.Println(field.Type, typeName)
	//}
	return

}
