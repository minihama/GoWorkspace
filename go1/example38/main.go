/*
왼쪽 예제를 실행하여 출력된 에러 메시지를 살펴 보세요.
*/
package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := fn()
	fmt.Printf("%+v", err)
}

func fn() error {
	e1 := errors.New("my error")
	e2 := errors.Wrap(e1, "inner")
	return errors.Wrap(e2, "outer")
}
