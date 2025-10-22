package problems

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := map[string]struct {
		target int
		nums   []int
		want   int
	}{
		"case 1": {
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		"case 2": {
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		"case 3": {
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
		"case 4": {
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			want:   3,
		},
		"case 5": {
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			want:   5,
		},
		"case 6": {
			target: 6,
			nums:   []int{10, 2, 3},
			want:   1,
		},
		"empty array": {
			target: 5,
			nums:   []int{},
			want:   0,
		},
		"single element equal to target": {
			target: 5,
			nums:   []int{5},
			want:   1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
