package main

import (
	"bufio"
	"os"
	"testing"
)

func TestCountsChaining(t *testing.T) {
	inbuiltMap := make(map[string]int, 8)
	customMap := NewHashMapChaining[string, uint](DefaultCapacity, 8)

	f, err := os.Open("./tale.txt")
	if err != nil {
		panic("Unable to open file")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		// add to inbuilt map
		_, ok := inbuiltMap[w]
		if !ok {
			inbuiltMap[w] = 0
		}
		inbuiltMap[w]++

		// add to custom map
		customMap.Put(w, customMap.GetOrDefault(w, 0) + 1)
		customMap.clearNumCompares()
	}
	if customMap.n != uint(len(inbuiltMap)) {
		t.Errorf("Expected length=%d, got=%d", len(inbuiltMap), customMap.n)
	}
	for w, expectedCount := range inbuiltMap {
		actualCount, ok := customMap.Get(w)
		if !ok {
			t.Errorf("%s not found in map", w)
			continue
		}
		if expectedCount != int(actualCount) {
			t.Errorf("for word=%s, expected=%d, got=%d", w, expectedCount, actualCount)
		}	
	}
}