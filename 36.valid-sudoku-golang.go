package leetcode

// @leet start
func isValidSudoku(board [][]byte) bool {
	var rows, cols, squares [9][9]bool
	for i, row := range board {
		for j, n := range row {
			if n != '.' {
				k := int(n) - 49
				if rows[i][k] || cols[j][k] || squares[i/3*3+j/3][k] {
					return false
				}
				rows[i][k], cols[j][k], squares[i/3*3+j/3][k] = true, true, true
			}
		}
	}
	return true
}

// @leet end

