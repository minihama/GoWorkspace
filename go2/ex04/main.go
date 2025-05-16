package main

import (
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
)

const JsonString = `
{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}
`

func main() {
	var jsonData interface{}
	if err := json.Unmarshal([]byte(JsonString), &jsonData); err != nil {
		panic(err)
	}
	var lookupStrings = []string{
		"$.expensive",
		"$.store.book[0].price",
		"$.store.book[-1].isbn",
		"$.store.book[0,1].price",
		"$.store.book[1:2].price",
		"$.store.book[?(@.isbn)].price",
		"$.store.book[?(@.price < 10)].price",
	}
	for _, str := range lookupStrings {
		res, err := lookup(jsonData, str)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s => %s\n", str, res)
	}
}

func lookup(data interface{}, expression string) (string, error) {
	res, err := jsonpath.JsonPathLookup(data, expression)
	if err != nil {
		return "", err
	}
	encoded, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(encoded), nil
}
