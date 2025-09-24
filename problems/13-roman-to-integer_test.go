package problems

import "testing"

func TestRomanToInteger(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected int
	}{
		"III": {
			input:    "III",
			expected: 3,
		},
		"LVIII": {
			input:    "LVIII",
			expected: 58,
		},
		"MCMXCIV": {
			input:    "MCMXCIV",
			expected: 1994,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.expected != romanToInteger(tc.input) {
				t.Errorf("%s: expected %d, got %d", name, tc.expected, romanToInteger(tc.input))
			}
		})
	}
}
