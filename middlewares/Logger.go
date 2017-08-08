package middlewares

import (
	"github.com/kataras/iris/context"
)

func Logger(ctx context.Context) {
	ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
	ctx.Next()
}
