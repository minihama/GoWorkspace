/**
  x 변수를 논리형 타입으로 선언 한 후에 값의 초기화 없이 그
값을 출력하세요. x 변수를 true로 대입한 후에 다시 그 값을
출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	var x bool
	fmt.Printf("%t\n", x)
	x = true
	fmt.Printf("%t\n", x)
}
