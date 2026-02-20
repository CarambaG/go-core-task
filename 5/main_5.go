package main

import "fmt"

func Intersection(a, b []int) (bool, []int) {
	set := make(map[int]int)
	for _, v := range a {
		set[v]++
	}

	result := make([]int, 0, len(b))
	for _, v := range b {
		if val, ok := set[v]; ok && val > 0 {
			result = append(result, v)
			set[v]--
		}
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	exists, result := Intersection(a, b)

	fmt.Printf("a: %v\nb: %v\nresult: %t, %v", a, b, exists, result)

}
