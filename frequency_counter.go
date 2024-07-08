package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

// Finds the most frequent word whose length is
// greater than or equal to the threshold,
// and prints the result to  standard output
func main() {
	minLen := flag.Uint("n", 8, "minimum length of words to consider")
	mapImpl := flag.String("impl", "", "hash map implementation to use - chain, linear")
	initialCap := flag.Uint("cap", DefaultCapacity, "Initial capacity of hash map")
	loadFactor := flag.Float64("load", math.NaN(), "Load factor of hash map")

	flag.Parse()

	var hashMap HashMap[string, uint]
	switch *mapImpl {
		case "chain":
			if math.IsNaN(*loadFactor) {
				*loadFactor = float64(DefaultLoadFactorChaining)
			}
			hashMap = NewHashMapChaining[string, uint](*initialCap, *loadFactor) 
		case "linear":
			if math.IsNaN(*loadFactor) {
				*loadFactor = float64(DefaultLoadFactorProbing)
			}
			hashMap = NewHashMapLinearProbing[string, uint](*initialCap, *loadFactor)
		default:
			panic("Unknown impl: " + *mapImpl)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := 0
	var totalCompares uint

	fmt.Println("s_no,key,num_equality_compares,cumulative_avg")

	for scanner.Scan() {
		w := scanner.Text()
		if uint(len(w)) < *minLen {
			continue
		}
		n++
		var compares uint
		prevCount := hashMap.GetOrDefault(w, 0)
		hashMap.Put(w, prevCount + 1)
		
		compares = hashMap.getNumCompares()
		totalCompares += compares
		fmt.Printf("%d,%s,%d,%.2f\n", n, w, compares, float64(totalCompares) / float64(n))
		
		hashMap.clearNumCompares()
		}
}