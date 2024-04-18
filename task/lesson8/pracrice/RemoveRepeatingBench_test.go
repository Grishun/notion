package pracrice

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

func BenchmarkRemoveRepeatingV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveRepeating(testArray(100))
	}
}

func BenchmarkRemoveRepeatingV2(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
