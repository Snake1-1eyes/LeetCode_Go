package Plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{8}, []int{9}},
		{[]int{9}, []int{1, 0}},
		{[]int{5, 9}, []int{6, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, test := range tests {
		result := plusOne(test.digits)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For %v, expected %v but got %v", test.digits, test.expected, result)
		}
	}
}
