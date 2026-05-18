package problems

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := map[string]struct {
		tasks   []byte
		n       int
		expected int
	}{
		"example 1": {
			tasks:   []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:       2,
			expected: 8,
		},
		"example 2": {
			tasks:   []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:       1,
			expected: 6,
		},
		"example 3": {
			tasks:   []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:       3,
			expected: 10,
		},

	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := leastInterval(test.tasks, test.n)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
