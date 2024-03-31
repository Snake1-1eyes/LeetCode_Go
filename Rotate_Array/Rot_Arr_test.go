package Rot_Arr

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		result []int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:   "Example 2",
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			result: []int{3, 99, -1, -100},
		},
		{
			name:   "Rotate by length",
			nums:   []int{1, 2, 3, 4, 5},
			k:      5,
			result: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Rotate by multiple of length",
			nums:   []int{1, 2, 3, 4, 5},
			k:      10,
			result: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.result) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.result)
			}
		})
	}
}
