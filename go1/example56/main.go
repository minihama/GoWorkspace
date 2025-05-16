package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // WaitGroup 선언

func sum(start, end int) {
	defer wg.Done() // 하나의 작업이 끝날 때 마다 호출
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	fmt.Printf("sum(%d ~ %d) = %d\n", start, end, s)
}

func main() {
	wg.Add(5) // 5개의 작업을 처리해야 함.
	s1 := time.Now().UnixMilli()
	go sum(1234567, 3456789012)
	go sum(2345678, 3456789012)
	go sum(3456789, 3456789012)
	go sum(4567890, 3456789012)
	go sum(5678901, 3456789012)
	wg.Wait() // 모든 작업이 끝나기를 기다림
	fmt.Printf("Elapsed Time : %dms\n", time.Now().UnixMilli()-s1)
}
