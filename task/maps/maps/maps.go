package maps

import "notion/task/arrays/pkg/arrays"

var test = map[int]bool{}

func Set(arr []int) []int {
	seen := make(map[int]struct{})

	for i, v := range arr {

		_, entry := seen[v]

		if entry {
			arr[i] = 0
		} else {
			seen[v] = struct{}{}
		}
	}

	return arr

}

// Unique возвращает новый массив, содержащий только уникальные значения из исходного массива
func Unique(arr []int) (res []int) {
	seen := make(map[int]struct{})

	for _, v := range arr {
		_, entry := seen[v]

		if !entry {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}

var ages = map[string]int{
	"dad":  44,
	"mom":  40,
	"lexa": 24,
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

	seen := make(map[int]int)

	for _, num := range arr {
		seen[num]++
	}

	maxCount := 0
	element := 0

	for key, value := range seen {
		if value > maxCount {
			maxCount = value
			element = key
		}
	}

	return element
}

func MiddleMarks(marks map[string][]int) map[string]int {

	midMarks := make(map[string]int)

	for key, marks := range marks {
		midMarks[key] = arrays.MidVal(marks)
	}

	return midMarks
}

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/

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
