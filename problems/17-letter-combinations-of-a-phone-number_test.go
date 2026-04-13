package problems

import "testing"

func TestLetterCombinations(t *testing.T) {
	tests := map[string]struct {
		digits string
		want   []string
	}{
		"empty": {
			digits: "",
			want:   []string{},
		},
		"single": {
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		"two": {
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"three": {
			digits: "234",
			want:   []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := letterCombinations(tc.digits)
			if !equal(got, tc.want) {
				t.Errorf("letterCombinations(%s) = %v, want %v", tc.digits, got, tc.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
