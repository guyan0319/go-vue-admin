package common

import (
"context"
)

var Captcha = captchaController{}

type captchaController struct {
}

// CaptchaImage 获取验证码
func (c *captchaController) CaptchaImage(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}



