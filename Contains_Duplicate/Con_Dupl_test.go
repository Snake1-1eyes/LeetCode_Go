package Con_Dupl

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result bool
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 2, 3, 1},
			result: true,
		},
		{
			name:   "Example 2",
			nums:   []int{1, 2, 3, 4},
			result: false,
		},
		{
			name:   "Example 3",
			nums:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			result: true,
		},
		{
			name:   "Empty slice",
			nums:   []int{},
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.result {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.result)
			}
		})
	}
}
