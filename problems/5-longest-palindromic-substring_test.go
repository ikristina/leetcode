package problems

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "aba",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input:    "ac",
			expected: "a",
		},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("longestPalindrome(%s) = %s; expected %s", tc.input, result, tc.expected)
		}
	}
}
