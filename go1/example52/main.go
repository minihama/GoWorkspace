package main

import (
	"fmt"
)

type Printer interface {
	Print()
}
type Emp struct {
	name, address string
}

func (e *Emp) Print() {
	fmt.Printf("Name: %s\nAddress: %s\n", e.name, e.address)
}
func (e *Emp) Assign(n, a string) {
	e.name = n
	e.address = a
}

func main() {
	var e Emp
	e.Assign("John Doe", "very long address string")
	var p Printer
	p = &e
	p.Print()
}
