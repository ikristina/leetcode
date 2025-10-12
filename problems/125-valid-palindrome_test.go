package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"case1": {"A man, a plan, a canal: Panama", true},
		"case2": {"race a car", false},
		"case3": {" ", true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
