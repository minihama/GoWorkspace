/*
numbers 슬라이스를 만들고, 100부터 120까지 정수를 numbers에 추가 하세요.
그리고 그 합을 구하여 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	sum := 0
	numbers := make([]int, 0)
	for i := 100; i <= 120; i++ {
		numbers = append(numbers, i)
	}
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}
