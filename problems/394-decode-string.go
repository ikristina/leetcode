package problems

import "strings"

func decodeString(s string) string {
	if s == "" {
		return ""
	}
	// using stack
	//
	stack := []rune{}

	for _, char := range s {
		if char != ']' {
			stack = append(stack, char)
			continue
		}
		var decodedString string
		var c rune
		c, stack = pop(stack)

		for c != '[' {
			decodedString = string(c) + decodedString
            c, stack = pop(stack)
		}

		k := 0
		base := 1

		c, stack = pop(stack)
		for c >= '0' && c <= '9'  {
			digit := int(c - '0')
			k += digit * base
			base = base * 10
            if len(stack) == 0 {
                break
            }
            c, stack = pop(stack)
		}
        if c < '0' || c > '9' {
            stack = append(stack, c)
        }
		var repeated strings.Builder
	 	for _ = range k {
			repeated.WriteString(decodedString)
		}
		for _, c := range repeated.String() {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func pop(stack []rune) (rune, []rune) {
    if len(stack) == 0 {
        return 0, stack
    }
    top := stack[len(stack)-1]
    return top, stack[:len(stack)-1]
}
