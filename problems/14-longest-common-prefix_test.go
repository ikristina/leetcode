package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	// test cases
	testCases := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}

	for _, tc := range testCases {
		got := longestCommonPrefix(tc.strs)
		if got != tc.want {
			t.Errorf("longestCommonPrefix(%v) = %v, want %v", tc.strs, got, tc.want)
		}
	}
}
