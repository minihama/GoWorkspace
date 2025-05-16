package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Student struct {
	Name   string
	Age    int
	Scores []Score
}

type Score struct {
	Subject string
	Score   int
}

func main() {
	s1 := Student{
		Name: "BloomFilter",
		Age:  21,
		Scores: []Score{
			{
				Subject: "Math",
				Score:   90,
			},
			{
				Subject: "Physics",
				Score:   85,
			},
		},
	}
	spew.Dump(s1)
}
