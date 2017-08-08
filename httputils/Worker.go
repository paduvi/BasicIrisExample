package httputils

import (
	"github.com/valyala/fasthttp"
	"net"
	"time"
	"strconv"
	"os"
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	HttpClient *fasthttp.Client
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		HttpClient: &fasthttp.Client{
			Dial: func(addr string) (net.Conn, error) {
				RequestTimeOut, _ := strconv.Atoi(os.Getenv("RequestTimeOut"))
				KeepAliveDuration, _ := strconv.Atoi(os.Getenv("KeepAliveDuration"))
				dialer := net.Dialer{
					Timeout:   time.Duration(RequestTimeOut) * time.Second,
					KeepAlive: time.Duration(KeepAliveDuration) * time.Second,
				}
				return dialer.Dial("tcp", addr)
			},
		},
		quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				job.Handle(w.HttpClient, job.Result, job.Payload)

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
