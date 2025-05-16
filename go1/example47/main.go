package main

import (
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Kim", 27},
		{"Lee", 43},
		{"Jung", 36},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	spew.Dump(people)
}
