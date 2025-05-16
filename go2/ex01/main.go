package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	encoded, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	spew.Dump(string(encoded))
	decoded := make(map[string]int)
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		panic(err)
	}
	spew.Dump(decoded)
}
