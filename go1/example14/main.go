/**
  x 변수를 1234567890123456.123으로 초기화 후에
"[%g]"와 "[%f]"와 "[%e]"로 출력 해 보세요.
*/
package main

import (
	"fmt"
)

func main() {
	x := 1234567890123456.123
	fmt.Printf("[%g] [%f] [%e]\n", x, x, x)
}