package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func main() {
	client := http.Client{}
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
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
