package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() StringIntMap {
	return StringIntMap{
		data: make(map[string]int),
	}
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.data[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.data, key)
}

func (sim *StringIntMap) Copy() map[string]int {
	mapCopy := make(map[string]int)
	for key, value := range sim.data {
		mapCopy[key] = value
	}
	return mapCopy
}

func (sim *StringIntMap) Exists(key string) bool {
	_, exists := sim.data[key]
	return exists
}

func (sim *StringIntMap) Get(key string) (int, bool) {
	value, exists := sim.data[key]
	return value, exists
}

func main() {
	sim := NewStringIntMap()

	sim.Add("test", 1)
	fmt.Printf("1. Add element: %v\n", sim.data)

	sim.Remove("test")
	fmt.Printf("2. Delete element: %v\n", sim.data)

	sim.Add("a", 1)
	sim.Add("b", 2)
	simCopy := sim.Copy()
	fmt.Printf("3. Copy StrintIntMap. Original: %v\tCopy: %v\n", sim, simCopy)

	fmt.Printf("4. Check Exists. Original map: %v\tCheck 'a': %t\tCheck 'c': %t\n", sim.data, sim.Exists("a"), sim.Exists("c"))

	getA, existsA := sim.Get("a")
	getC, existsC := sim.Get("c")
	fmt.Printf("5. Check Get. Original map: %v\tGet 'a': %d %t\tGet 'c': %d %t\n", sim.data, getA, existsA, getC, existsC)

}
