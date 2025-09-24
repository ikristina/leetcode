package problems

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
	for _, test := range tests {
		actual := lengthOfLastWord(test.s)
		if actual != test.expected {
			t.Errorf("s: %s, expected: %d, actual: %d", test.s, test.expected, actual)
		}
	}
}
