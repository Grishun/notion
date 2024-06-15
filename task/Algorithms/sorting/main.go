package main

import (
	"fmt"
	"notion/task/Algorithms/sorting/pkg/sort"
)

func main() {
	arr := []int{1, 3, 4, 5, 2, 6}
	fmt.Println(sort.QuickSort(arr))
}
