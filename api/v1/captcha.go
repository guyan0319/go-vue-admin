package v1

import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/captchaImage" tags:"验证码" method:"get" summary:"获取验证码"`
}
type CaptchaRes struct {
	g.Meta         `mime:"application/json"`
	Key            string `json:"key"`
	Img            string `json:"img"`
	Uuid           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled"`
}
type TestReq struct {
	g.Meta `path:"/test" tags:"验证码" method:"get" summary:"获取验证码"`
}
type TestRes struct {
	g.Meta `mime:"application/json"`
}
