package problems

import "fmt"

func strStr(haystack string, needle string) int {

	needleLen := len(needle)
	if len(haystack) == needleLen {
		if haystack == needle {
			return 0
		}
	}

	if needleLen == 0 {
		return -1
	}

	for i, _ := range haystack {

		if i+needleLen > len(haystack) {
			fmt.Println(i, needleLen)
			return -1
		}
		if haystack[i:i+needleLen] == needle {
			return i
		}

	}

	return -1
}
