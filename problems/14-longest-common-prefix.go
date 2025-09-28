package problems

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	compare := strs[0]

	for _, str := range strs {
		for i := 0; i < len(compare); i++ {
			if i >= len(str) || compare[i] != str[i] {
				compare = compare[:i]
				break
			}
		}
	}

	return compare
}
