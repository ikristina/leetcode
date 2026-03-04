package problems

import (
	"maps"
	"slices"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	resMap := map[string][]string{}

	for _, word := range strs {
		count := [26]int{}

		for _, c := range word {
			count[c-'a']++
		}

		key := ""
		for _, v := range count {
			key += strconv.Itoa(v) + "#"
		}

		resMap[key] = append(resMap[key], word)

	}

	return slices.Collect(maps.Values(resMap))
}
