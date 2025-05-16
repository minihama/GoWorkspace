/*
1 ~ 200까지 홀수만 더하여 총합을 구하세요
*/
package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 200; i += 2 {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
