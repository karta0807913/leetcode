package main

func calSquare(x, y int) int {
	return x/3 + (y / 3 * 3)
}

func isValidSudoku(board [][]byte) bool {
	var cols, rows, squares [9]int64
	for y, col := range board {
		for x, b := range col {
			if b == '.' {
				continue
			}
			i, s := calSquare(x, y), int64(1<<(b))
			if (rows[x]|cols[y]|squares[i])&s != 0 {
				return false
			}
			rows[x] |= s
			cols[y] |= s
			squares[i] |= s
		}
	}
	return true
}
