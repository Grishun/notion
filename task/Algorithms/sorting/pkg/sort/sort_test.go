package sort

import (
	"github.com/stretchr/testify/require"
	"notion/task/arrays/pkg/arrays"
	"strconv"
	"testing"
)

func TestQuickSort(t *testing.T) {
	randArr := arrays.RandArray(1000)

	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    randArr,
			expected: true,
		},

		{
			input:    []int{},
			expected: true,
		},

		{
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IsSorted(QuickSort(testCase.input)))
		})
	}
}

func TestQuickSortV2(t *testing.T) {
	randArr := arrays.RandArray(1000)

	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    randArr,
			expected: true,
		},

		{
			input:    []int{},
			expected: true,
		},

		{
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IsSorted(QuickSortV2(testCase.input)))
		})
	}
}

func TestBubbleSort(t *testing.T) {
	randArr := arrays.RandArray(1000)

	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    randArr,
			expected: true,
		},

		{
			input:    []int{},
			expected: true,
		},

		{
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IsSorted(BubbleSort(testCase.input)))
		})
	}
}

func TestSelectionSort(t *testing.T) {
	randArr := arrays.RandArray(1000)

	testCases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    randArr,
			expected: true,
		},

		{
			input:    []int{},
			expected: true,
		},

		{
			input:    []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IsSorted(SelectionSort(testCase.input)))
		})
	}
}
