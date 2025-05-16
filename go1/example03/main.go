/**
x 변수는 5로 초기화 한 후에, y = 2x + 1의 값을 계산하여
출력하는 프로그램을 작성하여 실행하세요.
*/
package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 2*x + 1
	fmt.Printf("x = %d, y = %d\n", x, y)
}
