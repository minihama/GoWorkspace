package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/shamaton/msgpack/v2"
)

type Student struct {
	ID           int
	Name         string
	ScoreMath    int
	ScoreEnglish int
}

func main() {
	data := []Student{
		{ID: 1, Name: "Kim", ScoreMath: 78, ScoreEnglish: 89},
		{ID: 2, Name: "Lee", ScoreMath: 100, ScoreEnglish: 92},
		{ID: 3, Name: "Park", ScoreMath: 68, ScoreEnglish: 96},
		{ID: 4, Name: "Choi", ScoreMath: 92, ScoreEnglish: 94},
	}
	ser1, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	spew.Dump(ser1)
	ser2, err := msgpack.MarshalAsArray(data)
	if err != nil {
		panic(err)
	}
	spew.Dump(ser2)
	unser := make([]Student, 0)
	if err := msgpack.UnmarshalAsArray(ser2, &unser); err != nil {
		panic(err)
	}
	spew.Dump(unser)
}
