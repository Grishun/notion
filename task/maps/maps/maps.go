package maps

import (
	"math"
	"notion/task/arrays/pkg/arrays"
)

//func Set(arr []int) []int {
//	seen := make(map[int]struct{})
//
//	for i, v := range arr {
//
//		_, entry := seen[v]
//
//		if entry {
//			arr[i] = 0
//		} else {
//			seen[v] = struct{}{}
//		}
//	}
//
//	return arr
//
//}

// Unique возвращает новый массив, содержащий только уникальные значения из исходного массива
func Unique(arr []int) []int {

	uniqueArr := arr[:0]

	seen := make(map[int]struct{})

	for _, v := range arr {

		_, entry := seen[v]

		if !entry {
			seen[v] = struct{}{}
			uniqueArr = append(uniqueArr, v)
		}

	}

	arr = uniqueArr

	return arr
}

var (
	M1 = map[string]int{
		"Bob":     12,
		"Charlie": 17,
		"Eve":     1,
		"Grace":   3,
		"Hannah":  11,
		"Ivan":    2,
		"Jack":    13,
	}

	M2 = map[string]int{
		"Alice":   28,
		"Bob":     34,
		"Charlie": 22,
		"David":   29,
		"Frank":   31,
		"Grace":   27,
		"Ivan":    30,
		"Jack":    23,
	}
)

func MergeMaps(map1, map2 map[string]int) map[string]int {

	copyMap := func(m map[string]int) map[string]int {

		copied := make(map[string]int)

		for key, value := range m {
			copied[key] = value
		}

		return copied
	}

	res := copyMap(map1)

	for key, value := range map2 {

		_, entry := map1[key]

		if entry {
			res[key] = map1[key] + map2[key]
		} else {
			res[key] = value
		}
	}

	return res
}

func MostFrequentEl(arr []int) int {

	entrys := map[int]int{}

	for _, num := range arr {
		entrys[num]++
	}

	maxEntrys := 0
	mostfreq := 0

	for key, value := range entrys {
		if value > maxEntrys {
			maxEntrys = value
			mostfreq = key
		}
	}

	return mostfreq

}

func MiddleMarks(marks map[string][]int) map[string]int {

	midMarks := make(map[string]int)

	for key, marks := range marks {
		midMarks[key] = arrays.MidVal(marks)
	}

	return midMarks
}

// SmallerNumbersThanCurrent https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/
func SmallerNumbersThanCurrent(nums []int) (res []int) {

	smallerNums := map[int]int{}

	smaller := func(arr []int, val int) (count int) {

		for i := 0; i < len(arr); i++ {
			if arr[i] < val {
				count++
			}
		}
		return
	}

	for _, v := range nums {
		smallerNums[v] = smaller(nums, v)
	}

	for i := 0; i < len(nums); i++ {
		res = append(res, smallerNums[nums[i]])
	}

	return

}

// Task https://leetcode.com/problems/number-of-good-pairs/description/
func Task(nums []int) int {

	isGood := func(arr []int, i, j int) bool {
		return arr[i] == arr[j] && i < j
	}

	goodPairs := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if isGood(nums, i, j+i+1) {
				goodPairs++
			}
		}

	}

	return goodPairs
}

// Task2 https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/description/
func Task2(nums []int, k int) int {

	isCorrespondPair := func(arr []int, i, j int) bool {
		return int(math.Abs(float64(arr[i])-float64(arr[j]))) == k
	}

	counter := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if isCorrespondPair(nums, i, i+j+1) {
				counter++
			}
		}
	}

	return counter
}
