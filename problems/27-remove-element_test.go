package problems

import "testing"
import "fmt"

func TestRemoveElement(t *testing.T) {
	tests := map[string]struct {
		nums          []int
		val           int
		expectedArray []int
		expected      int
	}{
		"happy path 1": {
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expectedArray: []int{2, 2},
			expected:      2,
		},
		"happy path 2": {
			nums:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:           2,
			expectedArray: []int{0, 1, 3, 0, 4},
			expected:      5,
		},
	}

	for name, tc := range tests {
		t.Run(
			name, func(t *testing.T) {
				nums := tc.nums
				got := RemoveElement(nums, tc.val)
				fmt.Println(nums)
				fmt.Println("expectedArray", tc.expectedArray)
				if got != tc.expected {
					t.Errorf("got %d, want %d", got, tc.expected)
				}

			},
		)
	}
}
