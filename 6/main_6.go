package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func RandomGenerator(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		rand.Seed(time.Now().UnixNano())
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- rand.Intn(max + 1):
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := RandomGenerator(ctx, 100)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
