package arrays

import (
	"testing"
)

var testArray = func(n int) (arr []int) {
	for i := 0; i < n; i++ {
		arr = append(arr, i)
		arr = append(arr, i)
	}

	return
}

func BenchmarkRemoveRepeating(b *testing.B) {
	for i := 0; i < 100; i++ {
		RemoveRepeating(testArray(1000))
	}
}

var arr1 = func(n int) (res []int) {
	for i := 1; i < n; i += 2 {
		res = append(res, i)
	}

	return res
}

var arr2 = func(n int) (res []int) {
	for i := 1; i < n/2; i++ {
		res = append(res, i)
	}

	return
}

func BenchmarkMergeArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveRepeating(MergeArrays(arr1(10), arr2(10)))
	}
}

func BenchmarkUniteArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniteArrays(arr1(10), arr2(10))
	}
}
