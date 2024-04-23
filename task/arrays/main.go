package main

import (
	"fmt"
	"notion/task/arrays/pkg/arrays"
)

func main() {
	arr1 := []int{1, 2, 3, -1, 5, 4}
	arr2 := []int{5, 1, 2, 3, 4, -1}
	fmt.Println(arrays.MergeArrays(arr1, arr2))
}
