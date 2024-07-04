package main

func Hash(key string) uint32 {
	return fnv1a32(key)
}

const fnvOffsetBasis32 uint32 = 2166136261
const fnvPrime32 uint32 = 16777619

// fnv1a32 implements the 32-bit version of the FNV-1a non-cryptographic hash function
// 
// The hash value is initialized to the fnv_offset_basis for 32 bits
// 
// Then for each byte, the hash is calculated as -
// hash = (hash ^ data[i]) * fnv_prime
func fnv1a32(data string) uint32 {
  var hash uint32 = fnvOffsetBasis32
  for i := 0; i < len(data); i++ {
		hash ^= uint32(data[i])
		hash *= fnvPrime32
	}
	return hash
}

func djb32(data string) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(data); i++ {
		//  hash(i) = hash(i - 1) * 33 + str[i]
		hash = (hash << 5 + hash) + uint32(data[i])
	}
	return hash
}