package search

import (
	"notion/task/arrays/pkg/arrays"
	"slices"
	"sort"
	"testing"
)

func BenchmarkBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinSearch(arrays.RandArray(10000), 1)
	}
}

func BenchmarkGoSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(arrays.RandArray(10000), 1)
	}
}

func BenchmarkSliceGoSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.BinarySearch(arrays.RandArray(10000), 1)
	}
}

func BenchmarkLineSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LineSearch(arrays.RandArray(10000), 1)
	}
}
