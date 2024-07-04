package main

import (
	"fmt"
	"strings"
)

// Single linked list
type linkedList[K comparable, V any] struct {
	head *node[K, V]
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
func (l *linkedList[K, V]) Push(key K, val V) {
	// check for existing key
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.key == key {
			curr.val = val
			return
		}
	}
	l.head = &node[K, V]{key: key, val: val, next: l.head}
}

func (l *linkedList[K, V]) Get(key K) (V, bool) {
	for curr := l.head; curr != nil; curr = curr.next {
		if (curr.key == key) {
			return curr.val.(V), true 
		}
	}
	var dummyVal V 
	return dummyVal, false
}

func (l *linkedList[K, V]) GetOrDefault(key K, defaultVal V) (V) {
	val, ok := l.Get(key)
	if !ok {
		return defaultVal
	}
	return val
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