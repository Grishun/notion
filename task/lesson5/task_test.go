package lesson5

import (
	"errors"
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

func TestCalcuclitor(t *testing.T) {
	type input struct {
		num1 int
		num2 int
		sign string
	}
	type expected struct {
		result int
		err    error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				num1: 0,
				num2: 0,
				sign: "+",
			},
			exp: expected{
				result: 0,
				err:    nil,
			},
		},
		{
			inp: input{
				num1: 0,
				num2: 0,
				sign: "-",
			},
			exp: expected{
				result: 0,
				err:    nil,
			},
		},
		{
			inp: input{
				num1: math.MaxInt64,
				num2: math.MinInt64,
				sign: "+",
			},
			exp: expected{
				result: -1,
				err:    nil,
			},
		}, {
			inp: input{
				num1: 2,
				num2: 3,
				sign: "*",
			},
			exp: expected{
				result: 6,
				err:    nil,
			},
		}, {
			inp: input{
				num1: 10,
				num2: 2,
				sign: "/",
			},
			exp: expected{
				result: 5,
				err:    nil,
			},
		}, {
			inp: input{
				num1: -1,
				num2: math.MinInt64,
				sign: "-",
			},
			exp: expected{
				result: math.MaxInt64,
				err:    nil,
			},
		},
		{
			inp: input{
				num1: 0,
				num2: 0,
				sign: "^",
			},
			exp: expected{
				result: 0,
				err:    errors.New("don't know this operator"),
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			actualResult, actualErr := Calcuclitor(testCase.inp.num1, testCase.inp.num2, testCase.inp.sign)
			require.Equal(t, testCase.exp.err, actualErr)
			require.Equal(t, testCase.exp.result, actualResult)
		})
	}
}
