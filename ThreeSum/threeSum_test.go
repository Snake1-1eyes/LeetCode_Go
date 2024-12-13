package ThreeSum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			nums:     []int{},
			expected: [][]int{},
		},
		{
			nums:     []int{0},
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		result := threeSum(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums %v, expected %v, but got %v", test.nums, test.expected, result)
		}
	}
}
