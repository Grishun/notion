package tasks

import (
	"fmt"
)

func task3() (res []int) {

	for i := 100; i <= 999; i++ {
		if i%99 == 0 && i%10 == 7 {
			res = append(res, i)
		}
	}

	return
}

func task4() {
	for i := 1; i <= 9; i++ {
		for j := 2; j <= 9; j++ {
			fmt.Printf("%v * %v = %v	", j, i, i*j)
		}
		fmt.Println()
	}
}
