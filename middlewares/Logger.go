package middlewares

import (
	"github.com/kataras/iris/context"
	"time"
)

func Logger(ctx context.Context) {
	start := time.Now()

	ctx.Next()

	ctx.Application().Logger().Infof("%s\t%s\t%s",
		ctx.Request().Method,
		ctx.Request().URL.String(),
		time.Since(start),
	)
}
