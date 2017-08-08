package todo

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	. "github.com/paduvi/BasicIrisExample/models"
	"github.com/paduvi/BasicIrisExample/actions"
	"io/ioutil"
	"io"
	"github.com/kataras/iris"
	"github.com/paduvi/BasicIrisExample/httputils"
)

func MessageIndex(ctx context.Context) {
	done := make(chan Result)
	work := httputils.Job{Result: done, Handle: actions.ListMessage}

	// Push the work onto the queue.
	httputils.JobQueue <- work

	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	var messages Messages
	json.Unmarshal(result.Data.([]byte), &messages)
	ctx.JSON(messages)
}

func MessageShow(ctx context.Context) {
	messageId, err := ctx.Params().GetInt("messageId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan Result)
	work := httputils.Job{Payload: Message{Id: messageId}, Result: done, Handle: actions.FindMessage}

	// Push the work onto the queue.
	httputils.JobQueue <- work

	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	var message Message
	json.Unmarshal(result.Data.([]byte), &message)
	ctx.JSON(message)
}

func MessageCreate(ctx context.Context) {
	var message Message
	body, err := ioutil.ReadAll(io.LimitReader(ctx.Request().Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := ctx.Request().Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &message); err != nil {
		ctx.ContentType("application/json")
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		if _, err := ctx.JSON(err); err != nil {
			panic(err)
		}
	}

	done := make(chan Result)
	work := httputils.Job{Payload: message, Result: done, Handle: actions.CreateMessage}

	// Push the work onto the queue.
	httputils.JobQueue <- work

	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	json.Unmarshal(result.Data.([]byte), &message)
	ctx.JSON(message)
}

func MessageDelete(ctx context.Context) {
	messageId, err := ctx.Params().GetInt("messageId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan Result)
	work := httputils.Job{Payload: Message{Id: messageId}, Result: done, Handle: actions.DestroyMessage}

	// Push the work onto the queue.
	httputils.JobQueue <- work

	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text(result.Data.(string))
}
