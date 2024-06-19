package cycles

import (
	"github.com/Grishun/problems/task/math/pkg/numbers"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestRemoveDigs(t *testing.T) {
	type inp struct {
		num int
		dig int
	}

	testCases := []struct {
		input    inp
		expected int
	}{
		{
			input: inp{
				num: 12313,
				dig: 3,
			},
			expected: 121,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveDigs(testCase.input.num, testCase.input.dig))
		})
	}
}

func TestRemoveDigByIndex(t *testing.T) {
	type inp struct {
		num   int
		index int
	}

	testCases := []struct {
		input    inp
		expected int
	}{
		{
			input: inp{
				num:   123,
				index: 0,
			},
			expected: 12,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveDigByIndex(testCase.input.num, testCase.input.index))
		})
	}
}

func TestPrimeNums(t *testing.T) {
	isPrimes := func(arr []int) bool {

		for _, v := range arr {
			if !numbers.IsPrime(v) {
				return false
			}
		}

		return true
	}

	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    PrimeNums(50),
			expected: true,
		},
		{
			input:    PrimeNums(0),
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, isPrimes(PrimeNums(50)))
		})
	}
}

func TestDecompNums(t *testing.T) {

	isDiv := func(arr []int, num uint64) bool {
		for _, v := range arr {
			if int(num)%v != 0 {
				return false
			}
		}

		return true
	}

	testCases := []struct {
		input    uint64
		expected bool
	}{
		{
			input:    10000,
			expected: true,
		},
		{
			input:    0,
			expected: true,
		},
		{
			input:    11,
			expected: true,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, isDiv(DecompNums(testCase.input), testCase.input))
		})
	}
}

func TestPerfectNums(t *testing.T) {

	isPerfect := func(num uint64) bool {

		if num < 6 {
			return false
		}
		sumOfDivs := 0

		for j := 0; j < len(DecompNums(num))-1; j++ {
			sumOfDivs += DecompNums(num)[j]
		}

		return sumOfDivs == int(num)
	}

	testCases := []struct {
		input    uint64
		expected bool
	}{
		{
			input:    6,
			expected: true,
		},
		{
			input:    0,
			expected: false,
		},
		{
			input:    11,
			expected: false,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, isPerfect(testCase.input))
		})
	}
}
