package main

import (
	"fmt"
	"time"
)

func sum(start, end int) {
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	fmt.Printf("sum(%d ~ %d) = %d\n", start, end, s)
}

func main() {
	s1 := time.Now().UnixMilli()
	go sum(1234567, 3456789012)
	go sum(2345678, 3456789012)
	go sum(3456789, 3456789012)
	go sum(4567890, 3456789012)
	go sum(5678901, 3456789012)
	time.Sleep(2 * time.Second)
	fmt.Printf("Elapsed Time : %dms\n", time.Now().UnixMilli()-s1)
}
