package main

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	h := NewHashMapChaining[string, int](DefaultCapacity, DefaultLoadFactor)
	words := []string{
		"ant", "ant", "ant",
		"bat",
		"cat", "cat",
	}
	for _, w := range words {
		h.Put(w, h.GetOrDefault(w, 0) + 1)
	}
	if val := h.GetOrDefault("ant", 0); val != 3 {
		t.Errorf("got=%d, want=%d", val, 3)
	}
	if val := h.GetOrDefault("bat", 0); val != 1 {
		t.Errorf("got=%d, want=%d", val, 1)
	}
	if val := h.GetOrDefault("cat", 0); val != 2 {
		t.Errorf("got=%d, want=%d", val, 2)
	}
}