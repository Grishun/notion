package cycles

import (
	"math"
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

func IsPrime(num int) bool {
	if num == 0 || num == 1 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
