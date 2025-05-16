/**
  문자열 포인터를 선언하여 왼쪽 예제와 같은 방식으로
프로그램을 작성하여, "hello"에서 "world"로 변환되는 과정을
각각 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	var x = "hello"
	var p *string
	p = &x
	fmt.Println(x)
	*p = "world"
	fmt.Println(x)
}
