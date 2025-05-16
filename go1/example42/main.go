package main

import (
	"github.com/davecgh/go-spew/spew"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
	geo     geometry
}

type geometry struct {
	area      float64
	perimeter float64
}

func main() {
	var rect1 rectangle
	rect1.length = 10
	rect1.breadth = 15
	rect1.geo.area = rect1.length * rect1.breadth
	spew.Dump(rect1)

	rect2 := rectangle{10, 20, "Green", geometry{200, 60}}
	spew.Dump(rect2)

	rect3 := rectangle{length: 5, breadth: 10}
	spew.Dump(rect3)

	rect4 := new(rectangle)
	rect4.length = 5
	rect4.breadth = 12
	spew.Dump(rect4)
}
