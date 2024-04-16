package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func randArray(n int) []int {
	randArray := make([]int, n)
	for i := 0; i < n; i++ {
		randArray = append(randArray, rand.Intn(100000))
	}

	return randArray
}

func orderedArray(n int) []int {
	ordered := make([]int, n)
	for i := 0; i < n; i++ {
		ordered = append(ordered, i)
	}

	return ordered
}

func inversedArray(n int) []int {
	inversed := make([]int, n)
	for i := n - 1; i < 0; i-- {
		inversed = append(inversed, i)
	}

	return inversed
}

func BenchmarkSelectionSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(randArray(1000))
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(randArray(1000))
	}
}

func BenchmarkGoSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SliceStable(randArray(1000), func(i, j int) bool {
			return i < j
		})
	}
}

func BenchmarkQuickSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randArray(1000))
	}
}

func BenchmarkSelectionSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(orderedArray(1000))
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(orderedArray(1000))
	}
}

func BenchmarkGoSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SliceStable(orderedArray(1000), func(i, j int) bool {
			return i < j
		})
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(orderedArray(1000))
	}
}

func BenchmarkSelectionSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(inversedArray(1000))
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(inversedArray(1000))
	}
}

func BenchmarkGoSor3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SliceStable(inversedArray(1000), func(i, j int) bool {
			return i < j
		})
	}
}

func BenchmarkQuickSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(inversedArray(1000))
	}
}
