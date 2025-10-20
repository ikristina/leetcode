package problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"case 1": {
			input: "abcabcbb",
			want:  3,
		},
		"case 2": {
			input: "bbbbb",
			want:  1,
		},
		"case 3": {
			input: "pwwkew",
			want:  3,
		},
		"case 4": {
			input: " ",
			want:  1,
		},
		"case 5": {
			input: "au",
			want:  2,
		},
		"case 6": {
			input: "abb",
			want:  2,
		},
		"case 7": {
			input: "dvdf",
			want:  3,
		},
		"empty string": {
			input: "",
			want:  0,
		},
		"single char": {
			input: "a",
			want:  1,
		},
		"all unique chars": {
			input: "abcdef",
			want:  6,
		},
		"repeating pattern": {
			input: "abcabcabc",
			want:  3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.input)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
