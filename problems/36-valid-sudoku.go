package problems

func isValidSudoku(board [][]byte) bool {

	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// [{'val':true},{},{}]

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}
			if _, ok := rows[i][num]; ok {
				return false
			}

			if _, ok := cols[j][num]; ok {
				return false
			}

			boxInd := (i/3)*3 + j/3
			if _, ok := boxes[boxInd][num]; ok {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[boxInd][num] = true
		}
	}

	return true
}
