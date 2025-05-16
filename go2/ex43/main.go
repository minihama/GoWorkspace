package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("args is missed")
	}
	validate := validator.New()
	if errs := validate.Var(args[0], "required,email"); errs != nil {
		panic(errs)
	}
	fmt.Printf(args[0])
}

// go run main.go email
// go run main.go email@address.com
