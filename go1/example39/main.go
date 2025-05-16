/*
왼쪽 예제를 실행하여 결과를 확인하세요.
*/
package main

import (
	"fmt"
)

func main() {
	variadicExample("red")
	variadicExample("yellow", "green", "blue")
}

func variadicExample(s ...string) {
	fmt.Printf("%#v\n", s)
}
