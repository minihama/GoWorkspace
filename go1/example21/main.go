/**
  age 변수를 19로 초기화 하세요. 그리고 그 변수가 10대
이하면 "10대 이하", 20대 이면 "20대", 30대 이면 "30대",
40대 이상이면 "40대 이상"을 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	age := 19
	if age < 20 {
		fmt.Println("10대 이하")
	} else if age < 30 {
		fmt.Println("20대")
	} else if age < 40 {
		fmt.Println("30대")
	} else {
		fmt.Println("40대 이상")
	}
}
