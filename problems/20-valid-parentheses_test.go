package problems

import "testing"

func TestIsValid(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"": {
			input:    "",
			expected: false,
		},
		"()": {
			input:    "()",
			expected: true,
		},
		"()[]": {
			input:    "()[]",
			expected: true,
		},
		"([(]": {
			input:    "([(]",
			expected: false,
		},
		"([])": {
			input:    "([])",
			expected: true,
		},
		"([{}])": {
			input:    "([{}])",
			expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := isValid(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
