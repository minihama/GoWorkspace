/*
연습문제 24를 왼쪽 예제처럼 작성하세요.
*/
package main

import (
	"fmt"
)

func main() {
	sum, i := 0, 1
	for i <= 200 {
		sum += i
		i += 2
	}
	fmt.Println("sum = ", sum)
}
