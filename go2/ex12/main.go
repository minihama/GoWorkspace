package main

import (
	"fmt"
	"sync"
)

func main() {
	var cm sync.Map
	cm.Store("c1", 689)
	cm.Store("c2", "Have a nice day")
	cm.Store(123, 567)
	if v, ok := cm.Load("c2"); ok {
		fmt.Printf("c2: %s\n", v)
	}
}
