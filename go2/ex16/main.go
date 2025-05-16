package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

func main() {
	mc := memcache.New("192.168.0.16:11211")
	err := mc.Set(&memcache.Item{
		Key:        "foo",
		Value:      []byte("Hello World"),
		Expiration: 15 * 60,
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(it.Value))
}
