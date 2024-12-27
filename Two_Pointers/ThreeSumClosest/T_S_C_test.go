package threesumclosest

import "testing"

// ...existing code...
func TestThreeSumCosest(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2, // -1 + 2 + 1 = 2
		},
		{
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0, // 0 + 0 + 0 = 0
		},
		{
			nums:   []int{1, 1, 1},
			target: 3,
			want:   3, // 1 + 1 + 1 = 3
		},
		{
			nums:   []int{-3, -1, 2, 4, 5},
			target: 1,
			want:   1,
		},
	}

	for _, tt := range tests {
		got := threeSumCosest(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("threeSumCosest(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
