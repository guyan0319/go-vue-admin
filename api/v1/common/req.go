package common

type PageReq struct {
	PageNum  int               `p:"pageNum"`
	PageSize int               `p:"pageSize"`
	OrderBy  string            `p:"orderBy"` //排序方式
	Params   map[string]string `p:"params"`  //时间范围
}
