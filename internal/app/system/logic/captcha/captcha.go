package captcha

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
	"go-vue-admin/internal/app/system/service"
)

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	return &sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
}

type sCaptcha struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

var (
	captcha = sCaptcha{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjkmnpqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
)

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	driver := s.driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, s.store)
	idKeyC, base64stringC, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(s.driver, s.store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
