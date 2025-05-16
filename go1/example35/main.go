/*
str 변수를 "hello"로 초기화 하세요. 그리고 str의 주소를 p에 저장하세요.
그리고 spew.Dump 함수를 사용하여 그 주소값을 출력하세요.
*/
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	x := "hello world"
	y := &x
	fmt.Printf("%#v\n", y)
	spew.Dump(y)
}
