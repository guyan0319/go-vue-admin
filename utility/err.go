package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func WriteErrLog(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		g.Log().Error(ctx, err.Error())
		if len(msg) > 0 {
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}
