package arrays

import (
	"errors"
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

// This func counts how many times a number appears in an array
func Count(arr []int, value int) (res int) {
	for _, v := range arr {
		if v == value {
			res++
		}
	}

	return res
}

var letters = []string{"а", "б", "в", "г", "д", "е", "ё", "ж", "з",
	"и", "й", "к", "л", "м", "н", "о", "п", "р",
	"с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ",
	"ъ", "ы", "ь", "э", "ю", "я", "!", ",", " "}

func Message(message []int) (res string, err error) {

	for _, v := range message {
		if v > 35 || v < 0 {
			err = errors.New("out of range")
		}
		res += letters[v]
	}

	return
}
