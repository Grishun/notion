package sort

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"notion/task/arrays/pkg/arrays"
)

func TestQuickSort(t *testing.T) {
	randArr := arrays.RandArray(1000)
	sorted := SelectionSort(randArr)

	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    randArr,
			expected: sorted,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, QuickSort(testCase.input))
		})
	}
}
