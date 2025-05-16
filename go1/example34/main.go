/*
내장 패키지의 strings를 사용하여 "seafood" 문자열이 "foo"를 포함하고 있는지 그 결과를 출력하세요.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "seafood"
	fmt.Println(strings.Contains(str, "foo"))
	fmt.Println(strings.Contains(str, "bar"))
}
