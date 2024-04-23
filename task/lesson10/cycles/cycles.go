package cycles

import (
	"github.com/Grishun/problems/task/math/pkg/numbers"
	"math"
	"strconv"
)

// This func returns number of even and odd digits in a number
// Example
// CountDigit(123) -> 1, 2

func CountDigits(num int) (evens, ods int) {

	for i := 0; i < len(strconv.Itoa(num)); i++ {
		if numbers.DigOfNum(uint64(num), uint64(i))%2 == 0 {
			evens++
		}
	}

	return evens, len(strconv.Itoa(num)) - evens
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
