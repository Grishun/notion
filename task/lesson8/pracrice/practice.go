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
