package actions

import (
	"github.com/garyburd/redigo/redis"
	"github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"time"
)

type ViewItemByUserIdPayload struct {
	ItemId int
	UserId int
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

	userId := payload.(ViewItemByUserIdPayload).UserId
	itemId := payload.(ViewItemByUserIdPayload).ItemId

	conn.Send("MULTI")
	conn.Send("ZADD", "user:"+strconv.Itoa(userId), time.Now().Unix(), itemId)
	conn.Send("EXPIRE", "user:"+strconv.Itoa(userId), 120)
	conn.Send("SETEX", "user:item:"+strconv.Itoa(userId)+":"+strconv.Itoa(itemId), 60, time.Now().Unix())

	replies, err := conn.Do("EXEC")
	if err != nil {
		done <- models.Result{Error: err}
		return
	}
	done <- models.Result{Error: nil, Data: replies}
}

func ShowUserHistory(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(int)
	reply, err := conn.Do("ZRANGE", "user:"+strconv.Itoa(userId), 0, -1, "WITHSCORES")

	if err != nil {
		done <- models.Result{Error: err}
		return
	}

	done <- models.Result{Error: nil, Data: reply}
}
