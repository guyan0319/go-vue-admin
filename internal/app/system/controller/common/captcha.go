package common

import (
"context"
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
	res = &v1.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}



