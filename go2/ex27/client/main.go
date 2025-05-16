package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func main() {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(100 * time.Millisecond).
		SetRetryMaxWaitTime(5 * time.Second).
		AddRetryCondition(func(resp *resty.Response, err error) bool {
			return err != nil || resp.StatusCode() != http.StatusOK
		})
	resp, err := client.R().
		Get("http://localhost:1234/")
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}
