package middlewares

import "github.com/kataras/iris/context"

func ErrorHandler(ctx context.Context) {
	// .Values are used to communicate between handlers, middleware.
	errMessage := ctx.Values().GetString("error")
	if errMessage != "" {
		ctx.Writef("Internal server error: %s", errMessage)
		return
	}

	ctx.Writef("(Unexpected) internal server error")
}
