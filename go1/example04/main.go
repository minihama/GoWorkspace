/**
  x 변수를 uint8로 선언 한 후에 200으로 초기화하세요. x에
56을 더 한 후에 출력 해 보세요. y 변수를 int8로 선언한 후에
-128로 초기화 하세요. 그 다음에 3을 뺀 값을 출력 해 보세요.
*/
package main

import (
	"fmt"
)

func main() {
	var x uint8 = 200
	fmt.Printf("x = %d\n", x+56) // overflow
	var y int8 = -128
	fmt.Printf("y = %d\n", y-3) // underflow
}
