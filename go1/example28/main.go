/*
왼쪽 슬라이스를 사용하여, 원소의 개수가 2개씩만 가지도록 슬라이싱을 하여 출력하세요.
*/
package main

import (
	"fmt"
)

func main() {
	alphabet := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(alphabet)-1; i++ {
		fmt.Println(alphabet[i : i+2])
	}
}
