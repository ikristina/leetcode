package problems

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int) // letter and how many letters
	for _, c := range magazine {
		magazineMap[c]++
	}

	for _, c := range ransomNote {
		if n, ok := magazineMap[c]; !ok || n == 0 {
			return false
		}
		magazineMap[c]--
	}

	return true
}
