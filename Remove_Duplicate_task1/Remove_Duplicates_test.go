package Rem_Dup

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input          []int
		expectedResult int
		expectedArray  []int
	}{
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 2, 2, 3, 4, 4, 5, 5, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{}, 0, []int{}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{1}},
	}

	for _, test := range tests {
		result := RemoveDuplicates(test.input)
		if result != test.expectedResult {
			t.Errorf("For input %v, expected result %d, but got %d", test.input, test.expectedResult, result)
		}

		if !reflect.DeepEqual(test.input[:test.expectedResult], test.expectedArray) {
			t.Errorf("For input %v, expected array %v, but got %v", test.input, test.expectedArray, test.input[:test.expectedResult])
		}
	}
}
