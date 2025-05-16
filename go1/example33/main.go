/*
왼쪽 예제 프로그램을 그대로 작성하여 결과값을 확인 해 보세요.
*/
package main

func main() {
	num, str := 4, "hello"
	println(num, ":", str)
	change(&num, &str)
	println(num, ":", str)
}

func change(n *int, s *string) {
	*n *= 2
	*s += " world"
}
