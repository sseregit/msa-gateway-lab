package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-gateway/common"
	"go-gateway/config"
	"go-gateway/kafka"
)

const (
	_defaultBatchTime = 2
)

type HttpClient struct {
	client *resty.Client
	cfg    config.App

	producer kafka.Producer
}

func NewHttpClient(
	cfg config.App,
	producer map[string]kafka.Producer,
) HttpClient {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	if cfg.Http.BaseUrl == "" {
		panic("Base url not existed")
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal).
		SetBaseURL(cfg.Http.BaseUrl)

	return HttpClient{
		cfg:      cfg,
		client:   client,
		producer: producer[cfg.App.Name],
	}
}

func (h HttpClient) GET(url string, router config.Router) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	req = getRequest(h.client, router)
	if resp, err = req.Get(url); err != nil {
		return nil, err
	} else {
		fmt.Println(resp, url)
		return string(resp.Body()), nil
	}

}
func (h HttpClient) POST(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	req = getRequest(h.client, router).SetBody(requestBody)
	if resp, err = req.Post(url); err != nil {
		return nil, err
	} else {
		return string(resp.Body()), nil
	}
}
func (h HttpClient) DELETE(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	req = getRequest(h.client, router).SetBody(requestBody)
	if resp, err = req.Delete(url); err != nil {
		return nil, err
	} else {
		return string(resp.Body()), nil
	}
}
func (h HttpClient) PUT(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	req = getRequest(h.client, router).SetBody(requestBody)
	if resp, err = req.Put(url); err != nil {
		return nil, err
	} else {
		return string(resp.Body()), nil
	}
}

func getRequest(client *resty.Client, router config.Router) *resty.Request {
	req := client.R().EnableTrace()

	if router.Auth != nil {
		if len(router.Auth.Key) != 0 {
			req.SetAuthScheme(router.Auth.Key)
		}
		req.SetAuthToken(router.Auth.Token)
	}

	if router.Header != nil {
		req.SetHeaders(router.Header)
	}

	return req
}
