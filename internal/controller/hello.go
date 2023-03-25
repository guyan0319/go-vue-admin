package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/template-single/apiv1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
