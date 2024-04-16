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

// This func generates random int matrix,
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

func SymetricMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}

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

func EnumerateMatrix(matrix Matrix, i, j int) int {
	return matrix.High*i + j
}
