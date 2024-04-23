package search

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var testArr = func(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, i)
	}

	return
}

func TestLineSearch(t *testing.T) {
	type input struct {
		arr   []int
		value int
	}

	type expected struct {
		index int
		err   error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				arr:   testArr(100),
				value: 99,
			},
			exp: expected{
				index: 99,
				err:   nil,
			},
		}, {
			inp: input{
				arr:   testArr(100),
				value: 0,
			},
			exp: expected{
				index: 0,
				err:   nil,
			},
		}, {
			inp: input{
				arr:   testArr(100),
				value: 100,
			},
			exp: expected{
				index: 0,
				err:   errors.New("no value"),
			},
		}, {
			inp: input{
				arr:   []int{},
				value: 0,
			},
			exp: expected{
				index: 0,
				err:   errors.New("no value"),
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			actualIndex, actualErr := LineSearch(testCase.inp.arr, testCase.inp.value)
			require.Equal(t, testCase.exp.index, actualIndex)
			require.Equal(t, testCase.exp.err, actualErr)
		})
	}
}

func TestBinSearch(t *testing.T) {
	type input struct {
		arr   []int
		value int
	}

	type expected struct {
		index int
		err   error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				arr:   testArr(100),
				value: 99,
			},
			exp: expected{
				index: 99,
				err:   nil,
			},
		}, {
			inp: input{
				arr:   testArr(100),
				value: 0,
			},
			exp: expected{
				index: 0,
				err:   nil,
			},
		}, {
			inp: input{
				arr:   testArr(100),
				value: 100,
			},
			exp: expected{
				index: 0,
				err:   errors.New("no value"),
			},
		}, {
			inp: input{
				arr:   []int{},
				value: 0,
			},
			exp: expected{
				index: 0,
				err:   errors.New("empty array"),
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			actualIndex, actualErr := BinSearch(testCase.inp.arr, testCase.inp.value)
			require.Equal(t, testCase.exp.index, actualIndex)
			require.Equal(t, testCase.exp.err, actualErr)
		})
	}
}
