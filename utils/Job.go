package utils

import "github.com/paduvi/BasicIrisExample/config"
import "github.com/paduvi/BasicIrisExample/models"

// Job represents the job to be run
type Job struct {
	Payload interface{}
	Result  chan models.Result
}

// A buffered channel that we can send work requests on.
var JobQueue = make(chan Job, config.MaxQueue)
