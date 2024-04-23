package main

import (
	"fmt"
	"notion/task/Algorithms/Searching/pkg/search"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(search.BinSearch(arr, 4))
}
