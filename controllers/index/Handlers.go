package todo

import (
	"github.com/kataras/iris/context"
	"github.com/paduvi/BasicIrisExample/httputils"
	. "github.com/paduvi/BasicIrisExample/models"
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/actions"
)

func Index(ctx context.Context) {
	ctx.Text("Welcome!")
}

func SubIndex(ctx context.Context) {
	name := ctx.Params().Get("name")
	ctx.Text("Welcome, " + name + "!")
}

func PingRemote(ctx context.Context) {
	done := make(chan Result)
	work := httputils.Job{Result: done, Handle: actions.PingRemote}

	// Push the work onto the queue.
	httputils.JobQueue <- work

	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", "ping failed. "+result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text(result.Data.(string))
}
