/**
  x 변수를 true로, y변수를 "nice"로, z 변수를 3.14159로
초기화 한 다음에 %T와 %v를 사용하여 각 변수를 모두 출력해
보세요.
*/
package main

import (
	"fmt"
)

func main() {
	x, y, z := true, "nice", 3.14159
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}