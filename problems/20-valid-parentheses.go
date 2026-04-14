package problems
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	closeToOpen := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	stack := []rune{}

	for _, r := range s {

		if opener, ok := closeToOpen[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opener {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}

	}
	return len(stack) == 0
}
