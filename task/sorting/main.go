package main

import (
	"fmt"
	"notion/task/sorting/pkg/sort"
)

func main() {
	fmt.Println(sort.QuickSort([]int{0, 5, 1, 3, 4, 2}))
}
