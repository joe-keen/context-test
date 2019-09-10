package main

import (
	"context"
	"fmt"
	"github.com/joe-keen/context-test/test"
)

func main() {
	ctx := context.Background()

	ctx = test.Set(ctx)

	fmt.Println(test.Get(ctx))
	fmt.Println(ctx.Value("foo"))
}
