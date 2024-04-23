package main

import (
	"fmt"
	"notion/task/arrays/pkg/arrays"
)

func main() {
	arr := []int{1, 1, 2, 3}
	fmt.Println(arrays.RemoveRepeating(arr))
}
