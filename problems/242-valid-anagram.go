package problems

import "fmt"

// IsAnagram leetcode problem: https://leetcode.com/problems/valid-anagram/
// easy
// O(n) time and O(1) space
// #hashTable #map #string #sorting
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := map[rune]int{}
	mt := map[rune]int{}
	for _, r := range s {
		ms[r] += 1
	}

	for _, r := range t {
		mt[r] += 1
	}

	for r, v := range ms {
		if mt[r] != v {
			return false
		}
	}
	return true
}

// O(1) fixed array vs. O(k) map (k = unique characters).
func isAnagramOption2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [256]int{} // Fixed-size array for ASCII (or 52 for letters)
	for i := range s {
		fmt.Println(s[i])
		count[s[i]]++
		count[t[i]]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramOption3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var abc [26]int // 26 letters of alphabet

	for i := 0; i < len(s); i++ {
		abc[s[i]-'a']++
		abc[t[i]-'a']--
	}

	for _, l := range abc {
		if l != 0 {
			return false
		}
	}

	return true
}
