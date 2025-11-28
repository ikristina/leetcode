package problems

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := map[string]struct {
		coins  []int
		amount int
		want   int
	}{
		"Example 1": {
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		"Example 2": {
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		"Example 3": {
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		"Large amount": {
			coins:  []int{1, 2, 5},
			amount: 100,
			want:   20,
		},
		"Unreachable": {
			coins:  []int{2, 4, 6},
			amount: 7,
			want:   -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := CoinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
