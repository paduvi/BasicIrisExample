package redisutils

import (
	"github.com/garyburd/redigo/redis"
	"github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"regexp"
	"os"
	"fmt"
)

type UserItemPair struct {
	UserId int
	ItemId int
}

func init() {
	conn, err := redis.Dial("tcp", os.Getenv("RedisAddress"))
	if err != nil {
		panic(err)
	}
	psc := redis.PubSubConn{Conn: conn}

	// redis.conf notify-keyspace-events Kx (only subscribe for expired)
	psc.PSubscribe("__keyspace*__:user:item:*", "__keyspace*__:user:history:*")
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.PMessage:
				fmt.Printf("%s \t %s\n", v.Channel, v.Data)
				switch v.Pattern {
				case "__keyspace*__:user:item:*":
					ParseUserItemAndCreateJob(v.Channel)
					break
				case "__keyspace*__:user:history:*":
					ParseUserAndCreateJob(v.Channel)
					break
				}

			case redis.Subscription:
				fmt.Printf("%s \t %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Println(v.Error())
			}
		}
	}()
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

func ParseUserAndCreateJob(channel string) {
	regex, _ := regexp.Compile(":\\d+")

	userId, _ := strconv.Atoi(regex.FindString(channel))
	go func() {
		work := Job{
			Payload: userId,
			Handle:  RemoveExpiredUser,
		}
		JobQueue <- work
	}()
}

func ParseUserItemAndCreateJob(channel string) {
	regex, _ := regexp.Compile(":\\d+")
	matches := regex.FindAllString(channel, 2)
	userId, _ := strconv.Atoi(matches[0])
	itemId, _ := strconv.Atoi(matches[1])

	go func() {
		work := Job{
			Payload: UserItemPair{
				UserId: userId,
				ItemId: itemId,
			},
			Handle: RemoveOldHistory,
		}
		JobQueue <- work
	}()
}

func RemoveOldHistory(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(UserItemPair).UserId
	itemId := payload.(UserItemPair).ItemId

	script := redis.NewScript(2, RemoveOldHistoryScript)
	_, err := script.Do(conn,
		"user:history:"+strconv.Itoa(userId), // keys
		userId, itemId,                       // argv
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

func RemoveExpiredUser(redisPool redis.Pool, done chan models.Result, payload interface{}) {
	conn := redisPool.Get()
	defer conn.Close()

	userId := payload.(int)

	_, err := conn.Do("ZREM", "user:list", userId)
	if err != nil {
		go func() {
			work := Job{
				Payload: payload,
				Handle:  RemoveExpiredUser,
			}
			JobQueue <- work
		}()
	}
}
