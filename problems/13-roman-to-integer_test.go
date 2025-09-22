package problems

import "testing"

func TestRomanToInteger(t *testing.T) {

	tests := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if romanToInteger(name) != tc {
				t.Error("Expected ", tc, " got ", romanToInteger(name))
			}
		})
	}
}
