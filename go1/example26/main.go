/*
greetings라는 배열을 만들고 "good morning", "good afternoon", "good evening", "good night"로 초기화 하세요.
그리고 for 루프를 사용해서 값을 모두 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	greetings := []string{
		"good morning",
		"good afternoon",
		"good evening",
		"good night",
	}
	for i := 0; i < len(greetings); i++ {
		fmt.Println(greetings[i])
	}
}
