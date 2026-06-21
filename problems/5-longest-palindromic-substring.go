package problems


func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	res := [2]int{0, 0}
	dp := make([][]bool, len(s))
	for i := range len(s) {
		dp[i] = make([]bool, len(s))
	}

	for i := range len(s) {
		dp[i][i] = true
	}

	for i := range len(s)-1 {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			res = [2]int{i, i+1}
		}
	}

	// for 3 letters and more
	for length := 3; length <= len(s); length++ {
		i := 0
		for j := length-1; j < len(s); j++ {
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				res = [2]int{i, j}
			}
			i++
		}

	}

	return s[res[0]:res[1]+1]
}
