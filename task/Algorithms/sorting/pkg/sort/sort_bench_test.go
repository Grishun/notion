package sort

import (
	"math/rand"
	"notion/task/arrays/pkg/arrays"
	"slices"
	"sort"
	"testing"
)

var randArray = func(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, rand.Int())
	}

	return
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

func BenchmarkGoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SliceStable(randArray(1000), func(i, j int) bool {
			return i < j
		})
	}
}

func BenchmarkSliceGoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.Sort(arrays.RandArray(1000))
	}
}

func BenchmarkQuickSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randArray(1000))
	}
}
func BenchmarkQuickSortV21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(arrays.RandArray(1000))
	}
}

//func BenchmarkSelectionSort2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SelectionSort(arrays.OrderedArray(1000))
//	}
//}
//
//func BenchmarkBubbleSort2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SelectionSort(arrays.OrderedArray(1000))
//	}
//}
//
//func BenchmarkGoSort2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		sort.SliceStable(arrays.OrderedArray(1000), func(i, j int) bool {
//			return i < j
//		})
//	}
//}
//
//func BenchmarkQuickSort2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		QuickSort(arrays.OrderedArray(1000))
//	}
//}
//func BenchmarkQuickSortV22(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		QuickSort(arrays.OrderedArray(1000))
//	}
//}
//
//func BenchmarkSelectionSort3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SelectionSort(arrays.InversedArray(1000))
//	}
//}
//
//func BenchmarkBubbleSort3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SelectionSort(arrays.InversedArray(1000))
//	}
//}
//
//func BenchmarkGoSor3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		sort.SliceStable(arrays.InversedArray(1000), func(i, j int) bool {
//			return i < j
//		})
//	}
//}
//
//func BenchmarkQuickSort3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		QuickSort(arrays.InversedArray(1000))
//	}
//}
//
//func BenchmarkQuickSortV23(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		QuickSort(arrays.InversedArray(1000))
//	}
//}
