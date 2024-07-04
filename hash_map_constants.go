package main

const DefaultCapacity uint = 8
const DefaultLoadFactor float64 = 0.5

// HashMap defines a common interface for hash maps, also known as an associative array
type HashMap[K string, V any] interface {
	// Put inserts the give key-value pair into the hash map.
	// If the key already exists, the old value is overwritten.
	Put(key K, val V) 
	
	// Get returns the value associated with a specified key if it exists.
	Get(key K) (V, bool) 

	// Returns the value associated with the specified key if it exists,
	// else the default value if it doesn't exist.
	GetOrDefault(key K, defaultVal V) (V) 

	// Returns the number of key-value pairs in the hash map.
	Size() uint 

	// Returns a string representation of the hash map.
	String() string

	// Gives the number of equality comparisons used internally by the hash map.
	getNumCompares() uint 

	// Resets the counter for equality comparisons
	clearNumCompares() 
}