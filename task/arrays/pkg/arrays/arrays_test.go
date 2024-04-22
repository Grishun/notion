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
