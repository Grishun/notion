package main

import (
	"fmt"
	"notion/task/arrays/pkg/arrays"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arrays.RemoveByValueV1(arr, 1))
	fmt.Println(arr)
}
