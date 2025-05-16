/*
연습문제 21을 switch를 사용하되, switch 다음에 표현식
없이 작성하세요.
*/
package main

import (
	"fmt"
)

func main() {
	age := 19
	switch {
	case age < 20:
		fmt.Println("10대 이하")
	case age < 30:
		fmt.Println("20대")
	case age < 40:
		fmt.Println("30대")
	default:
		fmt.Println("40대 이상")
	}
}
