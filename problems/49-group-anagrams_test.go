package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  [][]string
	}{
		"case 1": {
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
		},
		"case 2": {
			input: []string{""},
			want:  [][]string{{""}},
		},
		"case 3": {
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.input[0], func(t *testing.T) {
			assert.ElementsMatch(t, tc.want, groupAnagrams(tc.input))
		})
	}
}
