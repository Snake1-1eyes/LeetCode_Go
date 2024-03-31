package BT_BaS

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			name:   "Example 2",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "Example 3",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "No profit",
			prices: []int{7, 6, 5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "Single price",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "Empty slice",
			prices: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
