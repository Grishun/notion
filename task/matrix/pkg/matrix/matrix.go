package matrix

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Matrix [][]int
	Width  int
	High   int
}

// RandMatrix func generates random int matrix,
// with size you put in it
func RandMatrix(matrix Matrix, n int) Matrix {
	matrix.Matrix = make([][]int, matrix.Width)

	for i := 0; i < matrix.Width; i++ {
		matrix.Matrix[i] = make([]int, matrix.High) // Создание width столбцов в каждой строке
		for j := 0; j < matrix.High; j++ {
			matrix.Matrix[i][j] = rand.Intn(n)
		}
	}

	return matrix
}

// This procedure outputs a matrix into the terminal, as a table

func Table(matrix Matrix) {
	for i := 0; i < len(matrix.Matrix); i++ {
		for j := 0; j < len(matrix.Matrix[i]); j++ {
			fmt.Printf("%v\t", matrix.Matrix[i][j])
		}
		fmt.Println()
	}
}

func TableV2(matrix Matrix) {
	for i := 0; i < len(matrix.Matrix); i++ {
		fmt.Println(matrix.Matrix[i])
	}
}

// GenerateMatrix generates matrix, and fills it with the values,
// you describe in the formula
// Example:
// GenerateMatrix([2][3]matrix, func(i,j int) {return 2^enumerate(matrix, i, j)})
// -> [[1, 2, 4], [8, 16, 32]]
func GenerateMatrix(matrix Matrix, formula func(i, j int) int) Matrix {
	matrix.Matrix = make([][]int, matrix.Width)

	for i := 0; i < matrix.Width; i++ {
		matrix.Matrix[i] = make([]int, matrix.High)
		for j := 0; j < matrix.High; j++ {
			matrix.Matrix[i][j] = formula(i, j)
		}
	}

	return matrix
}

// EnumerateMatrix enumerates matrix from its (0,0) element by chain
// to the last element.
// Example:
// matrix = [[1,2,3], [4,5,6], [7,8,9]]
// EnumerateMatrix(matrix, 0, 0) -> 0
// EnumerateMatrix(matrix, 1, 2) -> 5
// EnumerateMatrix(matrix, 3, 3) -> 8
func EnumerateMatrix(matrix Matrix, i, j int) int {
	return matrix.High*i + j
}

func SumOfElements(matrix Matrix) (res int) {
	for i := range matrix.Matrix {
		for _, v := range matrix.Matrix[i] {
			res += v
		}
	}

	return
}
