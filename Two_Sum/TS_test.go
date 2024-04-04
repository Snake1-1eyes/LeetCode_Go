package Two_Sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
	}

	for _, tc := range testCases {
		input := make([]int, len(tc.nums))
		copy(input, tc.nums)

		result := twoSum(tc.nums, tc.target)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For nums=%v and target=%d, expected %v but got %v", tc.nums, tc.target, tc.expected, result)
		}
	}
}
