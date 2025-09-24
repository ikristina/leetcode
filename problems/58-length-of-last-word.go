package problems

// import "fmt"

func lengthOfLastWord(s string) int {
	s = trim(s)
	wordLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		// fmt.Println("symbol: ", string(s[i]))
		if string(s[i]) == " " {
			return wordLength
		}
		wordLength++
	}
	return len(s)
}

func trim(s string) string {

	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			break
		}
		s = s[i:]
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			break
		}
		s = s[:i]
	}
	return s
}
