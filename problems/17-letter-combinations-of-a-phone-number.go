package problems

func letterCombinations(digits string) []string {

    if len(digits) == 0 {
        return []string{}
    }

    var result []string

    mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

    var bt func(index int, current string)
    bt = func(index int, current string) {
        if index == len(digits) {
            result = append(result, current)
            return
        }

        digit := digits[index] - '0'
        letters := mapping[digit]

        for i := 0; i < len(letters); i++ {
            // add the letter and move to the next digit
            bt(index+1, current + string(letters[i]))
        }

    }

    bt(0, "")
    return result
}
