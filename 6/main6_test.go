package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := RandomGenerator(ctx, 100)

	count := 0
	for range ch {
		count++
		if count >= 5 {
			cancel()
			break
		}
	}

	select {
	case _, ok := <-ch:
		if ok {
			t.Error("Channel not closed after context cancel")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Channel not closed within timeout")
	}
}

func TestRandomGenerator_RangeLoop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	ch := RandomGenerator(ctx, 50)
	numbers := 0

	for range ch {
		numbers++
	}

	if numbers == 0 {
		t.Error("No numbers generated")
	}
	t.Logf("Generated %d numbers", numbers)
}

func TestRandomGenerator_RangeMax(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := RandomGenerator(ctx, 10)
	valid := true

	for num := range ch {
		if num < 0 || num > 10 {
			valid = false
			t.Errorf("Out of range: %d", num)
			break
		}
	}

	if !valid {
		t.Fail()
	}
}

func TestRandomGenerator_UnbufferedBlocking(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := RandomGenerator(ctx, 100)

	for i := 0; i < 3; i++ {
		<-ch
	}

	cancel()
	time.Sleep(10 * time.Millisecond)

	_, ok := <-ch
	if ok {
		t.Error("Channel not closed after context cancel")
		t.Fail()
	}
}
