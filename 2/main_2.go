package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(originalSlice []int) []int {
	newSlice := make([]int, 0, len(originalSlice))
	for _, value := range originalSlice {
		if value%2 == 0 {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func addElements(originalSlice []int, value int) []int {
	return append(originalSlice, value)
}

func copySlice(originalSlice []int) []int {
	newSlice := make([]int, len(originalSlice))
	copy(newSlice, originalSlice)
	return newSlice
}

func removeElement(originalSlice []int, index int) ([]int, error) {
	if index >= len(originalSlice) {
		return originalSlice, errors.New("index out of range")
	}
	return append(originalSlice[:index], originalSlice[index+1:]...), nil
}

func main() {
	originalSlice := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Printf("1. originalSlice:%v\n", originalSlice)

	originalSlice = sliceExample(originalSlice)
	fmt.Printf("2. originalSlice after sliceExample: %v\n", originalSlice)

	originalSlice = addElements(originalSlice, 100)
	fmt.Printf("3. originalSlice after addElements: %v\n", originalSlice)

	copiedSlice := copySlice(originalSlice)
	copiedSlice[0] = -1
	fmt.Printf("4. originalSlice: %v\t copiedSlice: %v\n", originalSlice, copiedSlice)

	element := originalSlice[len(originalSlice)/2]
	var err error
	originalSlice, err = removeElement(originalSlice, len(originalSlice)/2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("5. originalSlice after removeElement (remove %d): %v\n", element, originalSlice)
	}
}
