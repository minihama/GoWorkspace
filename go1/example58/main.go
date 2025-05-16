package main

import (
	"fmt"
)

func main() {
	finished := make(chan bool) // 채널 생성
	go func() {
		fmt.Println("Hello World")
		finished <- true // 채널에 데이터를 보냄
	}()
	<-finished // 채널에서 데이터를 받을 때까지 block 됨
	fmt.Println("Program Ended")
}
