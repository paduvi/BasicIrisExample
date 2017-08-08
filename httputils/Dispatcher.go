package httputils

import (
	"os"
	"strconv"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func init() {
	MaxWorker, _ := strconv.Atoi(os.Getenv("MaxWorker"))
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < cap(d.WorkerPool); i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
