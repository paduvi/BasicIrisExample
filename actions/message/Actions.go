package message

import (
	. "github.com/paduvi/BasicIrisExample/models"
	"github.com/valyala/fasthttp"
	"github.com/paduvi/BasicIrisExample/config"
)

func PingMessage(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemotePingUrl)

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Response: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Response: string(bodyBytes), Error: nil}
	}
}
