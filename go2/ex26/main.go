package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func main() {
	client := resty.New().SetDebug(true)
	resp, err := client.R().
		EnableTrace().
		Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() != http.StatusOK {
		panic("HTTP ERROR: " + resp.Status())
	}
	spew.Dump(resp.Request.TraceInfo())
	spew.Dump(resp)
}
