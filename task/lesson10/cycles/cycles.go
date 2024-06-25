package cycles

import (
	"github.com/Grishun/problems/task/math/pkg/numbers"
	"math"
	"sort"
	"strconv"
)

// CountDigits returns number of even and odd digits in a number
// Example
// CountDigit(123) -> 1, 2
//func CountDigits(num int) (evens, ods int) {
//
//	for i := 0; i < len(strconv.Itoa(num)); i++ {
//		if numbers.DigOfNum(uint64(num), uint64(i))%2 == 0 {
//			evens++
//		}
//	}
//
//	return evens, len(strconv.Itoa(num)) - evens
//}

// RemoveDigs removes the digit from number
// Example
// RemoveDigByIndex(12313, 3) -> 121

func RemoveDigs(num, dig int) (res int) {

	digRank := 1
	for num != 0 {
		currentDig := num % 10
		num /= 10
		if currentDig != dig {
			res += currentDig * digRank
			digRank *= 10
		}
	}

	return
}

// RemoveDigByIndex removes the digit from number by index,
// Example
// RemoveDigByIndex(1234, 0) -> 123
// RemoveDigByIndex(1234, 1) -> 124
// RemoveDigByIndex(1234, 2) -> 134
func RemoveDigByIndex(num, index int) (res int) {

	sign := 1
	switch {
	case num < 0:
		num = -num
		sign = -1
	case index < 0:
		return
	}

	flag := 0
	for i := 0; i < len(strconv.Itoa(num)); i++ {
		if i == index {
			flag = 1
			continue
		}

		res += int(numbers.DigOfNum(uint64(num), uint64(i)) * uint64(math.Pow10(i-flag)))
	}

	return res * sign
}

func PrimeNums(n int) (res []int) {
	if n <= 1 {
		return
	}

	isPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p < n; p++ {
		if isPrime[p] {
			for mult := p * p; mult < n; mult += p {
				isPrime[mult] = false
			}
		}
	}

	for i, v := range isPrime {
		if v {
			res = append(res, i)
		}
	}

	return
}

func DecompNums(num uint64) (divs []int) {

	for i := 1; i*i <= int(num); i++ {
		if int(num)%i == 0 {
			divs = append(divs, i)

			if i != int(num)/i {
				divs = append(divs, int(num)/i)
			}
		}
	}

	sort.Ints(divs)

	return

}

func PerfectNums(n int) (res []int) {

	for i := 6; i < n; i++ {
		sumOfDivs := 0
		for j := 0; j < len(DecompNums(uint64(i)))-1; j++ {
			sumOfDivs += DecompNums(uint64(i))[j]
		}

		if sumOfDivs == i {
			res = append(res, i)
		}
	}

	return
}

//func Sirakuza(num int) {
//
//	fmt.Printf("%d -> ", num)
//
//	for num != 2 {
//		if num%2 == 0 {
//			num /= 2
//			fmt.Printf("%d -> ", num)
//		} else {
//			num = num*3 + 1
//			fmt.Printf("%d -> ", num)
//		}
//	}
//
//	fmt.Printf("1")
//
//}
