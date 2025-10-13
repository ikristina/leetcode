package problems

import "testing"

func TestIsSubscequence(t *testing.T) {
	tests := map[string]struct {
		s    string
		t    string
		want bool
	}{
		"case 1": {"abc", "ahbgdc", true},
		"case 2": {"axc", "ahbgdc", false},
		"case 3": {"", "ahbgdc", true},
		"case 4": {"b", "abc", true},
		"case 5": {"aaaaaa", "bbaaaa", false},
		"case 6": {"acb", "ahbgdc", false},
		"case 7": {"bb", "ahbgdc", false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isSubsequence(test.s, test.t)
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
