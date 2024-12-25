package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
		{
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			height:   []int{2, 3, 10, 5, 7, 8, 9},
			expected: 36,
		},
	}

	for _, test := range tests {
		result := maxArea(test.height)
		if result != test.expected {
			t.Errorf("For input %v, expected %d, but got %d", test.height, test.expected, result)
		}
	}
}
