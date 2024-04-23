package arrays

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var testSlice = func(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, i, i)
	}

	return
}

var numbers = func(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	return
}

func TestRemoveRepeating(t *testing.T) {
	testCases := []struct {
		input       []int
		expected    []int
		expectedLen int
		expectedCap int
	}{
		{
			input:       testSlice(5),
			expected:    []int{0, 1, 2, 3, 4},
			expectedLen: 5,
			expectedCap: 8,
		},
		{
			input:       testSlice(10),
			expected:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedLen: 10,
			expectedCap: 16,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			removed := RemoveRepeatingV2(testCase.input)
			require.Len(t, removed, testCase.expectedLen)
			require.Equal(t, testCase.expectedCap, cap(removed))
			require.Equal(t, testCase.expected, RemoveRepeatingV2(testCase.input))
		})
	}
}

func TestMergeArrays(t *testing.T) {
	type input struct {
		arr1 []int
		arr2 []int
	}

	testCases := []struct {
		inp      input
		expected []int
	}{
		{
			inp: input{
				arr1: numbers(10),
				arr2: numbers(10),
			},
			expected: testSlice(10),
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, MergeArrays(testCase.inp.arr1, testCase.inp.arr2))
		})
	}
}
