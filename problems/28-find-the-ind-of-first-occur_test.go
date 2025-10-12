package problems

import "testing"

func TestStrStr(t *testing.T) {
	tests := map[string]struct {
		haystack, needle string
		want             int
	}{
		"case1": {"hello", "ll", 2},
		"case2": {"aaaaa", "bba", -1},
		"case3": {"", "", 0},
		"case4": {"a", "a", 0},
		"case5": {"mississippi", "issip", 4},
	}

	for name, test := range tests {
		got := strStr(test.haystack, test.needle)
		if got != test.want {
			t.Errorf("StrStr(%s, %s) = %d, want %d. Test: %s", test.haystack, test.needle, got, test.want, name)
		}
	}
}
