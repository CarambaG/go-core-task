package main

import (
	"context"
	"fmt"
	"time"
)

func MergeChannels[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		done := make(chan struct{}, len(channels))

		for _, ch := range channels {
			go func(input <-chan T) {
				defer func() { done <- struct{}{} }()
				for {
					select {
					case <-ctx.Done():
						return
					case v, ok := <-input:
						if !ok {
							return
						}
						select {
						case <-ctx.Done():
							return
						case out <- v:
						}
					}
				}
			}(ch)
		}

		for i := 0; i < len(channels); i++ {
			select {
			case <-ctx.Done():
				return
			case <-done:
			}
		}
	}()
	return out
}

// Генератор строк для тестов
func stringGenerator(ctx context.Context, id string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- fmt.Sprintf("%s-%d", id, i):
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 3 канала с разными данными
	ch1 := stringGenerator(ctx, "A", 5)
	ch2 := stringGenerator(ctx, "B", 3)
	ch3 := stringGenerator(ctx, "C", 4)

	merged := MergeChannels(ctx, ch1, ch2, ch3)

	for msg := range merged {
		fmt.Println(msg)
	}
}
