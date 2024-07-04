package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Finds the most frequent word whose length is
// greater than or equal to the threshold,
// and prints the result to  standard output
func main() {
	minLen := flag.Uint("n", 8, "Minimum length of words to consider")
	mapImpl := flag.String("impl", "", "hash map implementation to use")
	flag.Parse()

	var hashMap HashMap[string, uint]
	switch *mapImpl {
		case "chain":
			hashMap = NewHashMapChaining[string, uint](DefaultCapacity, DefaultLoadFactor) 
		case "linear":
			hashMap = NewHashMapLinearProbing[string, uint](DefaultCapacity, DefaultLoadFactor)
		default:
			panic("Unknown impl: " + *mapImpl)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := 0
	var totalCompares uint

	fmt.Println("key,num_equality_compares,cumulative_avg")

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
		fmt.Printf("%s,%d,%.2f\n", w, compares, float64(totalCompares) / float64(n))
		hashMap.clearNumCompares()
		}
}