package main

import (
	"notion/task/matrix/pkg/matrix"
)

func main() {
	mtrx := matrix.Matrix{
		Matrix: make([][]int, 3),
		Width:  3,
		High:   4,
	}

	matrix.Table(matrix.GenerateMatrix(mtrx, func(i, j int) int {
		return matrix.EnumerateMatrix(mtrx, i, j) * 2
	}))

}
