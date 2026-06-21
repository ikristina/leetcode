package problems

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		input1  string
		input2  string
		expected int
	}{
		{
			input1:  "abcde",
			input2:  "ace",
			expected: 3,
		},
		{
			input1:  "abc",
			input2:  "def",
			expected: 0,
		},
		{
			input1:  "",
			input2:  "",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		result := longestCommonSubsequence(tc.input1, tc.input2)
		if result != tc.expected {
			t.Errorf("longestCommonSubsequence(%s, %s) = %d; expected %d", tc.input1, tc.input2, result, tc.expected)
		}
	}
}
