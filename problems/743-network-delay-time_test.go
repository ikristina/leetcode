package problems

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := map[string]struct {
		times    [][]int
		n        int
		k        int
		expected int
	}{
		"Example 1": {
			times:    [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:        4,
			k:        2,
			expected: 2,
		},
		"Example 2": {
			times:    [][]int{},
			n:        0,
			k:        0,
			expected: 0,
		},
		"Example 3": {
			times:    [][]int{{1, 2, 1}},
			n:        2,
			k:        1,
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := networkDelayTime(tt.times, tt.n, tt.k)
			if result != tt.expected {
				t.Errorf("networkDelayTime() = %v, expected %v", result, tt.expected)
			}
			resultPQ := networkDelayTimePQ(tt.times, tt.n, tt.k)
			if resultPQ != tt.expected {
				t.Errorf("networkDelayTimePQ() = %v, expected %v", resultPQ, tt.expected)
			}
		})
	}
}
