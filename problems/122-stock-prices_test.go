package problems

import (
	"testing"
)

func Test122(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{7, 1, 5, 3, 6, 4},
			want:  7,
		},

		{
			input: []int{1, 2, 3, 4, 5},
			want:  4,
		},

		{
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		{
			input: []int{6, 1, 3, 2, 4, 7},
			want:  2,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := MaxProfit(tc.input)
			if got != tc.want {
				t.Errorf("test failed")
			}
		})
	}
}
