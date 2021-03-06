package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Millisecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
