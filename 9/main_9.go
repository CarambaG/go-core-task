package main

import (
	"fmt"
	"time"
)

func CubePipeline(in <-chan uint8) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for num := range in {
			f := float64(num)
			cube := f * f * f
			out <- cube
		}
	}()

	return out
}

func uint8Generator(count int) <-chan uint8 {
	ch := make(chan uint8)

	go func() {
		defer close(ch)
		for i := uint8(0); i < uint8(count); i++ {
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}()

	return ch
}

func main() {
	input := uint8Generator(10)

	cubed := CubePipeline(input)

	i := 0
	for cube := range cubed {
		fmt.Printf("%d^3 = %.2f\n", i, cube)
		i++
	}
}
