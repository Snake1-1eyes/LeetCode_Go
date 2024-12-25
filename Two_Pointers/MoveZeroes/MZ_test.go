package Move_Z

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{[]int{1, 2, 3, 0}, []int{1, 2, 3, 0}},
		{[]int{0, 0, 0, 1, 0, 2}, []int{1, 2, 0, 0, 0, 0}},
		{[]int{1, 0, 0, 0, 2}, []int{1, 2, 0, 0, 0}},
	}

	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		moveZeroes(input)

		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("For %v, expected %v but got %v", tc.input, tc.expected, input)
		}
	}
}
