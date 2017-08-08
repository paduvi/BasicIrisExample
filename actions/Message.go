package actions

import (
	"github.com/valyala/fasthttp"
	"github.com/paduvi/BasicIrisExample/config"
	. "github.com/paduvi/BasicIrisExample/models"
	"strconv"
	"encoding/json"
)

func ListMessage(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemoteUrl + "/messages")

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Data: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Data: bodyBytes, Error: nil}
	}
}

func FindMessage(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	messageId := payload.(Message).Id
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemoteUrl + "/messages/" + strconv.Itoa(messageId))

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Data: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Data: bodyBytes, Error: nil}
	}
}

func CreateMessage(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemoteUrl + "/messages")
	req.Header.SetMethod("POST")
	body, err := json.Marshal(payload)

	if err != nil {
		done <- Result{Data: nil, Error: err}
		return;
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Data: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Data: bodyBytes, Error: nil}
	}
}

func DestroyMessage(client *fasthttp.Client, done chan Result, payload interface{}) {
	defer func() {
		close(done)
	}()
	messageId := payload.(Message).Id
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(config.RemoteUrl + "/messages/" + strconv.Itoa(messageId))
	req.Header.SetMethod("DELETE")

	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		done <- Result{Data: nil, Error: err}
	} else {
		bodyBytes := resp.Body()
		done <- Result{Data: string(bodyBytes), Error: nil}
	}
}
