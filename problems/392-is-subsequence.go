package problems

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	if len(s) > len(t) {
		return false
	}
	if len(s) == len(t) {
		return s == t
	}

	p1, p2 := 0, 0

	for p1 < len(s) && p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
		}
		p2++
	}

	return p1 == len(s)
}
