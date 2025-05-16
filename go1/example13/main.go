/**
  x 변수를 31로 초기화 한 후에, x 변수를  2진수로도,
8진수로도, 16진수로도, 10진수로도 출력하세요.
y 변수를 12로 초기화 후에 "[%03d]"로도 "[%3d]"로도
출력해 보세요.
*/
package main

import (
	"fmt"
)

func main() {
	x := 31
	fmt.Printf("%b %o %x %d\n", x, x, x, x)
	y := 12
	fmt.Printf("[%03d] [%3d]\n", y, y)
}
