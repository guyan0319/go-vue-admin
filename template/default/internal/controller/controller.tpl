package {{.packageName}}

import (
    "context"
)

type {{.nameController}} struct {
}

var {{.controllerName}} = {{.nameController}} {}
{{range  $value := .funcSilce}}
func (s *{{$value.NameController}} ) {{$value.FuncName}}(ctx context.Context, req *{{$value.ApiPackageName}}.{{$value.ServiceReq}}) (res *{{$value.ApiPackageName}}.{{$value.ServiceRes}}, err error) {
	res, err = service.{{$value.ServiceName}}().{{$value.FuncName}}(ctx, req)
	return
}
{{end}}

