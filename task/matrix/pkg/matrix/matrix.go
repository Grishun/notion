package matrix

import (
	"fmt"
	"math/rand"
)

// This func generates random int matrix,
// with size you put in it
func RandMatrix(strings, columns int) [][]int {
	res := make([][]int, strings)

	for i := 0; i < strings; i++ {
		res[i] = make([]int, columns) // Создание width столбцов в каждой строке
		for j := 0; j < columns; j++ {
			res[i][j] = rand.Intn(9)
		}
	}

	return res
}

// This procedure outputs a matrix into the terminal, as a table
func Table(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func TableV2(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
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
