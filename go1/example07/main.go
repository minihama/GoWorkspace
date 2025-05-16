/**
  b 변수를 byte 타입으로 'B'로 초기화 선언하세요. R 변수를
'한'으로 초기화 선언하세요. 그리고 이 값들을 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	var b byte = 'B'
	fmt.Printf("b = %s\n", string(b))
	R := '한'
	fmt.Printf("R = %s\n", string(R))
}
