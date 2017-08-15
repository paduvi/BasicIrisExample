package cronjob

import (
	"github.com/garyburd/redigo/redis"
	"github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"time"
	"regexp"
	"github.com/paduvi/BasicIrisExample/redisutils"
)

type UserItemPairWithExpiredTime struct {
	UserId      int
	ItemId      int
	ExpiredTime time.Time
}

func init() {
	// Get all user:item:* keys then parse to UserItemPairWithExpiredTime and push to queue
	done := make(chan models.Result)
	work := Job{
		Result: done,
		Handle: GetAllUserItemPairKeys,
	}

	// Push the work onto the queue.
	JobQueue <- work
	result := <-done

	for _, key := range result.Data.([]interface{}) {
		go func() {
			work := Job{
				Payload: string(key.([]byte)),
				Handle:  ParseUserItemAndCreateJob,
			}
			JobQueue <- work
		}()
	}
}

func GetAllUserItemPairKeys(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	// Get all user:item:* keys
	keys, err := conn.Do("KEYS", "user:item:*")

	if err != nil {
		done <- models.Result{Error: err}
		return
	}

	done <- models.Result{Data: keys, Error: nil}
}

func ParseUserItemAndCreateJob(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	key := payload.(string)

	regex, _ := regexp.Compile(":\\d+")
	matches := regex.FindAllString(key, 2)
	userId, _ := strconv.Atoi(matches[0])
	itemId, _ := strconv.Atoi(matches[1])

	result, err := conn.Do("GET", key)

	if err != nil {
		go func() {
			work := Job{
				Payload: key,
				Handle:  ParseUserItemAndCreateJob,
			}
			JobQueue <- work
		}()
		return
	}

	timestamp, _ := strconv.ParseInt(string(result.([]byte)), 10, 64)

	go func() {
		work := Job{
			Payload: UserItemPairWithExpiredTime{
				UserId:      userId,
				ItemId:      itemId,
				ExpiredTime: time.Unix(timestamp, 0),
			},
			Handle: RemoveOldHistory,
		}
		JobQueue <- work
	}()
}

func RemoveOldHistory(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	expiredTime := payload.(UserItemPairWithExpiredTime).ExpiredTime

	if expiredTime.After(time.Now()) {
		go func() {
			work := Job{
				Payload: payload,
				Handle:  RemoveOldHistory,
			}
			JobQueue <- work
		}()
		return
	}
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(UserItemPairWithExpiredTime).UserId
	itemId := payload.(UserItemPairWithExpiredTime).ItemId

	script := redis.NewScript(2, redisutils.RemoveOldHistoryScript)
	_, err := script.Do(conn,
		"user:history:"+strconv.Itoa(userId), "user:item:"+strconv.Itoa(userId)+":"+strconv.Itoa(itemId), // keys
		userId, itemId,                                                                                   // argv
	)

	if err != nil {
		go func() {
			work := Job{
				Payload: payload,
				Handle:  RemoveOldHistory,
			}
			JobQueue <- work
		}()
	}
}
