/**
  x를 7로 y를 2로 초기화하세요.  그리고
x & y, x | y, x ^ y, x << y, x >> y 의 값을 각각 구하세요.
*/
package main

import (
	"fmt"
)

func main() {
	x, y := 7, 2
	fmt.Println("x & y =", x&y)
	fmt.Println("x | y =", x|y)
	fmt.Println("x ^ y = ", x^y)
	fmt.Println("x << y = ", x<<y)
	fmt.Println("x >> y = ", x>>y)
}
