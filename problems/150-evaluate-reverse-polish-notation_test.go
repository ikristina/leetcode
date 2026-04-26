package problems

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := map[string]struct {
		tokens []string
		want   int
	}{
		"case0": {
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		"case1": {
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		"case2": {
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := evalRPN(tc.tokens)
			if got != tc.want {
				t.Errorf("wrong result: got: %d, want:%d", got, tc.want)
			}
		})
	}
}
