/**
  a, b, c, d 변수를 문자열형 타입으로 선언 한 후에 "nice",
"to", "meet", "you" 로 초기화 후에 각각 "[%3s]" 로
출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	a, b, c, d := "nice", "to", "meet", "you"
	fmt.Printf("[%3s] [%3s] [%3s] [%3s]\n", a, b, c, d)
}
