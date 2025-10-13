package problems

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := map[string]struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		"case1": {
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		"case2": {
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		"case3": {
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
