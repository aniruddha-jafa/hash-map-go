package main

import (
	"fmt"
	"strings"
)

type linkedList[K comparable, V any] struct {
	head *node[K, V]
}

type node[K comparable, V any] struct {
	key K
	val any
	next *node[K, V]
}

func (l *linkedList[K, V]) push(key K, val V) {
	// check for existing key
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.key == key {
			curr.val = val
			return
		}
	}
	// add at head
	l.head = &node[K, V]{key: key, val: val, next: l.head}
}

func (l *linkedList[K, V]) String() string {
	out := strings.Builder{}
	out.WriteString("[")
	for curr := l.head; curr != nil; curr = curr.next {
		out.WriteString(curr.string())
		if curr.next != nil {
			out.WriteString(", ")
		}
	}
	out.WriteString("]")
	return out.String()
}

func (n *node[K, V]) string() string {
	return fmt.Sprintf("{key=%v, value=%v}", n.key, n.val)
}

func (l *linkedList[K, V]) getOrDefault(key K, defaultVal V) (V) {
	for curr := l.head; curr != nil; curr = curr.next {
		if (curr.key == key) {
			return curr.val.(V)
		}
	}
	return defaultVal
}