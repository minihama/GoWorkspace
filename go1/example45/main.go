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
	r1 := rectangle{5, 4, "yellow"}
	r2 := r1
	r3 := &r1
	r2.color = "blue"
	spew.Dump(r1)
	r3.color = "magenta"
	spew.Dump(r1)
}
