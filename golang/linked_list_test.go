package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := &linkedList[string, int]{head: nil}
	list.push("1", 1)
	list.push("2", 2)
	if val := list.getOrDefault("1", -1); val != 1 {
		t.Errorf("Incorrect value key=%s, got=%d", "1", val)
	}
	if val := list.getOrDefault("2", -1); val != 2 {
		t.Errorf("Incorrect value key=%s, got=%d", "2", val)
	}
}
