package main

import (
	"github.com/davecgh/go-spew/spew"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	spew.Dump(rectangle{10.5, 12.5, "red"})
}
