package search

import (
	"notion/task/arrays/pkg/arrays"
	"sort"
	"testing"
)

func BenchmarkBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinSearch(arrays.OrderedArray(100), 25)
	}
}

func BenchmarkGoSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(arrays.OrderedArray(100), 25)
	}
}

func BenchmarkLineSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LineSearch(arrays.OrderedArray(100), 25)
	}
}
