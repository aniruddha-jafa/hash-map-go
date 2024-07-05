package main

import "fmt"

type HashMapLinearProbing[K string, V any] struct {
	cap uint // number of slots
	size uint // number of entries
	loadFactor float64 // desired upper bound for load factor
	keys []K 
	values []V
	_currentCompares uint
}

func NewHashMapLinearProbing[K string, V any](cap uint, loadFactor float64) *HashMapLinearProbing[K, V] {
	m := &HashMapLinearProbing[K, V]{cap: cap, loadFactor: loadFactor}
	m.size = 0
	m._currentCompares = 0
	m.keys = make([]K, cap)
	m.values = make([]V, cap)
	return m
}

// Inserts or updates a key with the given value.
//
// If the key already exists update the value,
// else write to first available null slot in the probe sequence
//
// Double the size if load exceeds the load factor
func (m *HashMapLinearProbing[K, V]) Put(key K, val V) {
	if float64(m.size) / float64(m.cap) > m.loadFactor {
		m.resize(m.cap * 2)
	}
	var keyNullVal K
	var i uint32 = 0
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		m._currentCompares++
		if m.keys[i] == key {
			m.values[i] = val
			return
		}
	}
	m.keys[i] = key 
	m.values[i] = val
	m.size++
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
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		m._currentCompares++
		if m.keys[i] == key {
			return m.values[i], true
		}
	}
	return valueNullVal, false
}

func (m *HashMapLinearProbing[K, V]) getNumCompares() uint {
	return m._currentCompares
}

func (m *HashMapLinearProbing[K, V]) clearNumCompares() {
	m._currentCompares = 0
}

func (m *HashMapLinearProbing[K, V]) String() string {
	return fmt.Sprintf("<HashMapLinearProbing n=%d, cap=%d, loadFactor=%f, _currentCompares=%d, keys=%s>", m.size, m.cap, m.loadFactor, m._currentCompares, m.keys)
}

func (m *HashMapLinearProbing[K, V]) Size() uint {
	return m.size
}


func (m *HashMapLinearProbing[K, V]) resize(newCap uint) {
	newMap := NewHashMapLinearProbing[K, V](newCap, m.loadFactor)
	var keyNullValue K
	for i := 0; i < int(m.cap); i++ {
		if m.keys[i] != keyNullValue {
			newMap.Put(m.keys[i], m.values[i])
		}
	}
	newMap.copyTo(m)
}


// copyTo copies this map into another map `other`
func (m *HashMapLinearProbing[K, V]) copyTo(other *HashMapLinearProbing[K, V]) {
	other.cap = m.cap
	other.size = m.size
	other.loadFactor = m.loadFactor
	other.keys = m.keys
	other.values = m.values
	other._currentCompares = m._currentCompares
}

func (m *HashMapLinearProbing[K, V]) hash(key string) uint32 {
	return Hash(key) % uint32(m.cap)
}
