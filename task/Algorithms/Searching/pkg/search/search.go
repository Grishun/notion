package search

import "errors"

func LineSearch(arr []int, value int) (index int, err error) {
	index = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			index = i
		}
	}

	if index == -1 {
		return 0, errors.New("no value")
	}

	return index, nil
}

func BinSearch(arr []int, value int) (index int, err error) {
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2

	for arr[mid] != value {
		if arr[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (low + high) / 2

		if low > high {
			mid, err = 0, errors.New("no value")
			break
		}
	}

	return mid, err
}
