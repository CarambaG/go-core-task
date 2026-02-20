package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		wantBool bool
		want     []int
	}{
		{
			name:     "example case",
			a:        []int{65, 3, 58, 678, 64},
			b:        []int{64, 2, 3, 43},
			wantBool: true,
			want:     []int{64, 3},
		},
		{
			name:     "no intersection",
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			wantBool: false,
			want:     []int{},
		},
		{
			name:     "all intersect",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			wantBool: true,
			want:     []int{1, 2, 3},
		},
		{
			name:     "a empty",
			a:        []int{},
			b:        []int{1, 2},
			wantBool: false,
			want:     []int{},
		},
		{
			name:     "b empty",
			a:        []int{1, 2},
			b:        []int{},
			wantBool: false,
			want:     []int{},
		},
		{
			name:     "duplicates a",
			a:        []int{1, 2, 2, 3},
			b:        []int{2, 3, 3},
			wantBool: true,
			want:     []int{2, 3}, // уникальные
		},
		{
			name:     "duplicates b",
			a:        []int{1, 2, 3},
			b:        []int{2, 2, 2},
			wantBool: true,
			want:     []int{2},
		},
		{
			name:     "duplicates a and b",
			a:        []int{1, 2, 2, 3},
			b:        []int{2, 2, 2},
			wantBool: true,
			want:     []int{2, 2},
		},
		{
			name:     "zero values",
			a:        []int{0, 1, 0},
			b:        []int{0, 2},
			wantBool: true,
			want:     []int{0},
		},
		{
			name:     "negative numbers",
			a:        []int{-1, 0, 1},
			b:        []int{0, 1, 2},
			wantBool: true,
			want:     []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, got := Intersection(tt.a, tt.b)
			if gotBool != tt.wantBool {
				t.Errorf("Intersection(%v, %v) bool = %t, want %t", tt.a, tt.b, gotBool, tt.wantBool)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIntersection_Commutative(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5}

	_, interAB := Intersection(a, b)
	_, interBA := Intersection(b, a)

	if !reflect.DeepEqual(interAB, interBA) {
		t.Errorf("Intersection not commutative: AB=%v, BA=%v", interAB, interBA)
	}
}
