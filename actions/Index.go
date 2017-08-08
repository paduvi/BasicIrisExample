package actions

import (
	"github.com/valyala/fasthttp"
	"github.com/paduvi/BasicIrisExample/config"
	. "github.com/paduvi/BasicIrisExample/models"
)

func PingRemote(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemoteUrl)

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Data: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Data: string(bodyBytes), Error: nil}
	}
}
