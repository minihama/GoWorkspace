package main

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func main() {
	jsonUrl := "https://jsonplaceholder.typicode.com/posts/1"
	parsedUrl, err := url.Parse(jsonUrl)
	if err != nil {
		panic(err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	jar.SetCookies(parsedUrl, []*http.Cookie{
		{
			Name:    "test_cookie",
			Value:   "testCookie",
			Path:    "/",
			Expires: time.Now().Add(time.Hour),
		},
	})

	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			IdleConnTimeout:   10 * time.Second,
		},
		CheckRedirect: noRedirect,
		Jar:           jar,
		Timeout:       time.Minute,
	}
	requestBody, err := json.Marshal(map[string]interface{}{
		"id":     1,
		"userId": 1,
		"title":  "new title",
		"bdoy":   "new body",
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, jsonUrl, bytes.NewReader(requestBody))
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		panic("HTTP ERROR: " + res.Status)
	}
	post := make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&post); err != nil {
		panic(err)
	}
	spew.Dump(post)
}

func noRedirect(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}
