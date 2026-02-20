package main

import (
	"reflect"
	"testing"
)

func TestStringIntMap_Add(t *testing.T) {
	m := NewStringIntMap()

	tests := []struct {
		key   string
		value int
	}{
		{"key1", 100},
		{"key2", 200},
		{"key1", 300}, // перезапись
	}

	for _, tt := range tests {
		m.Add(tt.key, tt.value)
	}

	if val, ok := m.Get("key1"); !(ok && val == 300) {
		t.Errorf("Add failed: key1 = %d (ok=%t), want 300", val, ok)
	}
	if val, ok := m.Get("key2"); !(ok && val == 200) {
		t.Errorf("Add failed: key2 = %d (ok=%t), want 200", val, ok)
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("test", 42)

	m.Remove("test")
	if _, ok := m.Get("test"); ok {
		t.Error("Remove failed: key 'test' still exists")
	}

	m.Remove("nonexistent")
	if len(m.data) != 0 {
		t.Error("Remove nonexistent key changed length")
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 42)

	copyMap := m.Copy()

	// Equal
	if !reflect.DeepEqual(copyMap, m.data) {
		t.Error("Copy content mismatch")
	}

	// Not equal
	copyMap["modified"] = 999
	if _, ok := m.Get("modified"); ok {
		t.Error("Copy modified original!")
	}

	m.Add("new", 123)
	if _, ok := copyMap["new"]; ok {
		t.Error("Original modified copy!")
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("exists", 1)

	tests := []struct {
		key  string
		want bool
	}{
		{"exists", true},
		{"missing", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := m.Exists(tt.key); got != tt.want {
			t.Errorf("Exists(%q) = %t, want %t", tt.key, got, tt.want)
		}
	}
}

func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 42)

	val, ok := m.Get("key")
	if !ok || val != 42 {
		t.Errorf("Get('key') = (%d, %t), want (42, true)", val, ok)
	}

	val, ok = m.Get("missing")
	if ok || val != 0 {
		t.Errorf("Get('missing') = (%d, %t), want (0, false)", val, ok)
	}
}

func TestStringIntMap_EdgeCases(t *testing.T) {
	// Empty map
	m := NewStringIntMap()
	if m.Exists("any") || len(m.data) != 0 {
		t.Error("New map should be empty")
	}

	// Empty key
	m.Add("", 123)
	if !m.Exists("") {
		t.Error("Empty key should work")
	}

	m.Remove("")
	if m.Exists("") {
		t.Error("Remove empty key failed")
	}
}

func TestStringIntMap_Len(t *testing.T) {
	m := NewStringIntMap()
	if len(m.data) != 0 {
		t.Error("New map len != 0")
	}

	m.Add("a", 1)
	if len(m.data) != 1 {
		t.Error("Len after 1 Add != 1")
	}

	m.Add("b", 2)
	m.Remove("a")
	if len(m.data) != 1 {
		t.Error("Len after Add+Remove != 1")
	}
}
