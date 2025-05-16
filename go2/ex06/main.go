package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/shamaton/msgpack/v2"
)

func main() {
	type Struct struct {
		String string
	}
	v := Struct{String: "msgpack"}

	d, err := msgpack.Marshal(v)
	if err != nil {
		panic(err)
	}
	spew.Dump(d)
	r := Struct{}
	err = msgpack.Unmarshal(d, &r)
	if err != nil {
		panic(err)
	}
	spew.Dump(r)
}
