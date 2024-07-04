package main

const DefaultCapacity uint = 8
const DefaultLoadFactor float64 = 0.5

type HashMap[K string, V any] interface {
	Put(key K, val V)
	Get(key K) (V, bool)
	GetOrDefault(key K, defaultVal V) (V)
	String() string
	getNumCompares() uint // Gives the number of equality comparisons used.
						  // Is reset whenever the hash table is resized.
	clearNumCompares() // Manually reset the counter for equality comparisons for profiling
}