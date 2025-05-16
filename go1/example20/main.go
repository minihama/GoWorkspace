/**
  a, b를 4와 3으로 초기화하세요. a가 b보다 크면, "크다"를
출력하고, 그렇지 않으면 "작다"를 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	a, b := 4, 3
	if a > b {
		fmt.Println("크다")
	} else {
		fmt.Println("작다")
	}
}