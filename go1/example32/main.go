/*
두 정수를 입력 받아 가감승제 한 결과 4가지 정수를 반환하는 함수를 작성하세요.
*/
package main

import (
	"fmt"
)

func main() {
	a, b, c, d := fourArithmetic(11, 5)
	fmt.Println(a, b, c, d)
}

func fourArithmetic(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
