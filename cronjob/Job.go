package cronjob

import (
	"github.com/paduvi/BasicIrisExample/models"
	"os"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

// Job represents the job to be run
type Job struct {
	Payload interface{}
	Result  chan models.Result
	Handle  func(redis.Pool, chan models.Result, interface{})
}

// A buffered channel that we can send work requests on.
var maxQueue, _ = strconv.Atoi(os.Getenv("MaxQueue"))
var JobQueue = make(chan Job, maxQueue)
