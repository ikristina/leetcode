package problems

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		expected []string
	}{
		"Example 1": {nums: []int{0, 1, 2, 4, 5, 7}, expected: []string{"0->2", "4->5", "7"}},
		"Example 2": {nums: []int{0, 2, 3, 4, 6, 8, 9}, expected: []string{"0", "2->4", "6", "8->9"}},
		"Empty":     {nums: []int{}, expected: []string{}},
		"Negative":  {nums: []int{-1}, expected: []string{"-1"}},
		"Zero":      {nums: []int{0}, expected: []string{"0"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := summaryRanges(test.nums)
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
			}
		})
	}
}
