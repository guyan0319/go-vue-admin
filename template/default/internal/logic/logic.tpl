package {{.packageName}}

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.Register{{.className}}(New())
}

func New() *s{{.className}} {
	return &s{{.className}}{}
}

type s{{.className}} struct {
}

{{range  $value := .logicSilce}}
// {{$value.Notes}}
func (s *s{{$value.ClassName}}) {{$value.FuncName}}(ctx context.Context, req *{{$value.ApiPackageName}}.{{$value.ParamReq}}) (res *{{$value.ApiPackageName}}.{{$value.ParamRes}}, err error) {
	{{$value.Content}}
	return
}
{{end}}