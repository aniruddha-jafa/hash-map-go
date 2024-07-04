package main

type HashMapLinearProbing[K string, V any] struct {
	cap uint // number of slots
	n uint // number of entries
	loadFactor float64 // desired upper bound for load factor
	keys []K 
	values []V
}

func NewHashMapLinearProbing[K string, V any](cap uint, loadFactor float64) *HashMapLinearProbing[K, V] {
	m := &HashMapLinearProbing[K, V]{cap: cap, loadFactor: loadFactor}
	m.n = 0
	m.keys = make([]K, cap)
	m.values = make([]V, cap)
	return m
}

// Inserts or updates a key with the given value.
//
// If the key already exists update the value,
// else write to first available null slot in the probe sequence
//
// Double the size if load exceeds 50%
func (m *HashMapLinearProbing[K, V]) Put(key K, val V) {
	if float64(m.n) / float64(m.cap) > m.loadFactor {
		m = m.resize(m.cap * 2)
	}
	var keyNullVal K
	var i uint32 = 0
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		if m.keys[i] == key {
			m.values[i] = val
			return
		}
	}
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
	for i = m.hash(string(key)); m.keys[i] != keyNullVal; i = (i + 1) % uint32(m.cap) {
		if m.keys[i] == key {
			return m.values[i], true
		}
	}
	return valueNullVal, false
}

func (m *HashMapLinearProbing[K, V]) resize(newCap uint) *HashMapLinearProbing[K, V] {
	newMap := NewHashMapLinearProbing[K, V](newCap, m.loadFactor)
	var keyNullValue K
	for i := 0; i < int(m.cap); i++ {
		if m.keys[i] != keyNullValue {
			newMap.Put(m.keys[i], m.values[i])
		}
	}
	return newMap
}


func (m *HashMapLinearProbing[K, V]) hash(key string) uint32 {
	return Hash(key) % uint32(m.cap)
}
