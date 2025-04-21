package client

import (
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
) {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal)
}
