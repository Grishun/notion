package arrays

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var (
	testSlice = func(n int) (res []int) {
		for i := 0; i < n; i++ {
			res = append(res, i, i)
		}

		return
	}

	numbers = func(n int) (res []int) {
		for i := 0; i < n; i++ {
			res = append(res, i)
		}

		return
	}
)

func TestRemoveByIndex(t *testing.T) {
	type input struct {
		arr   []int
		index int
	}

	testCases := []struct {
		inp      input
		expected []int
	}{
		{
			inp: input{
				arr:   []int{0, 1, 2, 3, 4},
				index: 2,
			},
			expected: []int{0, 1, 3, 4},
		},
		{
			inp: input{
				arr:   []int{},
				index: 0,
			},
			expected: []int{},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveByIndex(testCase.inp.arr, testCase.inp.index))
		})
	}
}

func TestRemoveByValue(t *testing.T) {
	type input struct {
		arr   []int
		value int
	}

	testCases := []struct {
		inp      input
		expected []int
	}{
		{
			inp: input{
				arr:   []int{0, 1, 2, 3, 4},
				value: 2,
			},
			expected: []int{0, 1, 3, 4},
		},
		{
			inp: input{
				arr:   []int{},
				value: 0,
			},
			expected: []int{},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveByValue(testCase.inp.arr, testCase.inp.value))
		})
	}
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
			removed := RemoveRepeating(testCase.input)
			require.Len(t, removed, testCase.expectedLen)
			require.Equal(t, testCase.expectedCap, cap(removed))
			require.Equal(t, testCase.expected, RemoveRepeating(testCase.input))
		})
	}
}

func TestUniteArrays(t *testing.T) {
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
				arr1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{2, 4, 6, 8, 10, 12},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, UniteArrays(testCase.inp.arr1, testCase.inp.arr2))
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
