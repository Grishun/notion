package leetcode

import (
	"notion/task/Algorithms/Searching/pkg/search"
	"strings"
)

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/

func DistinctInts(num1, num2 []int) (res [][]int) {

	seen := make(map[int]struct{})

	res = make([][]int, 2)

	for _, v := range num1 {
		_, err := search.BinSearch(num2, v)

		if err != nil {
			res[0] = append(res[0], v)
		} else {
			seen[v] = struct{}{}
		}
	}

	for _, v := range num2 {

		_, entry := seen[v]

		if !entry {
			res[1] = append(res[1], v)
		}
	}

	return
}

func DistinctIntsV2(num1, num2 []int) [][]int {

	res := make([][]int, 2)

	copyArrToMap := func(arr []int) map[int]struct{} {
		copied := make(map[int]struct{})
		for _, v := range arr {
			copied[v] = struct{}{}
		}

		return copied
	}

	Num2 := copyArrToMap(num2)

	for _, v := range num1 {

		_, entry := Num2[v]

		if !entry {
			res[0] = append(res[0], v)
		} else {
			delete(Num2, v)
		}
	}

	for key, _ := range Num2 {
		res[1] = append(res[1], key)
	}

	return res
}

func FindNumOfValidWords(words []string, puzzles []string) []int {

	res := make([]int, len(puzzles))

	strLettersToMap := func(s string) map[rune]struct{} {
		m := make(map[rune]struct{})
		for _, v := range s {
			m[v] = struct{}{}
		}

		return m
	}

	isValid := func(word, puzzle string) bool {
		puzzleLettersMap := strLettersToMap(puzzle)

		if !strings.Contains(word, string(puzzle[0])) {
			return false
		}

		for _, v := range word {
			_, entry := puzzleLettersMap[v]

			if !entry {
				return false
			}
		}

		return true
	}

	for i, puzzle := range puzzles {
		for _, word := range words {
			if isValid(word, puzzle) {
				res[i]++
			}
		}
	}

	return res
}

func GetBitMask(word string) int32 {
	var mask int32 = 0
	for _, char := range word {
		letterNum := char - 'a'
		mask |= 1 << letterNum
	}
	return mask
}

// FindNumOfValidWordsV2 https://leetcode.com/problems/number-of-valid-words-for-each-puzzle
func FindNumOfValidWordsV2(words []string, puzzles []string) []int {

	wordsMasks := make([]int32, len(words))

	for i, word := range words {
		wordsMasks[i] = GetBitMask(word)
	}

	res := make([]int, len(puzzles))

	for i, puzzle := range puzzles {
		puzzleMask := GetBitMask(puzzle)
		firstCharMask := GetBitMask(string(puzzle[0]))
		for _, wordMask := range wordsMasks {
			if wordMask&puzzleMask == wordMask && firstCharMask&wordMask != 0 {
				res[i]++
			}
		}
	}

	return res
}
