package problems

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected int
	}{
		"Example 1": {s: "Hello World", expected: 5},
		"Example 2": {s: "   fly me   to   the moon  ", expected: 4},
		"Example 3": {s: "luffy is still joyboy", expected: 6},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := lengthOfLastWord(test.s)
			if actual != test.expected {
				t.Errorf("s: %s, expected: %d, actual: %d", test.s, test.expected, actual)
			}
		})
	}
}
