import request from '@/utils/request'

{{range  $value := .jsFuncSilce}}
// {{$value.Notes}}
export function {{$value.FuncName}}(query) {
  return request({
    {{$value.Content}}
  })
}
{{end}}
