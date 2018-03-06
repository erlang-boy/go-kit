package sdk

import (
	"time"

	"github.com/go-resty/resty"
)

var client = resty.New()

func NewClient(timeout time.Duration) *resty.Client {
	return client.
		SetRetryCount(3).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(2 * time.Second).
		SetTimeout(timeout)
}
