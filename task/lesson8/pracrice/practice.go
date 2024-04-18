package pracrice

import (
	"notion/task/arrays/pkg/arrays"
	"sort"
)

// hard practice: This func searches for min & max values in an array
// and make all elements between them = 0
//
// Example:
// [1, -1, 1, 3, 4] -> [1, -1, 0, 0, 4]
func Obnulyator(arr []int) []int {
	IndexOfMin, err1 := arrays.IndexOfVal(arr, arrays.Min(arr))
	IndexOfMax, err2 := arrays.IndexOfVal(arr, arrays.Max(arr))

	if err1 != nil || err2 != nil {
		return []int{}
	}

	var (
		startFor int
		endFor   int
	)

	if IndexOfMin > IndexOfMax {
		startFor = IndexOfMax
		endFor = IndexOfMin
	} else {
		startFor = IndexOfMin
		endFor = IndexOfMax
	}

	for i := startFor + 1; i < endFor; i++ {
		arr[i] = 0
	}

	return arr
}

func InsertionV1(arr []int, value, index int) []int {

	return append(arr[:index], append([]int{value}, arr[index:]...)...)
}

func InsertionV2(arr []int, value, index int) []int {
	arr = append(arr, 0)

	copy(arr[index+1:], arr[index:len(arr)-1])

	arr[index] = value

	return arr
}

// This func counts how many times the value appears in the array
func CountValues(arr []int, value int) int {
	counter := 0
	for _, num := range arr {
		if num == value {
			counter++
		}
	}

	return counter
}

func RemoveByIndex(arr []int, index int) []int {
	if index < 0 || index > len(arr)-1 {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func RemoveByValue(arr []int, value int) (res []int) {
	for _, num := range arr {
		if num != value {
			res = append(res, num)
		}
	}

	return
}

func RemoveRepeating(arr []int) []int {
	for i, v := range arr {
		if CountValues(arr, v) > 1 {
			arr = RemoveByValue(arr, v)
			arr = InsertionV1(arr, v, i)
		}
	}

	return arr
}

//func UniteArrays()

func MergeArrays(arr1, arr2 []int) (res []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	return
}
