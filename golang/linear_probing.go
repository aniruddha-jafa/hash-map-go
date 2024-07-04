package main

import "fmt"

type HashMapLinearProbing[K string, V any] struct {
	cap uint // number of slots
	n uint // number of entries
	keys []K 
	values []V
}

func NewHashMapLinearProbing[K string, V any](cap uint) *HashMapLinearProbing[K, V] {
	m := &HashMapLinearProbing[K, V]{cap: cap}
	m.n = 0
	m.keys = make([]K, cap)
	m.values = make([]V, cap)
	return m
}

func (m *HashMapLinearProbing[K, V]) Put(key K, val V) {
	// resize when load exceeds 50%
	if m.n >= m.cap / 2 {
		m.resize(m.cap * 2)
		fmt.Printf("resized to: %d", m.cap)
		fmt.Println(m.keys, m.values)
	}
	var keyNullVal K
	var i uint32 = 0
	// if key already exists, update the value
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		if m.keys[i] == key {
			m.values[i] = val
			return
		}
	}
	// else write to first null slot
	m.keys[i] = key 
	m.values[i] = val
	m.n++
}

func (m *HashMapLinearProbing[K, V]) GetOrDefault(key K, defaultVal V) (V) {
	val, ok := m.Get(key)
	if !ok {
		return defaultVal
	}
	return val
}

func (m *HashMapLinearProbing[K, V]) Get(key K) (V, bool) {
	var keyNullVal K
	var valueNullVal V 
	var i uint32 = 0
	// if key already exists, update the value
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		if m.keys[i] == key {
			return m.values[i], true
		}
	}
	return valueNullVal, false
}

func (m *HashMapLinearProbing[K, V]) resize(newCap uint) {
	newMap := NewHashMapLinearProbing[K, V](newCap)
	var keyNullValue K
	for i := 0; i < int(m.cap); i++ {
		if m.keys[i] != keyNullValue {
			newMap.Put(m.keys[i], m.values[i])
		}
	}
	m = newMap
}


func (m *HashMapLinearProbing[K, V]) hash(key string) uint32 {
	return Hash(key) % uint32(m.cap)
}
