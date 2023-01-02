package main

func recursive(grid [][]int, x, y, answer, want, current int) int {
	if y < 0 || len(grid) <= y {
		return answer
	}
	if x < 0 || len(grid[y]) <= x {
		return answer
	}
	if grid[y][x] == 2 {
		if want == current {
			return answer + 1
		}
		return answer
	}
	if grid[y][x] != 0 {
		return answer
	}
	grid[y][x] = -1
	answer = recursive(grid, x-1, y, answer, want, current+1)
	answer = recursive(grid, x+1, y, answer, want, current+1)
	answer = recursive(grid, x, y+1, answer, want, current+1)
	answer = recursive(grid, x, y-1, answer, want, current+1)
	grid[y][x] = 0
	return answer
}

func uniquePathsIII(grid [][]int) int {
	startY, startX := 0, 0
	want := 0
	for y, column := range grid {
		for x, i := range column {
			switch i {
			case 1:
				startY = y
				startX = x
			case 0:
				want++
			}
		}
	}
	grid[startY][startX] = 0
	return recursive(grid, startX, startY, 0, want, -1)
}
