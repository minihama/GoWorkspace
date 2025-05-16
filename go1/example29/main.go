/*
왼쪽의 myArray의 총합을 구하여 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	sum := 0
	myArray := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := 0; i < len(myArray); i++ {
		for j := 0; j < len(myArray[i]); j++ {
			sum += myArray[i][j]
		}
	}
	fmt.Println(sum)
}
