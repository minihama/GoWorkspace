package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var vary interface{}
	vary = 123
	spew.Dump(vary)
	vary = "good morning"
	spew.Dump(vary)
	vary = map[string]int{"Mark": 10, "Jane": 20}
	spew.Dump(vary)
	vary = [3]string{"Korea", "Japan", "China"}
	spew.Dump(vary)
}
