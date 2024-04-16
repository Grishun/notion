package pracrice

import (
	"errors"
	"notion/task/arrays/pkg/arrays"
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

func Insert(arr []int, value, index int) (res []int, err error) {

	if index < 0 || index > len(arr) {
		return []int{}, errors.New("error")
	}
	res = append(res, arr[:index]...)
	res = append(res, value)
	res = append(res, arr[index:]...)

	return res, nil
}
