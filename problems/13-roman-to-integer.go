package problems

func romanToInteger(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0

	for i, r := range s {
		c := string(r)
		// XIV - 14, 10 - 1 + 5
		if len(s)-1 > i {
			next := romans[string(s[i+1])]
			if romans[c] < next {
				res -= romans[c]
				continue
			}
		}

		res += romans[c]
	}
	return res
}
