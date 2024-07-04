package main

import (
	"hash/fnv"
	"testing"
)

// Compare output with 32-bit FNV-1a hash function from go standard library
func fnv1aStandardLibHash32(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}


func TestFnvHash(t *testing.T) {
	var tests = []string{
		"indomitable",
		"aso97bciaas)*ACSv98",
		"a",
		"",
		"запустить",
		"радость",
		"失われた",
		"空気",
		"عظيم",
		"الجلوس",
		"It was the best of times it was the worst of times.",
		"12398103891",
	}
	for _, data := range tests {
		 testname := data
		 t.Run(testname, func(t *testing.T) {
            want := fnv1aStandardLibHash32(data)
			got := fnv1a32(data)
            if want != got {
                t.Errorf("got=%d, want=%d", got, want)
            }
        })
	}
}