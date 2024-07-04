package main

import "fmt"

type HashMapChaining[K string, V any] struct {
	cap uint // number of buckets
	n uint // number of entries
	loadFactor float64 // desired upper bound for load factor
	buckets []linkedList[K, V] // each bucket holds a linked list
	_currentCompares uint
}

func NewHashMapChaining[K string, V any](cap uint, loadFactor float64) *HashMapChaining[K, V] {
	m := &HashMapChaining[K, V]{cap: cap, loadFactor: loadFactor}
	m.n = 0
	m.buckets = make([]linkedList[K, V], cap)
	for i := 0; i < int(m.cap); i++ {
		m.buckets[i] = linkedList[K, V]{}
	}
	return m
}

func (m *HashMapChaining[K, V]) Put(key K, val V) {
	if float64(m.n) / float64(m.cap) > m.loadFactor {
		m.resize(m.cap * 2)
	}
	m.buckets[m.hash(string(key))].Push(key, val)
	m.n += 1
}

func (m *HashMapChaining[K, V]) GetOrDefault(key K, defaultVal V) (V) {
	return m.buckets[m.hash(string(key))].GetOrDefault(key, defaultVal)
}

func (m *HashMapChaining[K, V]) Get(key K) (V, bool) {
	val, ok := m.buckets[m.hash(string(key))].Get(key)
	var dummyVal V
	if !ok {
		return dummyVal, false
	}
	return val, true
}

func (m *HashMapChaining[K, V]) GetNumCompares() uint {
	return m._currentCompares
}

func (m *HashMapChaining[K, V]) ClearNumCompares() {
	m._currentCompares = 0
}

func (m *HashMapChaining[K, V]) String() string {
	return fmt.Sprintf("<HashMapChaining n=%d, cap=%d, loadFactor=%f, _currentCompares=%d>", m.n, m.cap, m.loadFactor, m._currentCompares)
}

func (m *HashMapChaining[K, V]) resize(newCap uint) {
	newMap := NewHashMapChaining[K, V](newCap, m.loadFactor)
	for i := 0; i < int(m.cap); i++ {
		node := m.buckets[i].head
		for node != nil {
			newMap.Put(node.key, node.val.(V))
			node = node.next
		}
	}
	newMap.ClearNumCompares()
	newMap.copyTo(m)
}

func (m *HashMapChaining[K, V]) copyTo(other *HashMapChaining[K, V]) {
	other.cap = m.cap
	other.n = m.n
	other.loadFactor = m.loadFactor
	other.buckets = m.buckets
	other._currentCompares = m._currentCompares
}

func (m *HashMapChaining[K, V]) hash(key string) uint32 {
	return Hash(key) % uint32(m.cap)
}
