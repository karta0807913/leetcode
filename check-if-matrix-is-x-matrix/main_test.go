package main

func checkXMatrix(grid [][]int) bool {
	left, right := 0, len(grid)-1
	for y, column := range grid {
		for x := range column {
			if x == left || x == right {
				if grid[y][x] == 0 {
					return false
				}
			} else if grid[y][x] != 0 {
				return false
			}
		}
		left += 1
		right -= 1
	}
	return true
}
