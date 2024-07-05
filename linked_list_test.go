package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := &linkedList[string, int]{head: nil}
	list.Push("1", 1)
	list.Push("2", 2)
	if val, _ := list.GetOrDefault("1", -1); val != 1 {
		t.Errorf("Incorrect value key=%s, got=%d", "1", val)
	}
	if val, _ := list.GetOrDefault("2", -1); val != 2 {
		t.Errorf("Incorrect value key=%s, got=%d", "2", val)
	}
}
