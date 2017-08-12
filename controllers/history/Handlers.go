package todo

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris"
	"strconv"
	"github.com/paduvi/BasicIrisExample/models"
	"github.com/paduvi/BasicIrisExample/actions"
	"github.com/paduvi/BasicIrisExample/redisutils"
	"regexp"
	"strings"
	"time"
)

func ListViewer(ctx context.Context) {
	q, err := strconv.Atoi(ctx.Request().URL.Query().Get("q"))
	if err != nil {
		q = 1
	}
	if q < 1 {
		ctx.Redirect(ctx.Request().URL.Path)
		return
	}

	done := make(chan models.Result)
	work := redisutils.Job{
		Payload: q,
		Result:  done,
		Handle:  actions.ListViewer,
	}

	// Push the work onto the queue.
	redisutils.JobQueue <- work
	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	data := []int{}
	for _, element := range result.Data.([]interface{}) {
		n, _ := strconv.Atoi(string(element.([]byte)))
		data = append(data, n)
	}
	ctx.JSON(data)
}

func ListViewerByItemId(ctx context.Context) {
	itemId, err := ctx.Params().GetInt("itemId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan models.Result)
	work := redisutils.Job{
		Payload: itemId,
		Result:  done,
		Handle:  actions.ListViewerByItemId,
	}

	// Push the work onto the queue.
	redisutils.JobQueue <- work
	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	data := []int{}
	regex, err := regexp.Compile(":\\d+:")
	for _, element := range result.Data.([]interface{}) {
		userId, _ := strconv.Atoi(strings.Trim(string(regex.Find(element.([]byte))), ":"))
		data = append(data, userId)
	}
	ctx.JSON(data)
}

func ViewItemByUserId(ctx context.Context) {
	userId, err := ctx.Params().GetInt("userId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	itemId, err := ctx.Params().GetInt("itemId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan models.Result)
	work := redisutils.Job{
		Payload: actions.UserItemPair{UserId: userId, ItemId: itemId},
		Result:  done,
		Handle:  actions.ViewItemByUserId,
	}

	// Push the work onto the queue.
	redisutils.JobQueue <- work
	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.Text("OK")
}

func ShowUserHistory(ctx context.Context) {
	userId, err := ctx.Params().GetInt("userId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan models.Result)
	work := redisutils.Job{
		Payload: userId,
		Result:  done,
		Handle:  actions.ShowUserHistory,
	}

	// Push the work onto the queue.
	redisutils.JobQueue <- work
	result := <-done
	if result.Error != nil {
		ctx.Values().Set("error", result.Error.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	list := result.Data.([]interface{})
	histories := map[int]time.Time{}
	for i := 0; i < len(list); i += 2 {
		itemId, _ := strconv.Atoi(string(list[i].([]byte)))
		timeStamp, _ := strconv.ParseInt(string(list[i+1].([]byte)), 10, 64)
		histories[itemId] = time.Unix(timeStamp, 0)
	}
	ctx.JSON(models.User{Id: userId, Histories: histories})
}
