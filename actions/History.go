package actions

import (
	"github.com/garyburd/redigo/redis"
	"github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"time"
	"github.com/paduvi/BasicIrisExample/redisutils"
	"github.com/paduvi/BasicIrisExample/cronjob"
)

type UserItemPair struct {
	ItemId int
	UserId int
}

func ListViewer(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	q := payload.(int)
	reply, err := conn.Do("ZRANGEBYSCORE", "user:list", q, "+inf")

	if err != nil {
		done <- models.Result{Error: err}
		return
	}
	done <- models.Result{Error: nil, Data: reply}
}

func ListViewerByItemId(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	itemId := payload.(int)
	reply, err := conn.Do("KEYS", "user:item:*:"+strconv.Itoa(itemId))

	if err != nil {
		done <- models.Result{Error: err}
		return
	}

	done <- models.Result{Error: nil, Data: reply}
}

func ViewItemByUserId(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(UserItemPair).UserId
	itemId := payload.(UserItemPair).ItemId

	script := redis.NewScript(2, redisutils.ViewItemByUserIdScript)
	_, err := script.Do(conn,
		"user:history:"+strconv.Itoa(userId), "user:item:"+strconv.Itoa(userId)+":"+strconv.Itoa(itemId), // keys
		userId, itemId, time.Now().Unix(),                                                                // argv
	)

	if err != nil {
		done <- models.Result{Error: err}
		return
	}
	done <- models.Result{Error: nil}

	go func() {
		work := cronjob.Job{
			Payload: cronjob.UserItemPairWithExpiredTime{UserId: userId, ItemId: itemId, ExpiredTime: time.Now().Add(time.Second * time.Duration(120))},
			Handle:  cronjob.RemoveOldHistory,
		}
		cronjob.JobQueue <- work
	}()
}

func ShowUserHistory(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(int)
	reply, err := conn.Do("ZRANGE", "user:history:"+strconv.Itoa(userId), 0, -1, "WITHSCORES")

	if err != nil {
		done <- models.Result{Error: err}
		return
	}

	done <- models.Result{Error: nil, Data: reply}
}
