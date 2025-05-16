package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sum(start, end int, sum *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	t := 0
	for i := start; i <= end; i++ {
		t += i
	}
	atomic.AddUint64(sum, uint64(t))
}

func main() {
	var s = uint64(0)
	var wg sync.WaitGroup
	wg.Add(5)
	go sum(1, 10000000, &s, &wg)
	go sum(10000001, 20000000, &s, &wg)
	go sum(20000001, 30000000, &s, &wg)
	go sum(30000001, 40000000, &s, &wg)
	go sum(40000001, 50000000, &s, &wg)
	wg.Wait()
	fmt.Println(atomic.LoadUint64(&s))
}
