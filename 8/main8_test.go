package main

import (
	"sync"
	"testing"
	"time"
)

func TestCustomWaitGroup_Basic(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()

	start := time.Now()
	wg.Wait()
	duration := time.Since(start)

	if duration < 40*time.Millisecond {
		t.Error("Wait returned too early")
	}
}

func TestCustomWaitGroup_MultipleGoroutines(t *testing.T) {
	wg := NewCustomWaitGroup()
	n := 10

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestCustomWaitGroup_ZeroInitial(t *testing.T) {
	wg := NewCustomWaitGroup()

	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait on zero count blocked")
	}
}

func TestCustomWaitGroup_NegativePanic(t *testing.T) {
	wg := NewCustomWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Done without Add should panic")
		}
	}()
	wg.Done()
}

func TestCustomWaitGroup_AddNegativePanic(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)
	wg.Done()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Negative counter should panic")
		}
	}()
	wg.Add(-2)
}

func TestCustomWaitGroup_ConcurrentSafety(t *testing.T) {
	wg := NewCustomWaitGroup()
	var mu sync.Mutex
	var calls int

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			calls++
			mu.Unlock()
		}()
	}

	wg.Wait()

	mu.Lock()
	if calls != 100 {
		t.Errorf("Expected 100 calls, got %d", calls)
	}
	mu.Unlock()
}

func TestCustomWaitGroup_AddAfterDone(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(1)
	go func() {
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		wg.Done()
	}()
	wg.Wait()
}
