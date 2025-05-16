/**
  a, b, c, d를 2,3,4,5로 초기화하세요. 그리고 다음의 값이
얼마인지 출력해 보세요. (b >= a) && !(c < d)
*/
package main

import (
	"fmt"
)

func main() {
	a, b, c, d := 2, 3, 4, 5
	fmt.Println((b >= a) && !(c < d))
}