package main

import (
	"fmt"
	"notion/task/Algorithms/sorting/pkg/sort"
)

func main() {
	arr := []int{3, 9, 6, 1, 4}
	fmt.Println(sort.QuickSort(arr))
}
