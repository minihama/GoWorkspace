package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func main() {
	m := cmap.New[any]()
	m.Set("a1", 1234)
	m.Set("a2", "morning")
	m.Set("a3", []int{2, 3, 4, 5})
	if v, ok := m.Get("a2"); ok {
		fmt.Println(v)
	}
	m.IterCb(func(key string, val any) {
		fmt.Printf("%s: %#v\n", key, val)
	})
}
