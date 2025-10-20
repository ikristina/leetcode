package problems

func lengthOfLongestSubstring(s string) int {
	// dynamic sliding window?
	// not sure what else to improve, it says it beats only 35% of other solutions for runtime.

	ht := map[byte]bool{}

	start := 0
	end := 0

	window := 0

	for end < len(s) {
		if _, ok := ht[s[end]]; !ok {
			ht[s[end]] = true
			end++
			if window <= len(s[start:end]) {
				window = end - start
			}

		} else {
			delete(ht, s[start])
			start++
		}
	}
	return window
}
