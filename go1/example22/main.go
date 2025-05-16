/*
왼쪽 예제를 if/else 로 변경하여 작성하세요. 단, x는 3으로
초기화 하세요.
*/
package main

import (
	"fmt"
)

func main() {
	x := 3
	if x%4 == 1 {
		fmt.Println("나머지는 1입니다")
	} else if x%4 == 2 || x%4 == 3 {
		fmt.Println("나머지는 2 또는 3입니다")
	} else {
		fmt.Println("나머지는 0입니다.")
	}
}
