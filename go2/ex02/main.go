package main

import (
	"github.com/davecgh/go-spew/spew"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	encoded, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}
	spew.Dump(string(encoded))

	json := jsoniter.ConfigFastest
	decoded := make(map[string]int)
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		panic(err)
	}
	spew.Dump(decoded)
}
