package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomthing(ctx)
}

func doSomthing(ctx context.Context) {
	fmt.Println(ctx.Value("myKey"))
}
