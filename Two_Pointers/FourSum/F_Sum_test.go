package foursum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	type testCase struct {
		nums     []int
		target   int
		expected [][]int
	}
	tests := []testCase{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			expected: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			expected: [][]int{
				{2, 2, 2, 2},
			},
		},
	}

	for _, tc := range tests {
		got := fourSum(tc.nums, tc.target)
		// Поскольку порядок может отличаться, простая проверка reflect.DeepEqual
		// может не подойти в некоторых случаях. Для наглядности здесь — простой вариант:
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("For nums %v and target %d, expected %v, got %v", tc.nums, tc.target, tc.expected, got)
		}
	}
}
