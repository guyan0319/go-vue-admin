package {{.packageName}}

import (
	"github.com/gogf/gf/v2/frame/g"
)
{{range  $value := .apiSilce}}
type  {{$value.TypeName}} struct {
   {{$value.TypeContent}}
}
{{end}}
