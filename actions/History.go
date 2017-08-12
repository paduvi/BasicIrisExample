package actions

import (
	"github.com/garyburd/redigo/redis"
	"github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"time"
	"github.com/paduvi/BasicIrisExample/redisutils"
)

type ViewItemByUserIdPayload struct {
	ItemId int
	UserId int
}

type RemoveOldHistoryPayload struct {
	UserId      int
	ItemId      int
	ExpiredTime time.Time
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
	conn.Send("SETEX", "user:item:"+strconv.Itoa(userId)+":"+strconv.Itoa(itemId), 360, time.Now().Unix())

	replies, err := conn.Do("EXEC")
	if err != nil {
		done <- models.Result{Error: err}
		return
	}
	done <- models.Result{Error: nil, Data: replies}

	work := redisutils.Job{
		Payload: RemoveOldHistoryPayload{UserId: userId, ItemId: itemId, ExpiredTime: time.Now().Add(time.Second * time.Duration(120))},
		Handle:  RemoveOldHistory,
	}
	redisutils.JobQueue <- work
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

func RemoveOldHistory(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	expiredTime := payload.(RemoveOldHistoryPayload).ExpiredTime

	if expiredTime.After(time.Now()) {
		go func() {
			work := redisutils.Job{
				Payload: payload,
				Handle:  RemoveOldHistory,
			}
			redisutils.JobQueue <- work
		}()
		return
	}
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(RemoveOldHistoryPayload).UserId
	itemId := payload.(RemoveOldHistoryPayload).ItemId

	conn.Send("MULTI")
	conn.Send("ZREM", "user:"+strconv.Itoa(userId), itemId)
	conn.Send("DEL", "user:item:"+strconv.Itoa(userId)+":"+strconv.Itoa(itemId))
	conn.Do("EXEC")
}
