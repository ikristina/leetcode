package problems

import "strings"

// A map[byte]struct{} uses less memory than map[byte]bool because struct{} has zero size.
var valids = map[byte]struct{}{
	'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {}, 'h': {}, 'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {}, 'o': {}, 'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {}, 'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
	'0': {}, '1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {},
}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s) // O(N)
	start := 0
	end := len(s) - 1
	for start < end {
		if _, ok := valids[s[start]]; !ok {
			start++
			continue
		}
		if _, ok := valids[s[end]]; !ok {
			end--
			continue
		}
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
