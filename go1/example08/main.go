/**
  x 변수를 uint8로 선언 한 후에 200으로 초기화하세요. y
변수는 int8 로 선언한 후에 -100으로 초기화 하세요. z
변수를 int 형으로 선언한 후에 x + y 의 값을 대입한 후에
출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	var x uint8 = 200 // x := uint8(200)
	var y int8 = -100 // y := int8(-100)
	z := int(x) + int(y)
	fmt.Printf("%d\n", z)
}
