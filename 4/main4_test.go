package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "example case",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "no matches",
			slice1: []string{"a", "b", "c"},
			slice2: []string{"x", "y", "z"},
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "all matches",
			slice1: []string{"a", "b"},
			slice2: []string{"a", "b", "c"},
			want:   []string{},
		},
		{
			name:   "slice1 empty",
			slice1: []string{},
			slice2: []string{"a", "b"},
			want:   []string{},
		},
		{
			name:   "slice2 empty",
			slice1: []string{"a", "b"},
			slice2: []string{},
			want:   []string{"a", "b"},
		},
		{
			name:   "duplicates in slice2",
			slice1: []string{"apple", "banana"},
			slice2: []string{"banana", "banana"},
			want:   []string{"apple"},
		},
		{
			name:   "duplicates in slice1",
			slice1: []string{"apple", "apple", "banana"},
			slice2: []string{"banana"},
			want:   []string{"apple", "apple"},
		},
		{
			name:   "empty strings",
			slice1: []string{"", "a", ""},
			slice2: []string{"a"},
			want:   []string{"", ""},
		},
		{
			name:   "case sensitive",
			slice1: []string{"Apple", "apple", "Banana"},
			slice2: []string{"apple", "BANANA"},
			want:   []string{"Apple", "Banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference(%v, %v) = %v, want %v", tt.slice1, tt.slice2, got, tt.want)
			}
		})
	}
}

func TestDifference_Performance(t *testing.T) {
	slice1 := make([]string, 1000)
	slice2 := make([]string, 100)

	for i := range slice1 {
		slice1[i] = fmt.Sprintf("str%d", i)
	}
	for i := range slice2 {
		slice2[i] = fmt.Sprintf("str%d", i*10)
	}

	result := Difference(slice1, slice2)
	if len(result) != 900 { // 100 совпадений из 1000
		t.Errorf("Performance test: unexpected result length %d", len(result))
	}
}
