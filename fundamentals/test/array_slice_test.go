package fundamentalstest

import (
	"reflect"
	"testing"
)

func Sum(numbers []int) int {
	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}

func TestSum(t *testing.T) {
	testCases := []struct {
		desc     string
		numbers  []int
		expected int
	}{
		{
			desc:     "sum numbers 1",
			numbers:  []int{1, 2, 3},
			expected: 6,
		},
		{
			desc:     "sum numbers 2",
			numbers:  []int{5, 2, 3, 4},
			expected: 14,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := Sum(tC.numbers)

			if tC.expected != result {
				t.Error("expected:", tC.expected, "got:", result)
			}
		})
	}
}

// numbersToSum... => variadic function
func SumAll(numbersToSum ...[]int) []int {
	result := []int{}

	for _, v := range numbersToSum {
		sumOfNumber := Sum(v)
		result = append(result, sumOfNumber)
	}

	return result
}

func TestSumAll(t *testing.T) {
	testCases := []struct {
		desc     string
		numbers1 []int
		numbers2 []int
		expected []int
	}{
		{
			desc:     "sums slices 1",
			numbers1: []int{1, 2, 3},
			numbers2: []int{4, 5, 6},
			expected: []int{6, 15},
		},
		{
			desc:     "sums slices 2",
			numbers1: []int{1, 3},
			numbers2: []int{4, 5, 6},
			expected: []int{4, 15},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SumAll(tC.numbers1, tC.numbers2)

			if !reflect.DeepEqual(result, tC.expected) {
				t.Error("got", result, "want", tC.expected)
			}
		})
	}
}
