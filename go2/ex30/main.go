package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := regexp.MustCompile(`a`)
	fmt.Printf("%#v\n", a.Split("banana", -1))
	fmt.Printf("%#v\n", a.Split("banana", 0))
	fmt.Printf("%#v\n", a.Split("banana", 1))
	fmt.Printf("%#v\n", a.Split("banana", 2))
	fmt.Printf("%#v\n", a.Split("banana", 3))
	fmt.Printf("%#v\n", a.Split("banana", 4))
	zp := regexp.MustCompile(`z+`)
	fmt.Printf("%#v\n", zp.Split("pizza", -1))
	fmt.Printf("%#v\n", zp.Split("pizza", 0))
	fmt.Printf("%#v\n", zp.Split("pizza", 1))
	fmt.Printf("%#v\n", zp.Split("pizza", 2))
	fmt.Printf("%#v\n", zp.Split("pizza", 3))
}
