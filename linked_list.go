package main

import (
	"fmt"
	"strings"
)

// Single linked list
type linkedList[K comparable, V any] struct {
	head *node[K, V]
	size uint // number of items
}

// Node for singly linked list
type node[K comparable, V any] struct {
	key K
	val any
	next *node[K, V]
}

// Push inserts or updates a key with the given value.
//
// If the key already exists update its value,
// else add a new node at the head of the list.
//
// Returns - 
// a bool indicating if the key already existed in the hash map
// the number of equality comparisons used
func (l *linkedList[K, V]) Push(key K, val V) (bool, uint) {
	var compares uint = 0
	// check for existing key
	for curr := l.head; curr != nil; curr = curr.next {
		compares++
		if curr.key == key {
			curr.val = val
			return true, compares
		}
	}
	l.head = &node[K, V]{key: key, val: val, next: l.head}
	l.size++
	return false, compares
}

func (l *linkedList[K, V]) Get(key K) (V, bool, uint) {
	var compares uint = 0
	for curr := l.head; curr != nil; curr = curr.next {
		compares++
		if (curr.key == key) {
			return curr.val.(V), true, compares
		}
	}
	var dummyVal V 
	return dummyVal, false, compares
}

func (l *linkedList[K, V]) GetOrDefault(key K, defaultVal V) (V, uint) {
	val, ok, compares := l.Get(key)
	if !ok {
		return defaultVal, compares
	}
	return val, compares
}

func (l *linkedList[K, V]) String() string {
	out := strings.Builder{}
	out.WriteString("[")
	for curr := l.head; curr != nil; curr = curr.next {
		out.WriteString(curr.String())
		if curr.next != nil {
			out.WriteString(", ")
		}
	}
	out.WriteString("]")
	return out.String()
}

func (n *node[K, V]) String() string {
	return fmt.Sprintf("{key=%v, value=%v}", n.key, n.val)
}