/*
정수를 가변 인자로 받는 sum 함수를 작성하여 누적합을 반환하는 함수로 작성하세요.
그리고 main 함수에서 1, 3, 5, 7, 9를 입력하여 그 sum값이 얼마인지 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 5, 7, 9))
}

func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
