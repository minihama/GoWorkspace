/*
왼쪽 설명대로 실행하여 seafood  안에  foo 가 포함되어 있는지 출력하세요.
*/
package main

import (
	"fmt"

	"example36/src/string_util"
)

func main() {
	fmt.Println(string_util.MyContains("seafood", "foo"))
}
