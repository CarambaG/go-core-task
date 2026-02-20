package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	input := make(chan uint8)
	go func() {
		defer close(input)
		input <- 2 // 2^3 = 8
		input <- 3 // 3^3 = 27
		input <- 5 // 5^3 = 125
	}()

	result := CubePipeline(input)

	expected := []float64{8.0, 27.0, 125.0}
	actual := make([]float64, 0, len(expected))

	for cube := range result {
		actual = append(actual, cube)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCubePipeline_Empty(t *testing.T) {
	input := make(chan uint8)
	close(input)

	result := CubePipeline(input)

	select {
	case _, ok := <-result:
		if ok {
			t.Error("Empty input should produce empty output")
		}
	default:
		// OK
	}
}

func TestCubePipeline_Zero(t *testing.T) {
	input := make(chan uint8)
	go func() {
		defer close(input)
		input <- 0 // 0^3 = 0
	}()

	result := CubePipeline(input)
	cube := <-result

	if cube != 0.0 {
		t.Errorf("Expected 0.0, got %f", cube)
	}
}

func TestCubePipeline_MaxUint8(t *testing.T) {
	input := make(chan uint8)
	go func() {
		defer close(input)
		input <- 255 // 255^3 = 16_581_375
	}()

	result := CubePipeline(input)
	cube := <-result

	expected := 255.0 * 255.0 * 255.0
	if cube != expected {
		t.Errorf("Expected %.0f, got %f", expected, cube)
	}
}

func TestUint8Generator(t *testing.T) {
	ch := uint8Generator(3)

	expected := []uint8{0, 1, 2}
	actual := make([]uint8, 0, len(expected))

	for num := range ch {
		actual = append(actual, num)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFullPipeline(t *testing.T) {
	input := uint8Generator(5)
	result := CubePipeline(input)

	expected := []float64{0, 1, 8, 27, 64}
	actual := make([]float64, 0, len(expected))

	for cube := range result {
		actual = append(actual, cube)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Full pipeline expected %v, got %v", expected, actual)
	}
}
