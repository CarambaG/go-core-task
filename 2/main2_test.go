package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "even numbers only",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{2, 4, 6},
		},
		{
			name:  "no even numbers",
			input: []int{1, 3, 5, 7},
			want:  []int{},
		},
		{
			name:  "all even numbers",
			input: []int{2, 4, 6, 8},
			want:  []int{2, 4, 6, 8},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "zero included",
			input: []int{0, 1, 2, 3},
			want:  []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.input, got, tt.want)
			}

			if len(tt.input) != len(tt.input) {
				t.Error("original slice was modified!")
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		value int
		want  []int
	}{
		{"empty", []int{}, 42, []int{42}},
		{"non-empty", []int{1, 2, 3}, 99, []int{1, 2, 3, 99}},
		{"negative", []int{10}, -5, []int{10, -5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.input, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements(%v, %d) = %v, want %v", tt.input, tt.value, got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"empty", []int{}},
		{"single", []int{42}},
		{"multiple", []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copied := copySlice(tt.input)

			if !reflect.DeepEqual(copied, tt.input) {
				t.Errorf("copySlice(%v) content mismatch: got %v, want %v", tt.input, copied, tt.input)
			}

			if len(tt.input) > 0 {
				copied[0] = -999
				if tt.input[0] == -999 {
					t.Error("copySlice modified original slice!")
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		index     int
		want      []int
		wantPanic bool
	}{
		{
			name:      "middle",
			input:     []int{1, 2, 3, 4, 5},
			index:     2,
			want:      []int{1, 2, 4, 5},
			wantPanic: false,
		},
		{
			name:      "first",
			input:     []int{1, 2, 3},
			index:     0,
			want:      []int{2, 3},
			wantPanic: false,
		},
		{
			name:      "last",
			input:     []int{1, 2, 3},
			index:     2,
			want:      []int{1, 2},
			wantPanic: false,
		},
		{
			name:      "empty",
			input:     []int{},
			index:     0,
			want:      []int{},
			wantPanic: true,
		},
		{
			name:      "out of range",
			input:     []int{1, 2, 3},
			index:     3,
			want:      []int{1, 2, 3},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.input, tt.index)
			if (err != nil) && !tt.wantPanic {
				t.Errorf("removeElement(%v, %v) error = %v, wantErr %v", tt.input, tt.index, err, tt.wantPanic)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.input, tt.index, got, tt.want)
			}
		})
	}
}

func TestMainLogic(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6}

	originalSlice = sliceExample(originalSlice)
	if len(originalSlice) != 3 { // 2,4,6
		t.Errorf("sliceExample должен вернуть 3 чётных, получил %d", len(originalSlice))
	}

	originalSlice = addElements(originalSlice, 100)
	if len(originalSlice) != 4 || originalSlice[3] != 100 {
		t.Errorf("addElements не добавил 100: %v", originalSlice)
	}

	copiedSlice := copySlice(originalSlice)
	copiedSlice[0] = -1
	if originalSlice[0] == -1 {
		t.Error("copySlice изменил оригинал!")
	}

	midIndex := len(originalSlice) / 2
	originalSlice, _ = removeElement(originalSlice, midIndex)
	if len(originalSlice) != 3 {
		t.Errorf("removeElement не удалил элемент: %v", originalSlice)
	}
}
