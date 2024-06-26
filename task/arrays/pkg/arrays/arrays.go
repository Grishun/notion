package arrays

import (
	"errors"
	"math/rand"
	"slices"
	"sort"
)

// Max returns max element in an array
func Max(arr []int) int {

	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}

	return max
}

// Min returns min element in an array
func Min(arr []int) int {

	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}

	return min
}

// MidVal returns middle value of all numbers in the array
func MidVal(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / len(arr)
}

// IndexOfVal - func takes an array and a value,
// and returns index of the value in the array
// If array doesn't contain the value, func returns an error
func IndexOfVal(arr []int, value int) (res int, err error) {
	var flag bool
	for i, num := range arr {
		if num == value {
			res = i
			flag = true
			break
		}
	}
	if !flag {
		return 0, errors.New("no value")
	}

	return
}

func RemoveByIndex(arr []int, index int) []int {
	if index > len(arr) || index < 0 {
		return nil
	}

	newArr := arr[:0]

	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, arr[index+1:]...)

	return newArr
}

func RemoveByValue(arr []int, value int) []int {
	res := arr[:0]
	for _, num := range arr {
		if num != value {
			res = append(res, num)
		}
	}

	arr = res

	return arr
}

// Insert inserts value to the array, to the index
// Example
// Insert([1, 3, 4], 2, 1) -> [1, 2, 3, 4]
func Insert(arr []int, value, index int) []int {
	newArr := arr[:0]
	newArr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	arr = newArr
	return arr
}

// InsertInOrderedArr inserts value in ordered array and doesn't break the orderliness
func InsertInOrderedArr(arr []int, value int) []int {
	index := sort.Search(len(arr), func(i int) bool {
		return value < arr[i]
	})

	return Insert(arr, value, index)
}

// CountValues counts how many times the value appears in the array
//func CountValues(arr []int, value int) int {
//	counter := 0
//	for _, num := range arr {
//		if num == value {
//			counter++
//		}
//	}
//
//	return counter
//}

func RemoveRepeating(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	encountered := map[int]bool{}
	result := []int{}

	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

// UniteArrays unites elements from 2 arrays
// Example
// UniteArrays([1,2,3,4], [2,4,6,8]) -> [1,2,3,4,6,8]
func UniteArrays(arr1, arr2 []int) []int {
	united := make(map[int]struct{})

	for _, num1 := range arr1 {
		united[num1] = struct{}{}
	}

	for _, num2 := range arr2 {
		united[num2] = struct{}{}
	}

	res := make([]int, 0, len(arr1)+len(arr2))
	for key, _ := range united {
		res = append(res, key)
	}

	sort.Ints(res)

	return slices.Clip(res)
}

// MergeArrays merges 2 ordered arrays, without breaking ordering
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

	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)

	return
}

func RandArray(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(n))
	}

	return
}

func OrderedArray(n int) []int {
	ordered := make([]int, n)
	for i := 0; i < n; i++ {
		ordered[i] = i + 1
	}

	return ordered
}

func InversedArray(n int) []int {
	inversed := make([]int, n)
	for i := n - 1; i < 0; i-- {
		inversed = append(inversed, i)
	}

	return inversed
}
