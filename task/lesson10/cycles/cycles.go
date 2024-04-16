package cycles

import (
	"strconv"
)

// This func returns number of even and odd digits in a number
// Example
// CountDigit(123) -> 1, 2

func CountDigits(num int) (evens, ods int) {

	strNum := strconv.Itoa(num)

	for _, v := range strNum {
		if (int(v)-'0')%2 == 0 {
			evens++
		}
	}
	return evens, len(strNum) - evens
}
