package problems

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := map[string]struct {
		s, t string
		want bool
	}{
		"case1": {
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		"case2": {
			s:    "rat",
			t:    "car",
			want: false,
		},
		"case3": {
			s:    "nagaram",
			t:    "ape",
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsAnagram() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestIsAnagramOption2(t *testing.T) {
	tests := map[string]struct {
		s, t string
		want bool
	}{
		"case1": {
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		"case2": {
			s:    "rat",
			t:    "car",
			want: false,
		},
		"case3": {
			s:    "nagaram",
			t:    "ape",
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isAnagramOption2(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsAnagram() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestIsAnagramOption3(t *testing.T) {
	tests := map[string]struct {
		s, t string
		want bool
	}{
		"case1": {
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		"case2": {
			s:    "rat",
			t:    "car",
			want: false,
		},
		"case3": {
			s:    "nagaram",
			t:    "ape",
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isAnagramOption3(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsAnagram() = %v; want %v", got, tc.want)
			}
		})
	}
}
