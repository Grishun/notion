package arrays

import (
	"errors"
	"math/rand"
	"notion/task/Algorithms/Searching/pkg/search"
	"sort"
)

var matrix = [2][3]int{
	{1, 2, 3},
	{4, 5, 6},
}

// This func returns max element in an array
func Max(arr []int) int {

	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}

	return max
}

// This func returns min element in an array
func Min(arr []int) int {

	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}

	return min
}

// This func returnes middle value of all numbers in the array
func MidVal(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / len(arr)
}

// func takes an array and a value,
// and returns index of the value in the array
// If array doesn't contain the value, func returnes an error
func IndexOfVal(arr []int, value int) (res int, err error) {
	var flag bool
	for i, num := range arr {
		if num == value {
			res = i
			flag = true
			break
		}
	}
	if flag != true {
		return 0, errors.New("no value")
	}

	return
}

func PreMaxV2(arr []int) int {
	var arr2 []int
	for _, v := range arr {
		if v != Max(arr) {
			arr2 = append(arr2, v)
		}
	}

	return Max(arr2)
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

func Insert(arr []int, value, index int) []int {

	return append(arr[:index], append([]int{value}, arr[index:]...)...)
}

func InsertInOrderedArr(arr []int, value int) []int {
	index := sort.Search(len(arr), func(i int) bool {
		return value < arr[i]
	})

	return Insert(arr, value, index)
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

func RemoveRepeating(arr []int) []int {
	res := make([]int, 0)
	res = append(res, arr...)

	for i, v := range res {
		if CountValues(res, v) > 1 {
			res = RemoveByIndex(res, i)
		}
	}

	return res
}

func RemoveRepeatingV2(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	encountered := map[int]bool{}
	result := []int{}

	for _, v := range arr {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

// this func unites ellements from2 arrays
// Example
// UniteArrays([1,2,3,4], [2,4,6,8]) -> [1,2,3,4,6,8]

func UniteArrays(arr1, arr2 []int) (res []int) {
	res = append(res, arr1...)

	for _, v := range arr2 {
		_, err := search.BinSearch(arr1, v)

		if err != nil {
			res = InsertInOrderedArr(res, v)
		}
	}

	return
}

// this func merges 2 ordered arrays, without breaking ordering
// Example
// MergeArrays([2,4,6], [1,3,5]) -> [1,2,3,4,5]
// MergeArrays([1,2,3,4], [2,4,6,8]) -. [1,2,2,3,4,4,6,8]

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

func RandArray(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(n))
	}

	return
}
