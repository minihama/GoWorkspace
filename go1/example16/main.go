/**
  x 변수를 4로 초기화 한 뒤에 2를 증가시키세요. y 변수는
3으로 초기화 한 뒤에 1을 감소시키세요. x % y 한 값을
출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	x := 4
	x += 2
	y := 3
	y--
	fmt.Println(x % y)
}