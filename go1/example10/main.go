/**
  x 변수를 문자열형 "200"으로 초기화 한 후 int 형으로 캐스팅
한 후에 출력하세요. y 변수는 float64 로 선언한 후에
-100.0으로 초기화 하세요. 문자열형으로 캐스팅 한 후에
출력하세요.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "200"
	intX, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", intX)
	y := -100.0
	strY := strconv.FormatFloat(y, 'f', -1, 64)
	fmt.Printf("%s\n", strY)
}
