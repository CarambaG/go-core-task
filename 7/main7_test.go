package main

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestMergeChannels_Basic(t *testing.T) {
	ctx := context.Background()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hello"
		ch1 <- "world"
		close(ch1)
	}()
	go func() {
		ch2 <- "foo"
		ch2 <- "bar"
		close(ch2)
	}()

	merged := MergeChannels[string](ctx, ch1, ch2)

	received := make(map[string]bool)
	for msg := range merged {
		received[msg] = true
	}

	expected := map[string]bool{
		"hello": true,
		"world": true,
		"foo":   true,
		"bar":   true,
	}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("Expected %v, got %v", expected, received)
	}
}

func TestMergeChannels_EmptyChannels(t *testing.T) {
	ctx := context.Background()

	// 2 empty channels (already close)
	ch1 := make(chan string)
	ch2 := make(chan string)
	close(ch1)
	close(ch2)

	merged := MergeChannels(ctx, ch1, ch2)

	select {
	case _, ok := <-merged:
		if ok {
			t.Error("Merged channel from empty inputs should be closed immediately")
		}
	default:
		// OK
	}
}

func TestMergeChannels_NoChannels(t *testing.T) {
	ctx := context.Background()
	merged := MergeChannels[string](ctx) // 0 channels

	select {
	case _, ok := <-merged:
		if ok {
			t.Error("Empty merge should close immediately")
		}
	default:
		// OK
	}
}

func TestMergeChannels_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "delayed"
		close(ch1)
	}()

	merged := MergeChannels(ctx, ch1)

	cancel()

	select {
	case msg := <-merged:
		t.Logf("Received: %s", msg)
	case <-time.After(100 * time.Millisecond):
		t.Error("MergeChannels didn't respect context cancellation")
	}
}

func TestMergeChannels_StringGenerator(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Используем ваш stringGenerator
	ch1 := stringGenerator(ctx, "A", 3)
	ch2 := stringGenerator(ctx, "B", 2)

	merged := MergeChannels(ctx, ch1, ch2)

	messages := make(map[string]bool)
	count := 0

	for msg := range merged {
		messages[msg] = true
		count++
		if count > 10 {
			break
		}
	}

	// Checking for expected messages
	expected := []string{"A-0", "A-1", "A-2", "B-0", "B-1"}
	for _, exp := range expected {
		if !messages[exp] {
			t.Errorf("Missing expected message: %s", exp)
		}
	}

	t.Logf("Received %d unique messages", len(messages))
}

func TestMergeChannels_AllChannelsClose(t *testing.T) {
	ctx := context.Background()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "one"
		close(ch1)
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "two"
		close(ch2)
	}()

	merged := MergeChannels(ctx, ch1, ch2)

	msgs := make([]string, 0, 2)
	for msg := range merged {
		msgs = append(msgs, msg)
	}

	if len(msgs) != 2 {
		t.Errorf("Expected 2 messages, got %d: %v", len(msgs), msgs)
	}

	// Checking that merged has closed
	select {
	case _, ok := <-merged:
		if ok {
			t.Error("Output channel not closed after all inputs closed")
		}
	default:
	}
}
