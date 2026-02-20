package main

import "fmt"

func Difference(slice1, slice2 []string) []string {
	set := make(map[string]struct{})
	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	result := make([]string, 0, len(slice1))
	for _, v := range slice1 {
		if _, ok := set[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := Difference(slice1, slice2)

	fmt.Printf("slice1: %v\nslice2: %v\nresult: %v", slice1, slice2, result)
}
