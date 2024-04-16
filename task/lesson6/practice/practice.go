package practice

import (
	"fmt"
	"math"
)

func practice1() {
	for i := 4; i <= 20; i += 2 {
		fmt.Println(i)
	}
}

func practice2(startRange, endRange, div int) (res []int) {

	for i := startRange; i <= endRange; i++ {
		if i%div == 0 {
			res = append(res, i)
		}
	}

	return
}

func practice3(startRange, endRange, power float64) (res []float64) {
	for i := endRange; i >= startRange; i-- {
		res = append(res, math.Pow(float64(i), power))
	}

	return
}

func practice4(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}

func practice5(startRange, endRange, div int) (res int) {
	for i := startRange; i <= endRange; i++ {
		if i%div == 0 {
			res = i
			break
		}
	}

	return
}
