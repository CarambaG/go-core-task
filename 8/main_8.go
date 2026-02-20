package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	mu    sync.Mutex
	cond  *sync.Cond
	count int
}

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{}
	wg.cond = sync.NewCond(&wg.mu)
	return wg
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	wg.count += delta
	wg.mu.Unlock()

	if wg.count < 0 {
		panic("negative counter")
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.mu.Lock()
	wg.count--
	wasZero := wg.count == 0
	wg.mu.Unlock()

	if wasZero {
		wg.cond.Broadcast()
	}

	if wg.count < 0 {
		panic("negative counter")
	}
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	for wg.count > 0 {
		wg.cond.Wait()
	}
	wg.mu.Unlock()
}

func main() {
	wg := NewCustomWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	fmt.Println("Waiting for all goroutines...")
	wg.Wait()
	fmt.Println("All goroutines completed!")
}
