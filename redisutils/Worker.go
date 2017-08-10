package redisutils

import (
	"github.com/garyburd/redigo/redis"
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	RedisPool  redis.Pool
	quit       chan bool
}

func NewWorker(workerPool chan chan Job, redisPool redis.Pool) Worker {
	return Worker{
		WorkerPool: workerPool,
		RedisPool:  redisPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				job.Handle(w.RedisPool, job.Result, job.Payload)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
