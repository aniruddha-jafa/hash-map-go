package main

const DefaultCapacity uint = 16

type HashMapChaining[K string, V any] struct {
	cap uint // number of buckets
	n uint // number of entries
	buckets []linkedList[K, V]
}

func NewHashMapChaining[K string, V any](cap uint) *HashMapChaining[K, V] {
	m := &HashMapChaining[K, V]{cap: cap}
	m.n = 0
	m.buckets = make([]linkedList[K, V], cap)
	for i := 0; i < int(m.cap); i++ {
		m.buckets[i] = linkedList[K, V]{}
	}
	return m
}

func (m *HashMapChaining[K, V]) Put(key K, val V) {
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

func (m *HashMapChaining[K, V]) hash(key string) uint32 {
	return Hash(key) % uint32(m.cap)
}
