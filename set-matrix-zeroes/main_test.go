package main

func setZeroes(matrix [][]int) {
	firstCol := false
	firstRow := false
	for colIdx, col := range matrix {
		for rowIdx, val := range col {
			if val == 0 {
				if colIdx == 0 {
					firstRow = true
				} else {
					matrix[colIdx][0] = 0
				}
				if rowIdx == 0 {
					firstCol = true
				} else {
					matrix[0][rowIdx] = 0
				}
			}
		}
	}
	for colIdx := 1; colIdx < len(matrix); colIdx++ {
		if matrix[colIdx][0] == 0 {
			for idx := range matrix[colIdx] {
				matrix[colIdx][idx] = 0
			}
		}
	}
	for rowIdx := 1; rowIdx < len(matrix[0]); rowIdx++ {
		if matrix[0][rowIdx] == 0 {
			for idx := range matrix {
				matrix[idx][rowIdx] = 0
			}
		}
	}
	if firstRow {
		for idx := range matrix[0] {
			matrix[0][idx] = 0
		}
	}
	if firstCol {
		for idx := range matrix {
			matrix[idx][0] = 0
		}
	}
}
